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
	rawP := rawProjectD{}
	if err := json.NewDecoder(dec.src).Decode(&rawP); err != nil {
		return nil, fmt.Errorf("failed to decode project: %w", err)
	}

	dec.editRate = rawP.EditRate
	p._width = int(rawP.Width)
	p._height = int(rawP.Height)
	p._autoNormalizeLoudness = rawP.ShouldApplyLoudnessNormalization
	p._frameRate = FrameRate(rawP.VideoFormatFrameRate)
	p._backgroundColor = color.NRGBA{
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
				gap:        dec.rawToDuration(rawMedia.Start) - now,
				duration:   dec.rawToDuration(rawMedia.Duration),
				scaleX:     rawMedia.Parameters.Scale0.getFirstValue(),
				scaleY:     rawMedia.Parameters.Scale1.getFirstValue(),
				node:       dec.decodeNode(&rawMedia),
				Animations: dec.decodeAnimations(&rawMedia),
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
			Text:             media.Def.Text,
			Spans:            spans,
			Shape:            calloutShapeFrom(media.Def.Shape),
			Width:            int(media.Def.Width.getFirstValue()),
			Height:           int(media.Def.Height.getFirstValue()),
			TextFontSize:     media.Def.Font.Size,
			TextFontName:     media.Def.Font.Name,
			TextFontWeight:   media.Def.Font.Weight,
			TextFontTracking: media.Def.Font.Tracking,
			TextFontColor: color.NRGBA{
				R: uint8(media.Def.Font.ColorRed * 255),
				G: uint8(media.Def.Font.ColorGreen * 255),
				B: uint8(media.Def.Font.ColorBlue * 255),
				A: 255,
			},
		}
	default:
		return &Callout{
			Text:   "",
			Width:  200,
			Height: 200,
		}
	}
}

func (dec *Decoder) decodeAnimations(rawMedia *rawMedia) []*Animation {
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
