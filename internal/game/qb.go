package game

import (
	"fmt"
	"math"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// omit marks an argument left out of a statement (e.g. LOCATE ,,0).
const omit = math.MinInt32

// PRINT item terminators.
type printEnd int

const (
	nl   printEnd = iota // PRINT x
	semi                 // PRINT x;
	tab                  // PRINT x,
)

// Text view port: PRINT scrolls lines 1..24, line 25 is only reachable
// with LOCATE.
const viewBottom = 24

// ------------------------------------------------------------ screen

func (g *Game) attr() byte {
	if g.scr.Graphics() {
		return 3
	}
	return byte(g.fg&15 | (g.bg&7)<<4 | (g.fg&16)<<3)
}

// screen is the SCREEN statement.
func (g *Game) screen(args ...int) {
	mode := args[0]
	g.scr.Lock()
	defer g.scr.Unlock()
	graphics := mode == 1
	if graphics == g.scr.Graphics() && mode != omit {
		return
	}
	g.scr.SetMode(graphics)
	g.fg, g.bg = 7, 0
	if graphics {
		g.fg = 3
	}
	g.row, g.col = 1, 1
	g.lastX, g.lastY = 160, 100
	g.drawColor, g.drawScale, g.drawAngle = 3, 4, 0
}

func arg(args []int, i int) int {
	if i < len(args) {
		return args[i]
	}
	return omit
}

// color is the COLOR statement.
func (g *Game) color(args ...int) {
	a, b := arg(args, 0), arg(args, 1)
	g.scr.Lock()
	defer g.scr.Unlock()
	if g.scr.Graphics() {
		bg, pal := -1, -1
		if a != omit {
			bg = a
		}
		if b != omit {
			pal = b
		}
		g.scr.SetCGAColors(bg, pal)
		return
	}
	if a != omit {
		g.fg = a & 31
	}
	if b != omit {
		g.bg = b & 15
	}
	// the border colour (third argument) is not visible in DOSBox
}

// locate is the LOCATE statement.
func (g *Game) locate(args ...int) {
	r, c, cur := arg(args, 0), arg(args, 1), arg(args, 2)
	if r != omit {
		if r < 1 || r > 25 {
			g.runtimeError("Illegal function call")
		}
		g.row = r
	}
	if c != omit {
		if c < 1 || c > 40 {
			g.runtimeError("Illegal function call")
		}
		g.col = c
	}
	if cur != omit {
		g.cursorOn = cur != 0
	}
}

// cls is the CLS statement.
func (g *Game) cls(n int) {
	g.scr.Lock()
	defer g.scr.Unlock()
	g.scr.ClearRows(1, 25, g.clearAttr())
	g.row, g.col = 1, 1
	if g.scr.Graphics() {
		g.lastX, g.lastY = 160, 100
	}
}

func (g *Game) clearAttr() byte { return byte((g.bg & 7) << 4) }

// newline moves the cursor to the start of the next line, scrolling the
// view port when needed. Called with the screen locked.
func (g *Game) newline() {
	g.col = 1
	switch {
	case g.row < viewBottom:
		g.row++
	case g.row == viewBottom:
		g.scr.ScrollUp(1, viewBottom, g.clearAttr())
	default: // on line 25
		g.scr.ScrollUp(1, viewBottom, g.clearAttr())
		g.row = viewBottom
	}
}

// out writes one PRINT item.
func (g *Game) out(s string, end printEnd) {
	g.scr.Lock()
	defer g.scr.Unlock()
	if g.col > 1 && g.col+len(s)-1 > 40 {
		g.newline()
	}
	a := g.attr()
	for i := 0; i < len(s); i++ {
		if g.col > 40 {
			g.newline()
		}
		g.scr.PutChar(g.row, g.col, s[i], a)
		g.col++
	}
	switch end {
	case nl:
		g.newline()
	case tab:
		// next 14 column print zone; if it does not fit completely, go
		// to the next line
		z := ((g.col-1)/14+1)*14 + 1
		if z+13 > 40 {
			g.newline()
			break
		}
		for g.col < z {
			g.scr.PutChar(g.row, g.col, ' ', a)
			g.col++
		}
	}
	if g.col > 40 {
		g.newline()
	}
}

// printS is PRINT of a string item.
func (g *Game) printS(s string, end printEnd) { g.out(s, end) }

// printF is PRINT of a SINGLE item: sign or space, digits, space.
func (g *Game) printF(v float64, end printEnd) { g.out(fmtSingle(v)+" ", end) }

// fmtSingle formats a SINGLE like STR$: a leading space for non negative
// numbers, at most 7 significant digits, no leading zero before the point.
func fmtSingle(v float64) string {
	f := float64(float32(v))
	sign := " "
	if f < 0 {
		sign, f = "-", -f
	}
	s := strconv.FormatFloat(f, 'g', 7, 32)
	if mant, exp, ok := strings.Cut(s, "e"); ok {
		es := exp[:1]
		exp = strings.TrimLeft(exp[1:], "0")
		for len(exp) < 2 {
			exp = "0" + exp
		}
		s = mant + "E" + es + exp
	}
	if s != "0" {
		s = strings.TrimPrefix(s, "0")
	}
	return sign + s
}

// ------------------------------------------------------------ keyboard

// inputS is INPUT$(n) from the keyboard.
func (g *Game) inputS(n int) string {
	if g.Trace != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Fprintf(g.Trace, "INPUT$ at=%s:%d energy=%v pos=%v,%v hour=%v/%v money=%v glue=%v booze=%v police=%v seed=%06x\n",
			filepath.Base(file), line, g.f162, g.f146, g.f14a, g.f1a6, g.f1c2, g.f182, g.f1b6, g.f1b2, g.f192, g.seed)
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = g.m.ReadKey()
	}
	if g.Trace != nil {
		fmt.Fprintf(g.Trace, "   -> %q\n", b)
	}
	return string(b)
}

