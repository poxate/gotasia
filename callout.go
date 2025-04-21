package gotasia

import (
	"image/color"
)

type CalloutShape int
type CalloutFillStyle int
type CalloutStrokeStyle string
type FontWeight int
type CalloutHorizontalAlignment int
type CalloutVerticalAlignment int

const (
	CalloutShapeText CalloutShape = iota
	CalloutShapeSpeechBubble1
	CalloutShapeSpeechBubble2
	CalloutShapeThoughtBubble1
	CalloutShapeThoughtBubble2
	CalloutShapeTextArrow1
	CalloutShapeTextArrow2
	CalloutShapeTextRectangle
)

const (
	CalloutFillSolid CalloutFillStyle = iota
	CalloutFillGradient
)

const (
	CalloutStrokeSolid      CalloutStrokeStyle = "solid"
	CalloutStrokeDot        CalloutStrokeStyle = "dot"
	CalloutStrokeDash       CalloutStrokeStyle = "dash"
	CalloutStrokeDashDot    CalloutStrokeStyle = "dashdot"
	CalloutStrokeDashDotDot CalloutStrokeStyle = "dashdotdot"
)

const (
	FontWeightRegular FontWeight = iota
	FontWeightBold
)

const (
	CalloutHorizontalAlignmentCenter CalloutHorizontalAlignment = 0
	CalloutHorizontalAlignmentLeft   CalloutHorizontalAlignment = 1
	CalloutHorizontalAlignmentRight  CalloutHorizontalAlignment = 2

	CalloutVerticalAlignmentCenter CalloutVerticalAlignment = 0
	CalloutVerticalAlignmentTop    CalloutVerticalAlignment = 1
	CalloutVerticalAlignmentBottom CalloutVerticalAlignment = 2
)

type Callout struct {
	Text string // ignored if Spans != nil
	Font Font

	Shape CalloutShape

	// if FillStyle is not specified, will render to solid
	FillStyle   CalloutFillStyle
	FillColor   color.Color
	FillOpacity float64

	StrokeColor   color.Color
	StrokeOpacity float64
	StrokeWidth   float64
	StrokeStyle   CalloutStrokeStyle

	Width        float64 // Must be between 1 and 9999. If 0, it is rendered as 400px
	Height       float64 // Must be between 1 and 9999. If 0, it is rendered as 250px
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

func (c *Callout) width() int  { return int(c.Width) }
func (c *Callout) height() int { return int(c.Height) }

func (node *Callout) encodeIntoMedia(rMedia *rawMedia) {
	font := node.Font.withDefaults()
	var fontR, fontG, fontB, _ float64 = colorTo1Scale(font.Color)

	width := node.Width
	if width == 0 {
		width = 400
	}

	height := node.Height
	if height == 0 {
		height = 250
	}

	cornerRadius := float64(node.CornerRadius)
	if node.Shape == CalloutShapeText {
		cornerRadius = 8
	}

	def := &rawDef{
		Kind:                 "remix",
		Shape:                node.Shape.string(),
		Style:                "basic",
		Width:                keepZero(width),
		Height:               keepZero(height),
		CornerRadius:         keepZero(cornerRadius),
		EnableLigatures:      1.0,
		LineSpacing:          0.0,
		TextStrokeAlignment:  2.0,
		TextStrokeColorAlpha: 1.0,
		TextStrokeColorBlue:  0.0,
		TextStrokeColorGreen: 0.0,
		TextStrokeColorRed:   0.0,
		TextStrokeWidth:      0.0,
		WordWrap:             1.0,
		HorizontalAlignment:  node.HorizontalAlignment.string(),
		VerticalAlignment:    node.VerticalAlignment.string(),
		ResizeBehavior:       "resizeText",
		Text:                 node.Text,
		Font: rawFont{
			ColorRed:   keepZero(fontR),
			ColorGreen: keepZero(fontG),
			ColorBlue:  keepZero(fontB),
			Size:       keepZero(font.Size),
			Tracking:   keepZero(font.Tracking),
			Name:       font.Name,
			Weight:     font.Weight.string(),
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

		def.FillStyle = ref(node.FillStyle.String())
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
