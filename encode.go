package gotasia

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"time"
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
				FontSize:                 32,
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
		if element.duration == 0 {
			element.duration = 5 * time.Second
		}

		rMedia := rawMedia{
			ID:              p.id.gen(),
			Type:            element.Node.camType(),
			Start:           p.encodeTime(start + element.gap),
			Duration:        p.encodeTime(element.duration),
			MediaDuration:   p.encodeTime(element.duration),
			Scalar:          1,
			Effects:         []interface{}{},
			Metadata:        element._rawMetadata,
			AnimationTracks: struct{}{},
		}
		rMedia.Metadata.ClipSpeedAttribute.Type = "bool"

		switch node := element.Node.(type) {
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
		if track.Name == "" {
			track.Name = ""
		}

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
