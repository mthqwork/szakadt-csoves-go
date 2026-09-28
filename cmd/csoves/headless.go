package main

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"tclinux-csoves/internal/dos"
	"tclinux-csoves/internal/game"
)

// runHeadless runs the game without a window. Input comes from the script;
// the game runs as fast as possible. When the script is exhausted the game
// is stopped and the final screen is written to shot (if given).
func runHeadless(m *dos.Machine, g *game.Game, script, shot string) error {
	items, err := parseScript(script)
	if err != nil {
		return err
	}
	m.Fast = true
	var ioErr error
	quit := false
	m.Idle = func(waiting bool) bool {
		if !waiting {
			return false
		}
		for len(items) > 0 && items[0].shot != "" {
			if err := savePNG(m, items[0].shot); err != nil && ioErr == nil {
				ioErr = err
			}
			items = items[1:]
		}
		if len(items) == 0 {
			if os.Getenv("CSOVES_TRACE") != "" {
				fmt.Fprintln(os.Stderr, "WAIT")
			}
			quit = true
			m.Quit()
			return true
		}
		m.PushKey(items[0].keys...)
		items = items[1:]
		return true
	}
	g.Run()
	if !quit {
		// the program ended by itself (END) before the script ran out
		fmt.Fprintln(os.Stderr, "program ended")
	}
	if shot != "" {
		if err := savePNG(m, shot); err != nil {
			return err
		}
	}
	return ioErr
}

// savePNG writes the screen with blinking characters visible; if there are
// blinking characters on screen, the other blink phase is also written to
// NAME.blink.png (the tests accept either).
func savePNG(m *dos.Machine, name string) error {
	if m.Scr.HasBlink() {
		off := false
		m.Scr.Blink = &off
		err := writePNG(m, strings.TrimSuffix(name, ".png")+".blink.png")
		on := true
		m.Scr.Blink = &on
		if err != nil {
			return err
		}
	}
	return writePNG(m, name)
}

func writePNG(m *dos.Machine, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, m.Scr.Image()); err != nil {
		return fmt.Errorf("writing %s: %w", name, err)
	}
	return nil
}
