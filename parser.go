package gotasia

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
)

func Parse(src io.Reader) (*Project, error) {
	p := NewProject(0, 0)
	rawP := rawProject{}
	if err := json.NewDecoder(src).Decode(&rawP); err != nil {
		return nil, fmt.Errorf("failed to decode project: %w", err)
	}

	p.Width = int(rawP.Width)
	p.Height = int(rawP.Height)
	p.AutoNormalizeLoudness = rawP.ShouldApplyLoudnessNormalization
	p.FrameRate = FrameRate(rawP.VideoFormatFrameRate)
	p.BackgroundColor = color.NRGBA{
		R: uint8(rawP.Timeline.BackgroundColor[0]),
		G: uint8(rawP.Timeline.BackgroundColor[1]),
		B: uint8(rawP.Timeline.BackgroundColor[2]),
		A: uint8(rawP.Timeline.BackgroundColor[3]),
	}

	return p, nil
}
