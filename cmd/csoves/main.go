// Command csoves is the native port of the DOS game "Szakadt csöves '88"
// (Pató Pál Software, 1989).
//
//	csoves                      play in a window
//	csoves -headless -keys ...  run without a window (for tests), see keys.go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"tclinux-csoves/assets"
	"tclinux-csoves/internal/dos"
	"tclinux-csoves/internal/game"
	"tclinux-csoves/internal/vid"
)

func main() {
	headless := flag.Bool("headless", false, "run without a window, driven by -keys")
	keys := flag.String("keys", "", "scripted input for -headless (see keys.go)")
	shot := flag.String("shot", "", "headless: write a PNG of the final screen here")
	scale := flag.Int("scale", 2, "window scale factor")
	var tune tuneFlag
	flag.Var(&tune, "tune", "override a program constant for tests: NAME=VALUE (repeatable)")
	norand := flag.Bool("norandomize", false, "skip RANDOMIZE TIMER (reproducible games, for tests)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	scr := vid.New(assets.Font8x8, assets.Font8x16)
	m := dos.New(scr)
	g := game.New(m)
	g.NoRandomize = *norand
	g.Tune = tune
	if os.Getenv("CSOVES_TRACE") != "" {
		g.Trace = os.Stderr
	}

	var err error
	if *headless {
		on := true
		scr.Blink = &on
		err = runHeadless(m, g, *keys, *shot)
	} else {
		err = runWindow(m, g, *scale)
	}
	if s := m.Console.String(); s != "" {
		fmt.Print(s)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// tuneFlag collects -tune NAME=VALUE options.
type tuneFlag map[string]float64

func (t *tuneFlag) String() string { return fmt.Sprint(map[string]float64(*t)) }

func (t *tuneFlag) Set(s string) error {
	name, val, ok := strings.Cut(s, "=")
	v, err := strconv.ParseFloat(val, 64)
	if !ok || err != nil {
		return fmt.Errorf("want NAME=VALUE, got %q", s)
	}
	if *t == nil {
		*t = tuneFlag{}
	}
	(*t)[name] = v
	return nil
}
