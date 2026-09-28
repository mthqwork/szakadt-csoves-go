// Package assets embeds the BIOS fonts of the VGA card that DOSBox emulates,
// dumped from DOSBox (INT 10h AX=1130h).
package assets

import _ "embed"

// Font8x8 is the 8x8 font used in graphics modes (256 characters).
//
//go:embed font8x8.bin
var Font8x8 []byte

// Font8x16 is the 8x16 font of the VGA text modes (256 characters).
//
//go:embed font8x16.bin
var Font8x16 []byte
