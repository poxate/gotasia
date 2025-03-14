package gotasia

import (
	"encoding/json"
	"fmt"
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
	rawProject := rawProject{}
	if err := json.NewDecoder(dec.src).Decode(&rawProject); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	} else if rawProject.Version != "8.0" {
		return nil, fmt.Errorf("unsupported version: %s", rawProject.Version)
	}

	p.FrameRate(FrameRate(rawProject.VideoFormatFrameRate))
	p._width = int(rawProject.Width)
	p._height = int(rawProject.Height)
	p.editRate = rawProject.EditRate

	for trackIndex, rawTrack := range rawProject.Timeline.SceneTrack.Scenes[0].Csml.Tracks {
		attributes := rawProject.Timeline.TrackAttributes[trackIndex]
		p.Tracks = append(p.Tracks, NewTrack(attributes.Ident))
		_ = rawTrack
	}

	return p, nil
}

func (dec *Decoder) rawToDuration(rawtime int) time.Duration {
	sec := float64(rawtime) / float64(dec.editRate)
	return time.Duration(sec * float64(editRate))
}

func (dec *Decoder) decodeNode(media *rawMedia_old) Node {
	switch media.Type {
	case "Callout":
		if len(media.Def.TextAttributes.Keyframes) != 1 {
			panic("expected media's text attribute to have 1 keyframe")
		}

		spans := decodeSpans(media.Def.Text, media.Def.TextAttributes)

		return &Callout{
			_text:  media.Def.Text,
			Spans:  spans,
			Shape:  calloutShapeFrom(media.Def.Shape),
			Width:  int(media.Def.Width.getFirstValue()),
			Height: int(media.Def.Height.getFirstValue()),
			// TextFontSize:     media.Def.Font.Size,
			// TextFontName:     media.Def.Font.Name,
			// TextFontWeight:   media.Def.Font.Weight,
			// TextFontTracking: media.Def.Font.Tracking,
			// TextFontColor: color.NRGBA{
			// 	R: uint8(media.Def.Font.ColorRed * 255),
			// 	G: uint8(media.Def.Font.ColorGreen * 255),
			// 	B: uint8(media.Def.Font.ColorBlue * 255),
			// 	A: 255,
			// },
		}
	default:
		return &Callout{
			_text:  "",
			Width:  200,
			Height: 200,
		}
	}
}

func (dec *Decoder) decodeAnimations(rawMedia *rawMedia_old) []*Animation {
	animations := []*Animation{}
	animationStarts := map[int]*Animation{}
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.Translation0)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.Translation1)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.Translation2)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.Rotation1)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.Shear1)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.Scale0)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.Scale1)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.GeometryCrop0)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.GeometryCrop1)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.GeometryCrop2)
	loadAnimations(dec, &animations, animationStarts, rawMedia.Parameters.GeometryCrop3)

	for _, keyframe := range rawMedia.Parameters.Scale0.Keyframe.Keyframes {
		animation := animationStarts[keyframe.Time]
		animation.ToScaleX(keyframe.Value)
	}

	for _, keyframe := range rawMedia.Parameters.Scale1.Keyframe.Keyframes {
		animation := animationStarts[keyframe.Time]
		animation.ToScaleY(keyframe.Value)
	}

	return animations
}

func loadAnimations[T any](dec *Decoder, animations *[]*Animation, animationStarts map[int]*Animation, maybeKeyframes rawMaybeKeyframes[T]) {
	if maybeKeyframes.Static {
		return
	}

	var start time.Duration = 0
	for _, keyframe := range maybeKeyframes.Keyframe.Keyframes {
		if v, ok := animationStarts[keyframe.Time]; ok {
			if v.Duration != dec.rawToDuration(keyframe.Duration) {
				panic("overlaying keyframes with different durations")
			} else if v.Duration != dec.rawToDuration(keyframe.EndTime-keyframe.Time) {
				panic("overlaying keyframes with inconsistent time or endtime")
			}
			continue
		}

		animation := NewAnimation(dec.rawToDuration(keyframe.Time)-start, dec.rawToDuration(keyframe.Duration))
		*animations = append(*animations, animation)
		animationStarts[keyframe.Time] = animation
	}
}
