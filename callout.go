package gotasia

import (
	"fmt"
	"image/color"
	"maps"
	"slices"
)

type CalloutShape string
type CalloutFillStyle string
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
	Spans []Span

	Shape CalloutShape
	// if FillStyle is not specified, will render to solid
	FillStyle    CalloutFillStyle
	FillColor    color.Color
	Width        int
	Height       int
	TailX        int // distance from the center of the callout along the x-axis, only used for speech and thought bubbles
	TailY        int // distance from the center of the callout along the y-axis, only used for speech and thought bubbles
	CornerRadius int // only used for text-rectangle

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

func calloutShapeFrom(raw string) CalloutShape {
	switch raw {
	case "text":
		return CalloutShapeText
	case "speech-bubble":
		return CalloutShapeSpeechBubble1
	case "speech-bubble2":
		return CalloutShapeSpeechBubble2
	case "thought-bubble":
		return CalloutShapeThoughtBubble1
	case "thought-bubble2":
		return CalloutShapeThoughtBubble2
	case "text-arrow":
		return CalloutShapeTextArrow1
	case "text-arrow2":
		return CalloutShapeTextArrow2
	case "text-rectangle":
		return CalloutShapeTextRectangle
	default:
		panic("unknown callout shape: " + raw)
	}
}

func (shape CalloutShape) string() string {
	switch shape {
	case CalloutShapeText:
		return "text"
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
		panic("unknown callout shape representation: " + shape)
	}
}

func (c *Callout) width() int  { return c.Width }
func (c *Callout) height() int { return c.Height }

func (node *Callout) encodeIntoMedia(rMedia *rawMedia) {
	fR, fG, fB, _ := colorTo1Scale(node.Font.Color)

	rMedia.Def = &rawDef{
		Kind:                 "remix",
		Shape:                string(node.Shape),
		Style:                "basic",
		Width:                KeepZero(node.Width),
		Height:               KeepZero(node.Height),
		CornerRadius:         KeepZero(node.CornerRadius),
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
			ColorRed:   KeepZero(fR),
			ColorGreen: KeepZero(fG),
			ColorBlue:  KeepZero(fB),
			Size:       KeepZero(node.Font.Size),
			Tracking:   KeepZero(node.Font.Tracking),
			Name:       node.Font.Name,
			Weight:     string(node.Font.Weight),
		},
	}
}

func (node *Callout) encodeDef() jobj {
	fillStyle := node.FillStyle
	if fillStyle == "" {
		fillStyle = "solid"
	}

	fillR, fillG, fillB, fillA := colorTo1Scale(node.FillColor)
	fontR, fontG, fontB, _ := colorTo1Scale(nil)

	text := node._text

	keyframeValues := []jobj{}
	if len(node.Spans) > 0 {
		for _, t := range node.Spans {
			text += t.Text
		}

		start := 0
		for _, span := range node.Spans {
			r, g, b, a := colorTo255(span.Color)

			gen := func(name string, valueType string, value interface{}) jobj {
				return jobj{
					"name":       name,
					"value":      value,
					"valueType":  valueType,
					"rangeStart": start,
					"rangeEnd":   start + len(span.Text),
				}
			}

			keyframeValues = append(keyframeValues,
				gen("fontSize", "double", float64(span.FontSize)),
				gen("underline", "int", boolToInt(span.Underline)),
				gen("fgColor", "color", fmt.Sprintf("(%d,%d,%d,%d)", r, g, b, a)),
				gen("fontWeight", "int", span.Weight),
				gen("fontItalic", "int", boolToInt(span.Italic)),
				gen("fontName", "string", span.FontName),
				gen("kerning", "double", span.Kerning),
				gen("strikethrough", "int", boolToInt(span.Strikethrough)),
			)
			start += len(span.Text)
		}
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
		"fill-color-red":          fillR,
		"fill-color-green":        fillG,
		"fill-color-blue":         fillB,
		"fill-color-opacity":      fillA,
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
		"text":                    text,
		"vertical-alignment":      "center",
		"font": jobj{
			"color-blue":  fontB,
			"color-green": fontG,
			"color-red":   fontR,
			// "size":        node.TextFontSize,
			// "tracking":    node.TextFontTracking,
			// "name":        node.TextFontName,
			// "weight":      node.TextFontWeight,
		},
		"textAttributes": jobj{
			"type": "textAttributeList",
			"keyframes": []jobj{
				{
					"endTime":  0,
					"time":     0,
					"value":    keyframeValues,
					"duration": 0,
				},
			},
		},
	}
}

func decodeSpans(text string, attributes rawTextAttributes) []Span {
	spans := flattenRange(text, attributes)

	start := 0
	for i := range spans {
		span := &spans[i]
		for _, detail := range attributes.Keyframes[0].Value {
			if inRange := detail.RangeStart <= start && detail.RangeEnd >= start+len(span.Text); !inRange {
				continue
			}
			switch detail.Name {
			case "fontSize":
				span.FontSize = int(detail.Value.(float64))
			case "underline":
				span.Underline = detail.Value.(float64) != 0
			case "fgColor":
				colorValue := detail.Value.(string)
				var r, g, b, a uint8
				fmt.Sscanf(colorValue, "(%d,%d,%d,%d)", &r, &g, &b, &a)
				span.Color = color.RGBA{R: r, G: g, B: b, A: a}
			case "fontWeight":
				span.Weight = int(detail.Value.(float64))
			case "fontItalic":
				span.Italic = detail.Value.(float64) != 0
			case "fontName":
				span.FontName = detail.Value.(string)
			case "kerning":
				span.Kerning = detail.Value.(float64)
			case "strikethrough":
				span.Strikethrough = detail.Value.(float64) != 0
			}
		}
		start += len(span.Text)
	}

	return spans
}

func flattenRange(text string, attributes rawTextAttributes) []Span {
	spanStarts := map[int]struct{}{}

	for _, value := range attributes.Keyframes[0].Value {
		spanStarts[value.RangeStart] = struct{}{}
	}

	starts := slices.Sorted(maps.Keys(spanStarts))

	spans := []Span{}
	for i, start := range starts {
		end := len(text)
		if i < len(starts)-1 {
			end = starts[i+1]
		}

		spans = append(spans, Span{
			Text: text[start:end],
		})
	}

	return spans
}

func (c *Callout) node()           {}
func (c *Callout) camType() string { return "Callout" }
