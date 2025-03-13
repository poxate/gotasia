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
