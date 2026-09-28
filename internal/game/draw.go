package game

import (
	"encoding/binary"
	"math"
	"strings"
)

// VARPTR$ results are only used inside DRAW strings ("S=" + VARPTR$(n!)
// and "X" + VARPTR$(a$)). Instead of a memory address the port embeds the
// value itself behind a 0 byte, which cannot occur in a DRAW string.
const (
	vpSingle = 4
	vpString = 3
)

func varptrF(v float32) string {
	var b [6]byte
	b[1] = vpSingle
	binary.LittleEndian.PutUint32(b[2:], math.Float32bits(v))
	return string(b[:])
}

func varptrS(s string) string {
	var b [4]byte
	b[1] = vpString
	binary.LittleEndian.PutUint16(b[2:], uint16(len(s)))
	return string(b[:]) + s
}

// setPoint sets the coordinates of the next graphics statement (PAINT).
func (g *Game) setPoint(x, y int) {
	g.lastX, g.lastY = x, y
}

// line draws a line with the Bresenham algorithm, both ends included.
func (g *Game) line(x0, y0, x1, y1, c int) {
	dx, dy := abs(x1-x0), -abs(y1-y0)
	sx, sy := sign(x1-x0), sign(y1-y0)
	e := dx + dy
	for {
		g.scr.Pset(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := 2 * e
		if e2 >= dy {
			e += dy
			x0 += sx
		}
		if e2 <= dx {
			e += dx
			y0 += sy
		}
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func sign(v int) int {
	switch {
	case v < 0:
		return -1
	case v > 0:
		return 1
	}
	return 0
}

var drawDirs = map[byte][2]int{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0},
	'E': {1, -1}, 'F': {1, 1}, 'G': {-1, 1}, 'H': {-1, -1}}

// drawParser walks a DRAW command string.
type drawParser struct {
	s string
	i int
}

func (p *drawParser) skip() {
	for p.i < len(p.s) && (p.s[p.i] == ' ' || p.s[p.i] == ';') {
		p.i++
	}
}

// varptr reads an embedded VARPTR$ value if one follows.
func (p *drawParser) varptr() (typ byte, f float32, s string, ok bool) {
	if p.i+1 < len(p.s) && p.s[p.i] == 0 {
		typ = p.s[p.i+1]
		switch typ {
		case vpSingle:
			f = math.Float32frombits(binary.LittleEndian.Uint32([]byte(p.s[p.i+2 : p.i+6])))
			p.i += 6
		case vpString:
			n := int(binary.LittleEndian.Uint16([]byte(p.s[p.i+2 : p.i+4])))
			s = p.s[p.i+4 : p.i+4+n]
			p.i += 4 + n
		}
		return typ, f, s, true
	}
	return 0, 0, "", false
}

// number reads a numeric argument; ok is false if none is present.
func (p *drawParser) number() (v int, ok bool) {
	p.skip()
	if p.i < len(p.s) && p.s[p.i] == '=' {
		p.i++
		if _, f, _, ok := p.varptr(); ok {
			return cint(float64(f)), true
		}
		panic("Illegal function call")
	}
	neg := false
	if p.i < len(p.s) && (p.s[p.i] == '-' || p.s[p.i] == '+') {
		neg = p.s[p.i] == '-'
		p.i++
		p.skip()
	}
	st := p.i
	for p.i < len(p.s) && p.s[p.i] >= '0' && p.s[p.i] <= '9' {
		v = v*10 + int(p.s[p.i]-'0')
		p.i++
	}
	if p.i == st {
		return 0, false
	}
	if neg {
		v = -v
	}
	return v, true
}

// draw is the DRAW statement.
func (g *Game) draw(cmd string) {
	g.scr.Lock()
	defer g.scr.Unlock()
	g.drawString(cmd)
}

func (g *Game) drawString(cmd string) {
	p := &drawParser{s: cmd}
	blank, noUpdate := false, false
	for {
		p.skip()
		if p.i >= len(p.s) {
			return
		}
		c := p.s[p.i]
		if c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		p.i++
		switch c {
		case 'B':
			blank = true
			continue
		case 'N':
			noUpdate = true
			continue
		case 'U', 'D', 'L', 'R', 'E', 'F', 'G', 'H':
			n, ok := p.number()
			if !ok {
				n = 1
			}
			d := n * g.drawScale / 4
			dir := drawDirs[c]
			g.moveTo(g.lastX+dir[0]*d, g.lastY+dir[1]*d, blank, noUpdate)
		case 'M':
			p.skip()
			rel := p.i < len(p.s) && (p.s[p.i] == '+' || p.s[p.i] == '-')
			x, _ := p.number()
			p.skip()
			if p.i < len(p.s) && p.s[p.i] == ',' {
				p.i++
			}
			y, _ := p.number()
			if rel {
				x = g.lastX + x*g.drawScale/4
				y = g.lastY + y*g.drawScale/4
			}
			g.moveTo(x, y, blank, noUpdate)
		case 'C':
			n, _ := p.number()
			g.drawColor = n
		case 'S':
			n, _ := p.number()
			g.drawScale = n
		case 'A':
			n, _ := p.number()
			g.drawAngle = n
		case 'X':
			p.skip()
			if _, _, s, ok := p.varptr(); ok {
				g.drawString(s)
			}
		case 'P':
			pc, _ := p.number()
			p.skip()
			if p.i < len(p.s) && p.s[p.i] == ',' {
				p.i++
			}
			bc, _ := p.number()
			g.fill(g.lastX, g.lastY, pc, bc)
		default:
			panic("Illegal function call: DRAW " + strings.TrimSpace(string(c)))
		}
		blank, noUpdate = false, false
	}
}

func (g *Game) moveTo(x, y int, blank, noUpdate bool) {
	if !blank {
		g.line(g.lastX, g.lastY, x, y, g.drawColor)
	}
	if !noUpdate {
		g.lastX, g.lastY = x, y
	}
}

// paint is PAINT (lastX, lastY), paint, border.
func (g *Game) paint(pc, bc int) {
	g.scr.Lock()
	defer g.scr.Unlock()
	g.fill(g.lastX, g.lastY, pc, bc)
}

// fill floods the area around x, y bounded by the border colour.
func (g *Game) fill(x, y, pc, bc int) {
	if c := g.scr.Point(x, y); c < 0 || c == bc&3 {
		return
	}
	seen := map[[2]int]bool{}
	stack := [][2]int{{x, y}}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if seen[p] {
			continue
		}
		c := g.scr.Point(p[0], p[1])
		if c < 0 || c == bc&3 {
			continue
		}
		seen[p] = true
		g.scr.Pset(p[0], p[1], pc)
		stack = append(stack, [2]int{p[0] + 1, p[1]}, [2]int{p[0] - 1, p[1]},
			[2]int{p[0], p[1] + 1}, [2]int{p[0], p[1] - 1})
	}
}
