package gotasia

import (
	"image/color"
	"strconv"
)

func colorTo255(c color.Color) (r, g, b, a uint32) {
	if c == nil {
		return 0, 0, 0, 255
	}

	r, g, b, a = c.RGBA()
	r /= 257
	g /= 257
	b /= 257
	a /= 257
	return r, g, b, a
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func colorTo1Scale(c color.Color) (r, g, b, a float64) {
	if c == nil {
		return 0, 0, 0, 1
	} else if r, ok := c.(rawColor); ok {
		return r.r, r.g, r.b, 1
	}

	ir, ig, ib, ia := c.RGBA()
	fR := float64(ir) / 65535.0
	fG := float64(ig) / 65535.0
	fB := float64(ib) / 65535.0
	fA := float64(ia) / 65535.0
	return fR, fG, fB, fA
}

func scale1ToColor(r, g, b float64) color.Color {
	return color.NRGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 255,
	}
}

type keepZero float64

func (f keepZero) MarshalJSON() ([]byte, error) {
	if float64(f) == float64(int(f)) {
		return []byte(strconv.FormatFloat(float64(f), 'f', 1, 32)), nil
	}
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64)), nil
}

func ref[T any](value T) *T {
	return &value
}

type rawColor struct {
	r, g, b float64
}

func (c rawColor) RGBA() (r, g, b, a uint32) {
	return uint32(c.r * 255), uint32(c.g * 255), uint32(c.b * 255), 255
}
