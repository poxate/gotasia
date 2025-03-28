package gotasia

import (
	"image/color"
)

type CalloutShape string
type CalloutFillStyle string
type CalloutStrokeStyle string
type FontWeight string
type CalloutHorizontalAlignment string
type CalloutVerticalAlignment string

const (
	CalloutShapeText           CalloutShape = "text"
	CalloutShapeSpeechBubble1  CalloutShape = "speech-bubble"
	CalloutShapeSpeechBubble2  CalloutShape = "speech-bubble2"
	CalloutShapeThoughtBubble1 CalloutShape = "thought-bubble"
	CalloutShapeThoughtBubble2 CalloutShape = "thought-bubble2"
	CalloutShapeTextArrow1     CalloutShape = "text-arrow"
	CalloutShapeTextArrow2     CalloutShape = "text-arrow2"
	CalloutShapeTextRectangle  CalloutShape = "text-rectangle"
)

const (
	CalloutFillSolid    CalloutFillStyle = "solid"
	CalloutFillGradient CalloutFillStyle = "gradient"
)

const (
	CalloutStrokeSolid      CalloutStrokeStyle = "solid"
	CalloutStrokeDot        CalloutStrokeStyle = "dot"
	CalloutStrokeDash       CalloutStrokeStyle = "dash"
	CalloutStrokeDashDot    CalloutStrokeStyle = "dashdot"
	CalloutStrokeDashDotDot CalloutStrokeStyle = "dashdotdot"
)

const (
	FontWeightRegular FontWeight = "Regular"
	FontWeightBold    FontWeight = "Bold"
)

const (
	CalloutHorizontalAlignmentLeft   CalloutHorizontalAlignment = "left"
	CalloutHorizontalAlignmentCenter CalloutHorizontalAlignment = "center"
	CalloutHorizontalAlignmentRight  CalloutHorizontalAlignment = "right"

	CalloutVerticalAlignemntTop    CalloutVerticalAlignment = "top"
	CalloutVerticalAlignemntCenter CalloutVerticalAlignment = "center"
	CalloutVerticalAlignemntBottom CalloutVerticalAlignment = "bottom"
)

type Callout struct {
	_text string // ignored if spans is not nil
	Font  Font

	Shape CalloutShape

	// if FillStyle is not specified, will render to solid
	FillStyle   CalloutFillStyle
	FillColor   color.Color
	FillOpacity float64

	StrokeColor   color.Color
	StrokeOpacity float64
	StrokeWidth   float64
	StrokeStyle   CalloutStrokeStyle

	Width        float64
	Height       float64
	TailX        float64 // distance from the center of the callout along the x-axis, only used for speech and thought bubbles
	TailY        float64 // distance from the center of the callout along the y-axis, only used for speech and thought bubbles
	CornerRadius int     // only used for text-rectangle

	VerticalAlignment   CalloutVerticalAlignment
	HorizontalAlignment CalloutHorizontalAlignment
}

type Span struct {
	Text          string
	FontSize      int
	Underline     bool
	Color         color.Color
	Weight        int
	Italic        bool
	FontName      string
	Kerning       float64
	Strikethrough bool
}

type Font struct {
	Color    color.Color
	Size     float64
	Tracking float64
	Name     string
	Weight   FontWeight
}

func NewCallout() *Callout {
	return &Callout{
		Width:               400,
		Height:              250,
		Shape:               CalloutShapeText,
		CornerRadius:        8,
		VerticalAlignment:   CalloutVerticalAlignemntCenter,
		HorizontalAlignment: CalloutHorizontalAlignmentCenter,
		_text:               "ABC",
		FillOpacity:         1,
		Font: Font{
			Color:    color.White,
			Size:     96,
			Tracking: 0,
			Name:     "Montserrat",
			Weight:   "Regular",
		},
	}
}

func (c *Callout) SetText(newText string) *Callout {
	c._text = newText
	return c
}

func (c *Callout) width() int  { return int(c.Width) }
func (c *Callout) height() int { return int(c.Height) }

func (node *Callout) encodeIntoMedia(rMedia *rawMedia) {
	var fontR, fontG, fontB float64
	if raw, ok := node.Font.Color.(rawColor); ok {
		fontR, fontG, fontB = raw.r, raw.g, raw.b
	} else {
		fontR, fontG, fontB, _ = colorTo1Scale(node.Font.Color)
	}

	def := &rawDef{
		Kind:                 "remix",
		Shape:                string(node.Shape),
		Style:                "basic",
		Width:                keepZero(node.Width),
		Height:               keepZero(node.Height),
		CornerRadius:         keepZero(node.CornerRadius),
		EnableLigatures:      1.0,
		LineSpacing:          0.0,
		TextStrokeAlignment:  2.0,
		TextStrokeColorAlpha: 1.0,
		TextStrokeColorBlue:  0.0,
		TextStrokeColorGreen: 0.0,
		TextStrokeColorRed:   0.0,
		TextStrokeWidth:      0.0,
		WordWrap:             1.0,
		HorizontalAlignment:  string(node.HorizontalAlignment),
		VerticalAlignment:    string(node.VerticalAlignment),
		ResizeBehavior:       "resizeText",
		Text:                 node._text,
		Font: rawFont{
			ColorRed:   keepZero(fontR),
			ColorGreen: keepZero(fontG),
			ColorBlue:  keepZero(fontB),
			Size:       keepZero(node.Font.Size),
			Tracking:   keepZero(node.Font.Tracking),
			Name:       node.Font.Name,
			Weight:     string(node.Font.Weight),
		},
		TextAttributes: rawTextAttributes{
			Type:      "textAttributeList",
			Keyframes: []Keyframe[[]TextAttribute]{{}},
		},
	}

	if node.Shape != CalloutShapeText {
		if v, ok := node.FillColor.(rawColor); ok {
			def.FillColorRed = ref(keepZero(v.r))
			def.FillColorGreen = ref(keepZero(v.g))
			def.FillColorBlue = ref(keepZero(v.b))
		} else {
			r, g, b, _ := colorTo1Scale(node.FillColor)
			def.FillColorRed = ref(keepZero(r))
			def.FillColorGreen = ref(keepZero(g))
			def.FillColorBlue = ref(keepZero(b))
		}

		if v, ok := node.StrokeColor.(rawColor); ok {
			def.StrokeColorRed = ref(keepZero(v.r))
			def.StrokeColorGreen = ref(keepZero(v.g))
			def.StrokeColorBlue = ref(keepZero(v.b))
		} else {
			r, g, b, _ := colorTo1Scale(node.StrokeColor)
			def.StrokeColorRed = ref(keepZero(r))
			def.StrokeColorGreen = ref(keepZero(g))
			def.StrokeColorBlue = ref(keepZero(b))
		}

		def.FillStyle = (*string)(&node.FillStyle)
		def.FillColorOpacity = ref(keepZero(node.FillOpacity))
		def.StrokeColorOpacity = ref(keepZero(node.StrokeOpacity))
		def.StrokeWidth = (*keepZero)(&node.StrokeWidth)
		def.StrokeStyle = (*string)(&node.StrokeStyle)
		def.TailX = (*keepZero)(&node.TailX)
		def.TailY = (*keepZero)(&node.TailY)
	}

	rMedia.Def = def
	rMedia.Attributes.AutoRotateText = ref(true)
}

func (c *Callout) node()           {}
func (c *Callout) camType() string { return "Callout" }
