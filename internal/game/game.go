// Package game is the Go port of Szakadt csöves '88.
//
// The game logic (main_gen.go) is translated from the compiled QuickBASIC
// code of CSOVES.EXE by a custom translator. It keeps the structure of the
// original: the module level variables are the fields of vars, SUBs are
// methods, GOSUB/RETURN is a small return stack, and every statement calls
// the same QuickBASIC runtime operation as the original (qb.go, draw.go).
package game

import (
	"io"

	"tclinux-csoves/internal/dos"
	"tclinux-csoves/internal/vid"
)

// Game holds the complete state of the original program.
type Game struct {
	m   *dos.Machine
	scr *vid.Screen

	vars
	gosub []int

	// text cursor (1-based) and colours
	row, col  int
	fg, bg    int
	cursorOn  bool
	lastX     int // graphics "last point referenced"
	lastY     int
	drawColor int
	drawScale int
	drawAngle int

	seed uint32 // RND state (24 bits)

	// NoRandomize makes RANDOMIZE TIMER a no-op, so RND starts from the
	// built-in seed like in the patched original used by the tests.
	NoRandomize bool
	// Trace, if set, receives a line of the main variables at every
	// keyboard read (debugging aid).
	Trace io.Writer
	// Tune overrides a few constants of the program for tests (marked in
	// main_gen.go by g.tune), e.g. the start position.
	Tune map[string]float64
}

// New creates the game on the given machine.
func New(m *dos.Machine) *Game {
	return &Game{m: m, scr: m.Scr, row: 1, col: 1, fg: 7, seed: 0x50000,
		drawColor: 3, drawScale: 4}
}

// Run executes the program until END (or until the front end quits).
func (g *Game) Run() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(dos.HaltError); ok {
				return
			}
			panic(r)
		}
	}()
	g.mainProgram()
}

func (g *Game) tune(name string, v float64) float64 {
	if t, ok := g.Tune[name]; ok {
		return t
	}
	return v
}

func (g *Game) gosubPop() int {
	n := len(g.gosub)
	if n == 0 {
		g.runtimeError("RETURN without GOSUB")
	}
	r := g.gosub[n-1]
	g.gosub = g.gosub[:n-1]
	return r
}

// clear is the CLEAR statement: all variables are reset.
func (g *Game) clear() {
	g.vars = vars{}
	g.gosub = nil
}

// end is the END statement.
func (g *Game) end() { g.m.Halt() }

// caseElseError is the run-time error of a SELECT CASE without a matching
// CASE (the program never reaches it in normal play).
func (g *Game) caseElseError() { g.runtimeError("CASE ELSE expected") }

func (g *Game) runtimeError(msg string) {
	g.printS("", nl)
	g.printS(msg, nl)
	g.printS("Hit any key to return to system", nl)
	g.m.ReadKey()
	g.m.WriteLn(msg)
	g.m.Halt()
}
