package gotasia

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"time"
)

type Decoder struct {
	src      io.Reader
	editRate int
}

func NewDecoder(src io.Reader) *Decoder {
	return &Decoder{src: src}
}

func (dec *Decoder) Decode() (*Project, error) {
	p := NewProject(0, 0)
	rawP := rawProject{}
	if err := json.NewDecoder(dec.src).Decode(&rawP); err != nil {
		return nil, fmt.Errorf("failed to decode project: %w", err)
	}

	dec.editRate = rawP.EditRate
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

	if len(rawP.Timeline.TrackAttributes) != len(rawP.Timeline.SceneTrack.Scenes[0].Csml.Tracks) {
		return nil, fmt.Errorf("this project's format is invalid, inequal number of tracks (%d) to track attributes (%d)", len(rawP.Timeline.SceneTrack.Scenes[0].Csml.Tracks), len(rawP.Timeline.TrackAttributes))
	}

	for i, rawTrack := range rawP.Timeline.SceneTrack.Scenes[0].Csml.Tracks {
		attributes := rawP.Timeline.TrackAttributes[i]
		track := NewTrack(attributes.Ident)

		var now time.Duration
		for _, rawMedia := range rawTrack.Medias {
			track.Elements = append(track.Elements, &Element{
				gap:      dec.rawToDuration(rawMedia.Start) - now,
				duration: dec.rawToDuration(rawMedia.Duration),
				scale:    1,
				node:     dec.decodeNode(&rawMedia),
			})
		}

		p.Tracks = append(p.Tracks, track)
	}

	return p, nil
}

func (dec *Decoder) rawToDuration(rawtime int) time.Duration {
	sec := float64(rawtime) / float64(dec.editRate)
	return time.Duration(sec * float64(editRate))
}

func (dec *Decoder) decodeNode(media *rawMedia) Node {
	switch media.Type {
	case "Callout":
		if len(media.Def.TextAttributes.Keyframes) != 1 {
			panic("expected media's text attribute to have 1 keyframe")
		}

		spans := decodeSpans(media.Def.Text, media.Def.TextAttributes)

		return &Callout{
			Text:   media.Def.Text,
			Spans:  spans,
			Shape:  calloutShapeFrom(media.Def.Shape),
			Width:  int(media.Def.Width.getFirstValue()),
			Height: int(media.Def.Height.getFirstValue()),
		}
	default:
		return &Callout{
			Text:   "",
			Width:  200,
			Height: 200,
		}
	}
}
