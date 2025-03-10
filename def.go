package gotasia

import "github.com/sanity-io/litter"

func (p *Project) createDef(e *Element) jobj {
	switch node := e.Node.(type) {
	case *Callout:
		fillStyle := node.FillStyle
		if fillStyle == "" {
			fillStyle = "solid"
		}

		var fR, fG, fB, fA uint32
		if node.FillColor != nil {
			fR, fG, fB, fA = node.FillColor.RGBA()
		}

		return jobj{
			"kind":                    "remix",
			"width":                   node.Width,
			"height":                  node.Height,
			"shape":                   node.Shape.string(),
			"style":                   "basic",
			"corner-radius":           node.CornerRadius,
			"enable-ligatures":        1.0,
			"fill-style":              fillStyle,
			"fill-color-red":          float64(fR) / 65535.0,
			"fill-color-green":        float64(fG) / 65535.0,
			"fill-color-blue":         float64(fB) / 65535.0,
			"fill-color-opacity":      float64(fA) / 65535.0,
			"line-spacing":            0.0,
			"stroke-color-blue":       1.0,
			"stroke-color-green":      1.0,
			"stroke-color-opacity":    1.0,
			"stroke-color-red":        1.0,
			"stroke-width":            0.0,
			"tail-x":                  node.TailX,
			"tail-y":                  -node.TailY,
			"text-stroke-alignment":   2.0,
			"text-stroke-color-alpha": 1.0,
			"text-stroke-color-blue":  0.0,
			"text-stroke-color-green": 0.0,
			"text-stroke-color-red":   0.0,
			"text-stroke-width":       0.0,
			"word-wrap":               1.0,
			"horizontal-alignment":    "center",
			"resize-behavior":         "resizeText",
			"stroke-style":            "solid",
			"text":                    node.Text,
			"vertical-alignment":      "center",
			"font": jobj{
				"color-blue":  1.0,
				"color-green": 1.0,
				"color-red":   1.0,
				"size":        64.0,
				"tracking":    0.0,
				"name":        "Montserrat",
				"weight":      "Regular",
			},
			"textAttributes": jobj{
				"type": "textAttributeList",
				"keyframes": []jobj{
					{
						"endTime":  0,
						"time":     0,
						"value":    nil,
						"duration": 0,
					},
				},
			},
		}
	default:
		panic("unknown element node: " + litter.Sdump(node))
	}
}
