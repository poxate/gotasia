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
