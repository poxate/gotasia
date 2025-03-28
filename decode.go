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
		track := p.NewTrack(attributes.Ident)

		var start time.Duration
		for _, media := range rawTrack.Medias {
			track.Elements = append(track.Elements, &Element{
				gap:          p.decodeTime(media.Start) - start,
				duration:     p.decodeTime(media.Duration),
				Animations:   []*Animation{},
				node:         dec.decodeNode(media),
				_rawMetadata: media.Metadata,
			})
		}
	}

	return p, nil
}

func (p *Project) decodeTime(rawtime int) time.Duration {
	return time.Duration(float64(rawtime) / float64(p.editRate) * float64(time.Second))
}

func (dec *Decoder) rawToDuration(rawtime int) time.Duration {
	sec := float64(rawtime) / float64(dec.editRate)
	return time.Duration(sec * float64(editRate))
}

func (dec *Decoder) decodeNode(media rawMedia) Node {
	switch media.Type {
	case "Callout":
		c := NewCallout().SetText(media.Def.Text)
		c.Shape = CalloutShape(media.Def.Shape)
		c.Width = float64(media.Def.Width)
		c.Height = float64(media.Def.Height)

		c.Font = Font{
			Color: rawColor{
				r: float64(media.Def.Font.ColorRed),
				g: float64(media.Def.Font.ColorBlue),
				b: float64(media.Def.Font.ColorGreen),
			},
			Size:     float64(media.Def.Font.Size),
			Tracking: float64(media.Def.Font.Tracking),
			Name:     media.Def.Font.Name,
			Weight:   FontWeight(media.Def.Font.Weight),
		}

		if c.Shape != CalloutShapeText {
			c.FillColor = rawColor{
				r: float64(*media.Def.FillColorRed),
				g: float64(*media.Def.FillColorGreen),
				b: float64(*media.Def.FillColorBlue),
			}
			c.StrokeColor = rawColor{
				r: float64(*media.Def.StrokeColorRed),
				g: float64(*media.Def.StrokeColorGreen),
				b: float64(*media.Def.StrokeColorBlue),
			}
			c.FillStyle = CalloutFillStyle(*media.Def.FillStyle)
			c.FillOpacity = float64(*media.Def.FillColorOpacity)
			c.StrokeOpacity = float64(*media.Def.StrokeColorOpacity)
			c.StrokeWidth = float64(*media.Def.StrokeWidth)
			c.StrokeStyle = CalloutStrokeStyle(*media.Def.StrokeStyle)
			c.TailX = float64(*media.Def.TailX)
			c.TailY = float64(*media.Def.TailY)
		}

		return c
	default:
		return NewCallout().SetText("[unhandled media type]: " + media.Type)
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
