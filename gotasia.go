package gotasia

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"path/filepath"
	"time"

	"github.com/sanity-io/litter"
)

type jobj = map[string]any

type FrameRate int

const (
	FrameRate25 FrameRate = 25
	FrameRate30 FrameRate = 30
	FrameRate50 FrameRate = 50
	FrameRate60 FrameRate = 60
)

const editRate = 705600000

type Project struct {
	id                    *idTracker
	BackgroundColor       color.Color
	Width                 int
	Height                int
	AutoNormalizeLoudness bool
	FrameRate             FrameRate
	Tracks                []*Track
	MediaBin              MediaBin
	MediaItemId           map[*MediaItem]int
}

func NewProject(width, height int) *Project {
	return &Project{
		id:                    &idTracker{},
		Width:                 width,
		Height:                height,
		AutoNormalizeLoudness: true,
		FrameRate:             FrameRate60,
		Tracks:                []*Track{},
		MediaBin:              MediaBin{},
		MediaItemId:           map[*MediaItem]int{},
	}
}

func (p *Project) width() int {
	return p.Width
}

func (p *Project) height() int {
	return p.Height
}

func (p *Project) Encode(w io.Writer) error {
	if len(p.Tracks) == 0 {
		return fmt.Errorf("cannot encode project without at least 1 track")
	}

	sourcebin := p.encodeSourcebin()

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(jobj{
		"title":                            "",
		"description":                      "",
		"author":                           "",
		"targetLoudness":                   -18.0,
		"shouldApplyLoudnessNormalization": p.AutoNormalizeLoudness,
		"videoFormatFrameRate":             p.FrameRate,
		"audioFormatSampleRate":            44100,
		"allowSubFrameEditing":             false,
		"width":                            p.Width,
		"height":                           p.Height,
		"version":                          "8.0",
		"editRate":                         editRate,
		"sourceBin":                        sourcebin,
		"timeline":                         p.encodeTimeline(),
		"authoringClientName": jobj{
			"name":     "Camtasia",
			"platform": "Windows",
			"version":  "2024.0.1",
		},
		"metadata": jobj{
			"audioNarrationNotes": "",
			"calloutStyle":        "Basic",
		},
	})
}

func (p *Project) encodeTimeline() jobj {
	timelineId := p.id.gen()

	trackObjs := []jobj{}
	trackAttributeObjs := []jobj{}

	for i, track := range p.Tracks {
		trackObjs = append(trackObjs, jobj{
			"trackIndex": i,
			"medias":     p.encodeTrackMedias(track),
		})

		trackAttributeObjs = append(trackAttributeObjs, jobj{
			"ident": track.Name,
		})
	}

	var bgR, bgG, bgB, bgA uint32
	if p.BackgroundColor != nil {
		bgR, bgG, bgB, bgA = p.BackgroundColor.RGBA()
	}
	bgR /= 257
	bgG /= 257
	bgB /= 257
	bgA /= 257

	return jobj{
		"id": timelineId,
		"sceneTrack": jobj{
			"scenes": []jobj{{
				"csml": jobj{
					"tracks": trackObjs,
				},
			}},
		},
		"trackAttributes": trackAttributeObjs,
		"backgroundColor": []uint32{bgR, bgG, bgB, bgA},
	}
}

func (p *Project) encodeTrackMedias(track *Track) []jobj {
	list := []jobj{}

	var start time.Duration = 0
	for _, element := range track.Elements {
		start += element.gap

		width := int(element.scale * float64(element.node.width()))
		height := int(element.scale * float64(element.node.height()))

		var translateX int
		if element.xSet {
			translateX = p.coordX(int(width), element.x)
		}

		var translateY int
		if element.ySet {
			translateY = p.coordY(int(height), element.y)
		}

		obj := jobj{
			"id":            p.id.gen(),
			"start":         int(start.Seconds() * editRate),
			"duration":      int(element.duration.Seconds() * editRate),
			"mediaStart":    0,
			"mediaDuration": int(element.duration.Seconds() * editRate),
			"parameters": jobj{
				"scale0":       element.scale,
				"scale1":       element.scale,
				"translation0": translateX,
				"translation1": translateY,
			},
		}

		switch node := element.node.(type) {
		case *Callout:
			obj["_type"] = "Callout"
			obj["def"] = node.createDef()
		case *ImageFile:
			obj["_type"] = "IMFile"
			obj["src"] = p.MediaItemId[node.Src]
		default:
			panic("unknown element type: " + litter.Sdump(element.node))
		}

		list = append(list, obj)
		start += element.duration + 1
	}

	return list
}

func (p *Project) encodeSourcebin() []jobj {
	list := []jobj{}

	for _, item := range p.MediaBin {
		if item.Type == ImageMediaItem {
			id := p.id.gen()
			p.MediaItemId[item] = id
			list = append(list, jobj{
				"id":   id,
				"src":  item.Src,
				"rect": []int{0, 0, item.Width, item.Height},
				"sourceTracks": []jobj{
					{
						"range":          []int{0, 1},
						"type":           1,
						"editRate":       10000000,
						"trackRect":      []int{0, 0, item.Width, item.Height},
						"sampleRate":     0,
						"bitDepth":       24,
						"numChannels":    0,
						"integratedLUFS": 100.0,
						"peakLevel":      -1.0,
						"metaData":       filepath.Base(item.Src),
					},
				},
			})
		}
	}

	return list
}

func (p *Project) coordX(width, posX int) int {
	pCenter := p.Width / 2
	eCenter := posX + (width / 2)
	return eCenter - pCenter
}

func (p *Project) coordY(height, posY int) int {
	pCenter := p.Height / 2
	eCenter := posY + (height / 2)
	return pCenter - eCenter
}
