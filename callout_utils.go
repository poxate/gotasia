package gotasia

import "image/color"

func (s CalloutFillStyle) String() string {
	switch s {
	case CalloutFillSolid:
		return "solid"
	case CalloutFillGradient:
		return "gradient"
	default:
		return "solid"
	}
}

func calloutFillStyleFrom(value string) CalloutFillStyle {
	switch value {
	case "solid":
		return CalloutFillSolid
	case "gradient":
		return CalloutFillGradient
	default:
		return CalloutFillSolid
	}
}

func (a CalloutHorizontalAlignment) string() string {
	switch a {
	case CalloutHorizontalAlignmentLeft:
		return "left"
	case CalloutHorizontalAlignmentCenter:
		return "center"
	case CalloutHorizontalAlignmentRight:
		return "right"
	default:
		return "center"
	}
}

func (a CalloutVerticalAlignment) string() string {
	switch a {
	case CalloutVerticalAlignmentTop:
		return "top"
	case CalloutVerticalAlignmentCenter:
		return "center"
	case CalloutVerticalAlignmentBottom:
		return "bottom"
	default:
		return "center"
	}
}

func (a CalloutShape) string() string {
	switch a {
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
		return "text"
	}
}

func calloutShapeFrom(value string) CalloutShape {
	switch value {
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
		return CalloutShapeText
	}
}

func fontWeightFrom(value string) FontWeight {
	switch value {
	case "regular":
		return FontWeightRegular
	case "bold":
		return FontWeightBold
	default:
		return FontWeightRegular
	}
}

func (f Font) withDefaults() Font {
	if f.Size == 0 {
		f.Size = 96
	}

	if f.Name == "" {
		f.Name = "Montserrat"
	}

	if f.Color == nil {
		f.Color = color.White
	}

	return f
}

func (w FontWeight) string() string {
	switch w {
	case FontWeightRegular:
		return "Regular"
	case FontWeightBold:
		return "Bold"
	default:
		return "Regular"
	}
}
