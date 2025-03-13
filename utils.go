package gotasia

import "image/color"

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
	}

	ir, ig, ib, ia := c.RGBA()
	fR := float64(ir) / 65535.0
	fG := float64(ig) / 65535.0
	fB := float64(ib) / 65535.0
	fA := float64(ia) / 65535.0
	return fR, fG, fB, fA
}