// inkey is INKEY$.
func (g *Game) inkey() string {
	if !g.m.KeyPressed() {
		return ""
	}
	b := g.m.ReadKey()
	if b == 0 {
		return string([]byte{0, g.m.ReadKey()})
	}
	return string([]byte{b})
}

// ------------------------------------------------------------ numbers

// rnd is RND: the QuickBASIC 24 bit linear congruential generator.
func (g *Game) rnd() float64 {
	if g.Trace != nil {
		fmt.Fprintf(g.Trace, "   RND\n")
	}
	g.seed = (g.seed*0xFD43FD + 0xC39EC3) & 0xFFFFFF
	return float64(g.seed) / 16777216
}

// randomize is RANDOMIZE x (x passed as a DOUBLE): the middle two bytes of
// the seed become the XOR of the two high words of x.
func (g *Game) randomize(x float64) {
	if g.NoRandomize {
		return
	}
	b := math.Float64bits(x)
	w := uint32(b>>32&0xffff ^ b>>48&0xffff)
	g.seed = g.seed&0xff | w<<8
}

// timer is TIMER: seconds since midnight in 1/100 s resolution.
func (g *Game) timer() float64 {
	t := time.Now()
	n := ((t.Hour()*60+t.Minute())*60+t.Second())*100 + t.Nanosecond()/10_000_000
	return float64(float32(float64(n) / 100))
}

// cint rounds like the 8087 FISTP (to nearest, ties to even).
func cint(v float64) int {
	r := math.RoundToEven(v)
	if r < -32768 || r > 32767 {
		panic("Overflow")
	}
	return int(r)
}

func b2i(b bool) int {
	if b {
		return -1
	}
	return 0
}

// ------------------------------------------------------------ strings

func chr(n int) string { return string([]byte{byte(n)}) }

func asc(s string) int {
	if s == "" {
		panic("Illegal function call")
	}
	return int(s[0])
}

func mid(s string, start, n int) string {
	if start < 1 || n < 0 {
		panic("Illegal function call")
	}
	if start > len(s) {
		return ""
	}
	s = s[start-1:]
	if n < len(s) {
		s = s[:n]
	}
	return s
}

func left(s string, n int) string {
	if n < len(s) {
		return s[:n]
	}
	return s
}

func right(s string, n int) string {
	if n < len(s) {
		return s[len(s)-n:]
	}
	return s
}

func instr(start int, s, t string) int {
	if start > len(s) {
		return 0
	}
	i := strings.Index(s[start-1:], t)
	if i < 0 {
		return 0
	}
	return i + start
}

func spaces(n int) string { return strings.Repeat(" ", n) }

func stringFill(n, c int) string { return strings.Repeat(chr(c), n) }

// midStmt is MID$(dst, start, n) = src.
func midStmt(dst string, start, n int, src string) string {
	if start < 1 || start > len(dst) {
		panic("Illegal function call")
	}
	k := n
	if len(src) < k {
		k = len(src)
	}
	if len(dst)-start+1 < k {
		k = len(dst) - start + 1
	}
	b := []byte(dst)
	copy(b[start-1:], src[:k])
	return string(b)
}
