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

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(rawProject{
		Title:                            "",
		Description:                      "",
		Author:                           "",
		TargetLoudness:                   keepZero(p._targetLoudness),
		ShouldApplyLoudnessNormalization: p._autoNormalizeLoudness,
		VideoFormatFrameRate:             int(p._frameRate),
		AudioFormatSampleRate:            44100,
		AllowSubFrameEditing:             false,
		Width:                            keepZero(p._width),
		Height:                           keepZero(p._height),
		Version:                          "8.0",
		EditRate:                         editRate,
		AuthoringClientName: rawAuthoringClientName{
			Name:     "Camtasia",
			Platform: "Windows",
			Version:  "2024.0.1",
		},
		Timeline: rawTimeline{
			ID: p.id.gen(),
			SceneTrack: rawSceneTrack{
				Scenes: []rawScene{
					{
						Csml: rawCsml{
							Tracks: p.encodeRawTracks(),
						},
					},
				},
			},
			TrackAttributes: p.encodeTrackAttributes(),
			CaptionAttributes: rawCaptionAttributes{
				Enabled:                  true,
				FontName:                 "Arial",
				FontSize:                 64,
				BackgroundColor:          []int{0, 0, 0, 191},
				ForegroundColor:          []int{255, 255, 255, 255},
				Lang:                     "en",
				Alignment:                0,
				DefaultFontSize:          true,
				Opacity:                  0.5,
				BackgroundEnabled:        true,
				BackgroundOnlyAroundText: true,
			},
			Gain:                    1.0,
			LegacyAttenuateAudioMix: false,
			BackgroundColor:         []int{0, 0, 0, 255},
		},
		Metadata: rawMetadata{
			AutoSaveFile:        "",
			CanvasZoom:          rawCanvasZoom{Type: "int", Value: 60},
			Date:                "2025-03-13 08:38:36 AM",
			IsAutoSave:          rawIsAutoSave{Type: "bool", Value: false},
			IsStandalone:        rawIsAutoSave{Type: "bool", Value: true},
			Language:            "ENU",
			ProfileName:         "",
			ProjectThumbnail:    "iVBORw0KGgoAAAANSUhEUgAAAoAAAAFoCAYAAADHMkpRAAAAAXNSR0IArs4c6QAAAARnQU1BAACx\r\njwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAZxSURBVHhe7chBDcAwAMSw8ifdFUAQ3BzJn5zX\r\nBQDgV3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3IC\r\nALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALAr\r\nJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAAu3ICALArJwAA\r\nu3ICALArJwAAu3ICALArJwAAu3ICADDp3A/mI7CBEzt/RwAAAABJRU5ErkJggg==",
			RulerGuidesX:        []interface{}{},
			RulerGuidesY:        []interface{}{},
			RulerShowing:        rawIsAutoSave{Type: "bool", Value: false},
			Title:               "daword",
			AudioNarrationNotes: "",
			CalloutStyle:        "Basic",
		},
	})
}

func (p *Project) encodeRawTracks() []rawTrack {
	rawTracks := []rawTrack{}

	for i, track := range p.Tracks {
		_ = track
		rawTracks = append(rawTracks, rawTrack{
			TrackIndex: i,
			Medias:     p.encodeRawTrackMedias(track),
		})
	}

	return rawTracks
}

func (p *Project) encodeRawTrackMedias(track *Track) []rawMedia {
	medias := []rawMedia{}

	var start time.Duration

	for _, element := range track.Elements {
		rMedia := rawMedia{
			ID:              p.id.gen(),
			Type:            element.node.camType(),
			Start:           p.encodeTime(start + element.gap),
			Duration:        p.encodeTime(element.duration),
			MediaDuration:   p.encodeTime(element.duration),
			Scalar:          1,
			Effects:         []interface{}{},
			Metadata:        element._rawMetadata,
			AnimationTracks: struct{}{},
		}

		switch node := element.node.(type) {
		case *Callout:
			node.encodeIntoMedia(&rMedia)
		}

		medias = append(medias, rMedia)
		start += element.gap + element.duration + 1
	}

	return medias
}

func (p *Project) encodeTrackAttributes() []rawTrackAttribute {
	attributes := []rawTrackAttribute{}

	for _, track := range p.Tracks {
		attributes = append(attributes, rawTrackAttribute{
			Ident:       track.Name,
			AudioMuted:  false,
			VideoHidden: false,
			Magnetic:    false,
			Matte:       0,
			Solo:        false,
			Metadata: rawTrackAttributeMetadata{
				IsLocked:       "False",
				WinTrackHeight: "56",
			},
		})
	}

	return attributes
}

func (p *Project) encodeTrackMedias_old(track *Track) []jobj {
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
		case *ImageFile:
			obj["_type"] = "IMFile"
			obj["src"] = p.mediaItemId[node.Src]
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
			p.mediaItemId[item] = id
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
	pCenter := p._width / 2
	eCenter := posX + (width / 2)
	return eCenter - pCenter
}

func (p *Project) coordY(height, posY int) int {
	pCenter := p._height / 2
	eCenter := posY + (height / 2)
	return pCenter - eCenter
}

func (p *Project) encodeTime(dur time.Duration) int {
	return int(dur.Milliseconds()) * (p.editRate / 1000)
}
