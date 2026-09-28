// Package dos provides the small slice of DOS the game relies on: a
// BIOS-like keyboard buffer (with KeyPressed / ReadKey semantics), Delay and
// Halt.
//
// The game logic runs sequentially in its own goroutine exactly like the
// original program did; the front end (window or headless runner) feeds keys
// and displays the screen.
package dos

import (
	"strings"
	"sync"
	"time"

	"tclinux-csoves/internal/vid"
)

// HaltError is raised (via panic) by Halt and when the front end shuts down.
type HaltError struct{ Code int }

// Machine is the runtime environment shared by the game and the front end.
type Machine struct {
	Scr *vid.Screen

	mu   sync.Mutex
	cond *sync.Cond
	keys []byte
	quit bool
	// Idle, if set, is called (without the lock held) whenever the game polls
	// or waits for a key and the buffer is empty. waiting is true when the
	// game is really waiting for input (ReadKey, or KeyPressed polled
	// repeatedly) rather than draining the buffer. The headless runner uses
	// it to feed scripted input. It returns true if it did something (added
	// a key or called Quit), false to let the machine sleep and poll again.
	Idle       func(waiting bool) bool
	emptyPolls int
	lastOps    uint64
	// Fast disables all real-time waiting (Delay, retrace) for tests.
	Fast bool

	// Console collects what the program writes to the DOS console after
	// leaving graphics mode.
	Console strings.Builder
}

func New(scr *vid.Screen) *Machine {
	m := &Machine{Scr: scr}
	m.cond = sync.NewCond(&m.mu)
	return m
}

// PushKey appends a raw ReadKey byte to the keyboard buffer. Extended keys
// are pushed as 0 followed by the scan code.
func (m *Machine) PushKey(b ...byte) {
	m.mu.Lock()
	m.keys = append(m.keys, b...)
	m.mu.Unlock()
	m.cond.Broadcast()
}

// Quit makes the next keyboard or timing call in the game goroutine halt.
func (m *Machine) Quit() {
	m.mu.Lock()
	m.quit = true
	m.mu.Unlock()
	m.cond.Broadcast()
}

func (m *Machine) checkQuit() {
	if m.quit {
		m.mu.Unlock()
		panic(HaltError{})
	}
}

func (m *Machine) idle(waiting bool) {
	if m.Idle != nil && m.Idle(waiting) {
		return
	}
	time.Sleep(5 * time.Millisecond)
}

// KeyPressed reports whether a key is waiting (Crt.KeyPressed).
func (m *Machine) KeyPressed() bool {
	m.mu.Lock()
	m.checkQuit()
	n := len(m.keys)
	m.mu.Unlock()
	if n > 0 {
		return true
	}
	if ops := m.Scr.Ops(); ops != m.lastOps {
		m.lastOps = ops
		m.emptyPolls = 0
	}
	m.emptyPolls++
	if m.emptyPolls == 1 {
		// First empty poll: typically the end of a buffer drain loop.
		// Report without waiting so the drain finishes immediately.
		return false
	}
	m.idle(true)
	m.mu.Lock()
	m.checkQuit()
	n = len(m.keys)
	m.mu.Unlock()
	return n > 0
}

// ReadKey waits for and returns the next key byte (Crt.ReadKey).
func (m *Machine) ReadKey() byte {
	for {
		m.mu.Lock()
		m.checkQuit()
		if len(m.keys) > 0 {
			k := m.keys[0]
			m.keys = m.keys[1:]
			m.emptyPolls = 0
			m.mu.Unlock()
			return k
		}
		m.mu.Unlock()
		m.idle(true)
	}
}

// Delay sleeps ms milliseconds (Crt.Delay).
func (m *Machine) Delay(ms int) {
	m.sleep(time.Duration(ms) * time.Millisecond)
}

func (m *Machine) sleep(d time.Duration) {
	m.mu.Lock()
	m.checkQuit()
	m.mu.Unlock()
	if !m.Fast {
		time.Sleep(d)
	}
}

// Halt terminates the game program.
func (m *Machine) Halt() { panic(HaltError{}) }

// WriteLn writes a line to the DOS console (shown after the program ends).
func (m *Machine) WriteLn(s string) {
	m.Console.WriteString(s)
	m.Console.WriteByte('\n')
}
