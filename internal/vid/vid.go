// Package vid emulates the two video modes the game uses, as DOSBox
// (machine=svga_s3) shows them in its 640x400 window:
//
//   - SCREEN 0 with 40 columns: 40x25 character cells with attributes,
//     drawn with the VGA 8x16 font, every pixel doubled horizontally;
//   - SCREEN 1: CGA 320x200 with 4 colours, every pixel doubled both ways,
//     text drawn with the 8x8 font.
//
// The screen is shared between the game goroutine (which writes) and the
// front end (which takes snapshots), so all access is locked.
package vid

import (
	"image"
	"image/color"
	"sync"
	"time"
)

// Size of the rendered image (the DOSBox window).
const (
	Width  = 640
	Height = 400
)

// Text mode geometry.
const (
	Rows = 25
	Cols = 40
)

// EGA is the 16 colour text palette (6 bit DAC values scaled like DOSBox).
var EGA = [16]color.RGBA{
	{0, 0, 0, 255}, {0, 0, 170, 255}, {0, 170, 0, 255}, {0, 170, 170, 255},
	{170, 0, 0, 255}, {170, 0, 170, 255}, {170, 85, 0, 255}, {170, 170, 170, 255},
	{85, 85, 85, 255}, {85, 85, 255, 255}, {85, 255, 85, 255}, {85, 255, 255, 255},
	{255, 85, 85, 255}, {255, 85, 255, 255}, {255, 255, 85, 255}, {255, 255, 255, 255},
}

// cgaPalettes are the SCREEN 1 foreground colours 1..3 (low intensity) as
// EGA palette indices, for COLOR ,0 and COLOR ,1.
var cgaPalettes = [2][3]int{{2, 4, 6}, {3, 5, 7}}

// Cell is a text mode character cell.
type Cell struct {
	Ch   byte
	Attr byte // bit 0-3 foreground, 4-6 background, 7 blink
}

// Screen is the emulated display.
type Screen struct {
	mu    sync.Mutex
	ops   uint64
	font8 []byte
	font  []byte

	graphics bool
	text     [Rows][Cols]Cell
	pix      [200][320]byte // SCREEN 1 pixels, 0..3
	bg       int            // SCREEN 1 background (EGA index)
	palette  int            // SCREEN 1 palette 0/1

	start time.Time
	// Blink, if set, overrides the blink phase (true = blinking characters
	// visible). Used for deterministic screenshots.
	Blink *bool
}

// New creates a screen in 40 column text mode.
func New(font8x8, font8x16 []byte) *Screen {
	s := &Screen{font8: font8x8, font: font8x16, start: time.Now()}
	s.clearText(0x07)
	return s
}

// Lock / Unlock give the game exclusive access for a sequence of calls.
func (s *Screen) Lock() { s.mu.Lock() }

// Unlock ends an exclusive access and counts a screen change.
func (s *Screen) Unlock() {
	s.ops++
	s.mu.Unlock()
}

// Ops returns a counter that changes whenever the screen may have changed.
func (s *Screen) Ops() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ops
}

// The following methods must be called with the lock held.

// Graphics reports whether SCREEN 1 is active.
func (s *Screen) Graphics() bool { return s.graphics }

// SetMode switches between text mode (false) and SCREEN 1 (true); the
// screen is cleared like the BIOS mode set does.
func (s *Screen) SetMode(graphics bool) {
	s.graphics = graphics
	s.pix = [200][320]byte{}
	s.bg = 0
	s.palette = 1
	s.clearText(0x07)
}

// SetCGAColors sets the SCREEN 1 background colour and palette.
func (s *Screen) SetCGAColors(bg, palette int) {
	if bg >= 0 {
		s.bg = bg & 15
	}
	if palette >= 0 {
		s.palette = palette & 1
	}
}

func (s *Screen) clearText(attr byte) {
	for r := range s.text {
		for c := range s.text[r] {
			s.text[r][c] = Cell{' ', attr}
		}
	}
}

// ClearRows clears text rows top..bottom (1-based) with attr; in graphics
// mode it clears the corresponding pixel rows to colour 0.
func (s *Screen) ClearRows(top, bottom int, attr byte) {
	for r := top; r <= bottom; r++ {
		if s.graphics {
			for y := (r - 1) * 8; y < r*8; y++ {
				s.pix[y] = [320]byte{}
			}
		} else {
			for c := range s.text[r-1] {
				s.text[r-1][c] = Cell{' ', attr}
			}
		}
	}
}

