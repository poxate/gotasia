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
				Node:         dec.decodeNode(media),
				_rawMetadata: media.Metadata,
			})
		}
	}

	return p, nil
}

func (p *Project) decodeTime(rawtime int) time.Duration {
	return time.Duration(float64(rawtime) / float64(p.editRate) * float64(time.Second))
}

func (dec *Decoder) decodeNode(media rawMedia) Node {
	switch media.Type {
	case "Callout":
		c := &Callout{
			Text:   media.Def.Text,
			Shape:  calloutShapeFrom(media.Def.Shape),
			Width:  float64(media.Def.Width),
			Height: float64(media.Def.Height),
			Font: Font{
				Color: rawColor{
					r: float64(media.Def.Font.ColorRed),
					g: float64(media.Def.Font.ColorBlue),
					b: float64(media.Def.Font.ColorGreen),
				},
				Size:     float64(media.Def.Font.Size),
				Tracking: float64(media.Def.Font.Tracking),
				Name:     media.Def.Font.Name,
				Weight:   fontWeightFrom(media.Def.Font.Weight),
			},
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
			c.FillStyle = calloutFillStyleFrom(*media.Def.FillStyle)
			c.FillOpacity = float64(*media.Def.FillColorOpacity)
			c.StrokeOpacity = float64(*media.Def.StrokeColorOpacity)
			c.StrokeWidth = float64(*media.Def.StrokeWidth)
			c.StrokeStyle = CalloutStrokeStyle(*media.Def.StrokeStyle)
			c.TailX = float64(*media.Def.TailX)
			c.TailY = float64(*media.Def.TailY)
		}

		return c
	default:
		return &Callout{
			Text: "[unhandled media type]: " + media.Type,
		}
	}
}
