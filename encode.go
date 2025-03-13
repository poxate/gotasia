package gotasia

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/sanity-io/litter"
)

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

	bgR, bgG, bgB, bgA := colorTo255(p.BackgroundColor)

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

		width := int(element.scaleX * float64(element.node.width()))
		height := int(element.scaleY * float64(element.node.height()))

		translateX := rawMaybeKeyframes[int]{Static: true, StaticValue: 0, Keyframe: rawKeyframes[int]{Type: "double"}}
		if element.xSet {
			translateX.StaticValue = p.coordX(int(width), element.x)
		}

		translateY := rawMaybeKeyframes[int]{Static: true, StaticValue: 0, Keyframe: rawKeyframes[int]{Type: "double"}}
		if element.ySet {
			translateY.StaticValue = p.coordY(int(height), element.y)
		}

		scaleX := rawMaybeKeyframes[float64]{Static: true, StaticValue: element.scaleX, Keyframe: rawKeyframes[float64]{Type: "double"}}
		scaleY := rawMaybeKeyframes[float64]{Static: true, StaticValue: element.scaleY, Keyframe: rawKeyframes[float64]{Type: "double"}}

		var animateStart time.Duration = start
		lastScaleX := scaleX.StaticValue
		lastScaleY := scaleY.StaticValue
		for _, animation := range element.Animations {
			if animation.scaleX != nil {
				if scaleX.Static {
					scaleX.Static = false
					scaleX.Keyframe.DefaultValue = scaleX.StaticValue
				}

				scaleX.Keyframe.Keyframes = append(scaleX.Keyframe.Keyframes,
					keyframe[float64]{
						Time:     p.encodeTime(animateStart + animation.Gap),
						EndTime:  p.encodeTime(animateStart + animation.Gap + animation.Duration),
						Duration: p.encodeTime(animation.Duration),
						Value:    *animation.scaleX,
					},
				)
				lastScaleX = *animation.scaleX
			}

			if animation.scaleY != nil {
				if scaleY.Static {
					scaleY.Static = false
					scaleY.Keyframe.DefaultValue = scaleY.StaticValue
				}

				scaleY.Keyframe.Keyframes = append(scaleY.Keyframe.Keyframes,
					keyframe[float64]{
						Time:     p.encodeTime(animateStart + animation.Gap),
						EndTime:  p.encodeTime(animateStart + animation.Gap + animation.Duration),
						Duration: p.encodeTime(animation.Duration),
						Value:    *animation.scaleY,
					},
				)
				lastScaleY = *animation.scaleY
			}

			if animation.x != nil {
				if translateX.Static {
					translateX.Static = false
					translateX.Keyframe.DefaultValue = translateX.StaticValue
				}

				translateX.Keyframe.Keyframes = append(translateX.Keyframe.Keyframes,
					keyframe[int]{
						Time:     p.encodeTime(animateStart + animation.Gap),
						EndTime:  p.encodeTime(animateStart + animation.Gap + animation.Duration),
						Duration: p.encodeTime(animation.Duration),
						Value:    p.coordX(int(lastScaleX*float64(element.node.width())), *animation.x),
					},
				)
			}

			if animation.y != nil {
				if translateY.Static {
					translateY.Static = false
					translateY.Keyframe.DefaultValue = translateY.StaticValue
				}

				translateY.Keyframe.Keyframes = append(translateY.Keyframe.Keyframes,
					keyframe[int]{
						Time:     p.encodeTime(animateStart + animation.Gap),
						EndTime:  p.encodeTime(animateStart + animation.Gap + animation.Duration),
						Duration: p.encodeTime(animation.Duration),
						Value:    p.coordY(int(lastScaleY*float64(element.node.height())), *animation.y),
					},
				)
			}
		}

		obj := jobj{
			"id":            p.id.gen(),
			"start":         int(start.Seconds() * editRate),
			"duration":      int(element.duration.Seconds() * editRate),
			"mediaStart":    0,
			"mediaDuration": int(element.duration.Seconds() * editRate),
			"parameters": jobj{
				"scale0":       scaleX,
				"scale1":       scaleY,
				"translation0": translateX,
				"translation1": translateY,
			},
		}

		switch node := element.node.(type) {
		case *Callout:
			obj["_type"] = "Callout"
			obj["def"] = node.encodeDef()
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

func (p *Project) encodeTime(dur time.Duration) int {
	return int(dur.Seconds() * editRate)
}
