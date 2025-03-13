package gotasia

import (
	"image/color"
	"strconv"
)

type CalloutShape int
type CalloutFillStyle string

const (
	CalloutShapeSpeechBubble1 CalloutShape = iota
	CalloutShapeSpeechBubble2
	CalloutShapeThoughtBubble1
	CalloutShapeThoughtBubble2
	CalloutShapeTextArrow1
	CalloutShapeTextArrow2
	CalloutShapeTextRectangle
)

const (
	CalloutFillSolid    CalloutFillStyle = "solid"
	CalloutFillGradient CalloutFillStyle = "gradient"
)

type Callout struct {
	Text  string
	Shape CalloutShape
	// if FillStyle is not specified, will render to solid
	FillStyle    CalloutFillStyle
	FillColor    color.Color
	Width        int
	Height       int
	TailX        int // distance from the center of the callout along the x-axis, only used for speech and thought bubbles
	TailY        int // distance from the center of the callout along the y-axis, only used for speech and thought bubbles
	CornerRadius int // only used for text-rectangle
}

func (shape CalloutShape) string() string {
	switch shape {
	case CalloutShapeSpeechBubble1:
		return "speech-bubble"
	case CalloutShapeSpeechBubble2:
		return "speech-bubble2"
	case CalloutShapeThoughtBubble1:
		return "thought-bubble"
	case CalloutShapeThoughtBubble2:
		return "thought-bubble2"
	case CalloutShapeTextArrow1:
		return "text-arrow"
	case CalloutShapeTextArrow2:
		return "text-arrow2"
	case CalloutShapeTextRectangle:
		return "text-rectangle"
	default:
		panic("unknown callout shape representation: " + strconv.Itoa(int(shape)))
	}
}

func (c *Callout) width() int  { return c.Width }
func (c *Callout) height() int { return c.Height }

func (node *Callout) createDef() jobj {
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
}