// ScrollUp scrolls rows top..bottom (1-based) up by one line, filling the
// last line with blanks of attr (colour 0 in graphics mode).
func (s *Screen) ScrollUp(top, bottom int, attr byte) {
	if s.graphics {
		for y := (top - 1) * 8; y < (bottom-1)*8; y++ {
			s.pix[y] = s.pix[y+8]
		}
	} else {
		for r := top - 1; r < bottom-1; r++ {
			s.text[r] = s.text[r+1]
		}
	}
	s.ClearRows(bottom, bottom, attr)
}

// PutChar writes a character at row, col (1-based). In graphics mode attr
// is the pixel colour of the glyph (0..3) and the cell background is 0.
func (s *Screen) PutChar(row, col int, ch, attr byte) {
	if row < 1 || row > Rows || col < 1 || col > Cols {
		return
	}
	if !s.graphics {
		s.text[row-1][col-1] = Cell{ch, attr}
		return
	}
	g := s.font8[int(ch)*8 : int(ch)*8+8]
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			c := byte(0)
			if g[y]&(0x80>>x) != 0 {
				c = attr & 3
			}
			s.pix[(row-1)*8+y][(col-1)*8+x] = c
		}
	}
}

// Pset sets a SCREEN 1 pixel (clipped).
func (s *Screen) Pset(x, y, c int) {
	if x >= 0 && x < 320 && y >= 0 && y < 200 {
		s.pix[y][x] = byte(c & 3)
	}
}

// Point returns a SCREEN 1 pixel, -1 outside the screen.
func (s *Screen) Point(x, y int) int {
	if x >= 0 && x < 320 && y >= 0 && y < 200 {
		return int(s.pix[y][x])
	}
	return -1
}

// Cell returns a text cell (1-based).
func (s *Screen) Cell(row, col int) Cell { return s.text[row-1][col-1] }

// HasBlink reports whether blinking characters are on screen.
func (s *Screen) HasBlink() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics {
		return false
	}
	for r := range s.text {
		for _, c := range s.text[r] {
			if c.Attr&0x80 != 0 {
				return true
			}
		}
	}
	return false
}

// blinkOn reports the current blink phase: the VGA toggles blinking
// characters every 16 frames of its 70 Hz refresh.
func (s *Screen) blinkOn() bool {
	if s.Blink != nil {
		return *s.Blink
	}
	frames := int(time.Since(s.start).Seconds() * 70)
	return frames/16%2 == 0
}

// Snapshot renders the screen into dst (Width x Height).
func (s *Screen) Snapshot(dst *image.RGBA) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics {
		pal := [4]color.RGBA{EGA[s.bg]}
		for i, c := range cgaPalettes[s.palette] {
			pal[i+1] = EGA[c]
		}
		for y := 0; y < Height; y++ {
			row := &s.pix[y/2]
			o := dst.PixOffset(0, y)
			for x := 0; x < Width; x++ {
				c := pal[row[x/2]]
				dst.Pix[o], dst.Pix[o+1], dst.Pix[o+2], dst.Pix[o+3] = c.R, c.G, c.B, 255
				o += 4
			}
		}
		return
	}
	blink := s.blinkOn()
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			cell := s.text[r][c]
			fg, bg := EGA[cell.Attr&15], EGA[cell.Attr>>4&7]
			if cell.Attr&0x80 != 0 && !blink {
				fg = bg
			}
			g := s.font[int(cell.Ch)*16 : int(cell.Ch)*16+16]
			for y := 0; y < 16; y++ {
				o := dst.PixOffset(c*16, r*16+y)
				bits := g[y]
				for x := 0; x < 16; x++ {
					col := bg
					if bits&(0x80>>(x/2)) != 0 {
						col = fg
					}
					dst.Pix[o], dst.Pix[o+1], dst.Pix[o+2], dst.Pix[o+3] = col.R, col.G, col.B, 255
					o += 4
				}
			}
		}
	}
}

// Image returns a snapshot as a new image.
func (s *Screen) Image() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	s.Snapshot(img)
	return img
}
