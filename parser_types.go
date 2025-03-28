package gotasia

import (
	"encoding/json"
	"fmt"
)

type rawProjectD struct {
	Title                            string  `json:"title"`
	Description                      string  `json:"description"`
	Author                           string  `json:"author"`
	TargetLoudness                   float64 `json:"targetLoudness"`
	ShouldApplyLoudnessNormalization bool    `json:"shouldApplyLoudnessNormalization"`
	VideoFormatFrameRate             int     `json:"videoFormatFrameRate"`
	AudioFormatSampleRate            int     `json:"audioFormatSampleRate"`
	AllowSubFrameEditing             bool    `json:"allowSubFrameEditing"`
	Width                            float64 `json:"width"`
	Height                           float64 `json:"height"`
	Version                          string  `json:"version"`
	EditRate                         int     `json:"editRate"`
	AuthoringClientName              struct {
		Name     string `json:"name"`
		Platform string `json:"platform"`
		Version  string `json:"version"`
	} `json:"authoringClientName"`
	SourceBin []struct {
		ID                    int    `json:"id"`
		Src                   string `json:"src"`
		Rect                  []int  `json:"rect"`
		LastMod               string `json:"lastMod"`
		LoudnessNormalization bool   `json:"loudnessNormalization"`
		SourceTracks          []struct {
			Range          []int   `json:"range"`
			Type           int     `json:"type"`
			EditRate       int     `json:"editRate"`
			TrackRect      []int   `json:"trackRect"`
			SampleRate     int     `json:"sampleRate"`
			BitDepth       int     `json:"bitDepth"`
			NumChannels    int     `json:"numChannels"`
			IntegratedLUFS float64 `json:"integratedLUFS"`
			PeakLevel      float64 `json:"peakLevel"`
			MetaData       string  `json:"metaData"`
		} `json:"sourceTracks"`
		Metadata struct {
			TimeAdded string `json:"timeAdded"`
		} `json:"metadata"`
	} `json:"sourceBin"`
	Timeline struct {
		ID         int `json:"id"`
		SceneTrack struct {
			Scenes []struct {
				Csml struct {
					Tracks []struct {
						TrackIndex  int            `json:"trackIndex"`
						Medias      []rawMedia_old `json:"medias"`
						Transitions []struct {
							Name       string `json:"name"`
							Duration   int    `json:"duration"`
							LeftMedia  int    `json:"leftMedia"`
							Attributes struct {
								Random           float64 `json:"Random"`
								Bypass           bool    `json:"bypass"`
								Reverse          bool    `json:"reverse"`
								Trivial          bool    `json:"trivial"`
								UseAudioPreRoll  bool    `json:"useAudioPreRoll"`
								UseVisualPreRoll bool    `json:"useVisualPreRoll"`
							} `json:"attributes"`
						} `json:"transitions,omitempty"`
					} `json:"tracks"`
				} `json:"csml"`
			} `json:"scenes"`
		} `json:"sceneTrack"`
		TrackAttributes []struct {
			Ident       string `json:"ident"`
			AudioMuted  bool   `json:"audioMuted"`
			VideoHidden bool   `json:"videoHidden"`
			Magnetic    bool   `json:"magnetic"`
			Matte       int    `json:"matte"`
			Solo        bool   `json:"solo"`
			Metadata    struct {
				IsLocked       string `json:"IsLocked"`
				WinTrackHeight string `json:"WinTrackHeight"`
			} `json:"metadata"`
		} `json:"trackAttributes"`
		CaptionAttributes struct {
			Enabled                  bool    `json:"enabled"`
			FontName                 string  `json:"fontName"`
			FontSize                 int     `json:"fontSize"`
			BackgroundColor          []int   `json:"backgroundColor"`
			ForegroundColor          []int   `json:"foregroundColor"`
			Lang                     string  `json:"lang"`
			Alignment                int     `json:"alignment"`
			DefaultFontSize          bool    `json:"defaultFontSize"`
			Opacity                  float64 `json:"opacity"`
			BackgroundEnabled        bool    `json:"backgroundEnabled"`
			BackgroundOnlyAroundText bool    `json:"backgroundOnlyAroundText"`
		} `json:"captionAttributes"`
		Gain                    float64 `json:"gain"`
		LegacyAttenuateAudioMix bool    `json:"legacyAttenuateAudioMix"`
		BackgroundColor         []int   `json:"backgroundColor"`
	} `json:"timeline"`
	Metadata struct {
		AutoSaveFile string `json:"AutoSaveFile"`
		CanvasZoom   struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"CanvasZoom"`
		Date string `json:"Date"`
		Fit  struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"Fit"`
		IsAutoSave struct {
			Type  string `json:"type"`
			Value bool   `json:"value"`
		} `json:"IsAutoSave"`
		IsStandalone struct {
			Type  string `json:"type"`
			Value bool   `json:"value"`
		} `json:"IsStandalone"`
		ProfileName      string        `json:"ProfileName"`
		ProjectThumbnail string        `json:"ProjectThumbnail"`
		RulerGuidesX     []interface{} `json:"RulerGuidesX"`
		RulerGuidesY     []interface{} `json:"RulerGuidesY"`
		RulerShowing     struct {
			Type  string `json:"type"`
			Value bool   `json:"value"`
		} `json:"RulerShowing"`
		Title               string `json:"Title"`
		AudioNarrationNotes string `json:"audioNarrationNotes"`
		CalloutStyle        string `json:"calloutStyle"`
	} `json:"metadata"`
}

type rawMedia_old struct {
	ID   int    `json:"id"`
	Type string `json:"_type"`
	Def  struct {
		Kind                 string                     `json:"kind"`
		Shape                string                     `json:"shape"`
		Style                string                     `json:"style"`
		CornerRadius         any                        `json:"corner-radius"`
		EnableLigatures      any                        `json:"enable-ligatures"`
		Width                rawMaybeKeyframes[float64] `json:"width"`
		Height               rawMaybeKeyframes[float64] `json:"height"`
		LineSpacing          any                        `json:"line-spacing"`
		TextStrokeAlignment  any                        `json:"text-stroke-alignment"`
		TextStrokeColorAlpha any                        `json:"text-stroke-color-alpha"`
		TextStrokeColorBlue  any                        `json:"text-stroke-color-blue"`
		TextStrokeColorGreen any                        `json:"text-stroke-color-green"`
		TextStrokeColorRed   any                        `json:"text-stroke-color-red"`
		TextStrokeWidth      any                        `json:"text-stroke-width"`
		WordWrap             any                        `json:"word-wrap"`
		HorizontalAlignment  string                     `json:"horizontal-alignment"`
		ResizeBehavior       string                     `json:"resize-behavior"`
		Text                 string                     `json:"text"`
		VerticalAlignment    string                     `json:"vertical-alignment"`
		Font                 struct {
			ColorBlue  float64 `json:"color-blue"`
			ColorGreen float64 `json:"color-green"`
			ColorRed   float64 `json:"color-red"`
			Size       float64 `json:"size"`
			Tracking   float64 `json:"tracking"`
			Name       string  `json:"name"`
			Weight     string  `json:"weight"`
		} `json:"font"`
		TextAttributes rawTextAttributes `json:"textAttributes"`
	} `json:"def,omitempty"`
	Attributes struct {
		Ident          string `json:"ident"`
		AutoRotateText bool   `json:"autoRotateText"`
	} `json:"attributes"`
	Parameters struct {
		Translation0  rawMaybeKeyframes[float64] `json:"translation0"`
		Translation1  rawMaybeKeyframes[float64] `json:"translation1"`
		Translation2  rawMaybeKeyframes[float64] `json:"translation2"`
		Rotation1     rawMaybeKeyframes[float64] `json:"rotation1"`
		Shear1        rawMaybeKeyframes[float64] `json:"shear1"`
		Scale0        rawMaybeKeyframes[float64] `json:"scale0"`
		Scale1        rawMaybeKeyframes[float64] `json:"scale1"`
		GeometryCrop0 rawMaybeKeyframes[float64] `json:"geometryCrop0"`
		GeometryCrop1 rawMaybeKeyframes[float64] `json:"geometryCrop1"`
		GeometryCrop2 rawMaybeKeyframes[float64] `json:"geometryCrop2"`
		GeometryCrop3 rawMaybeKeyframes[float64] `json:"geometryCrop3"`
	} `json:"parameters"`
	Effects []struct {
		EffectName string `json:"effectName"`
		Bypassed   bool   `json:"bypassed"`
		Category   string `json:"category"`
		Parameters struct {
			Radius    rawMaybeKeyframes[float64] `json:"radius"`
			Intensity rawMaybeKeyframes[float64] `json:"intensity"`
		} `json:"parameters"`
		Metadata struct {
			DefaultGlowIntensity struct {
				Type  string  `json:"type"`
				Value float64 `json:"value"`
			} `json:"default-Glow_intensity"`
			DefaultGlowRadius struct {
				Type  string  `json:"type"`
				Value float64 `json:"value"`
			} `json:"default-Glow_radius"`
			PresetName string `json:"presetName"`
		} `json:"metadata"`
	} `json:"effects"`
	Start         int `json:"start"`
	Duration      int `json:"duration"`
	MediaStart    int `json:"mediaStart"`
	MediaDuration int `json:"mediaDuration"`
	Scalar        int `json:"scalar"`
	Metadata      struct {
		AudiateLinkedSession string `json:"audiateLinkedSession"`
		ClipSpeedAttribute   struct {
			Type  string `json:"type"`
			Value bool   `json:"value"`
		} `json:"clipSpeedAttribute"`
		DefaultHAlign    string `json:"default-HAlign"`
		DefaultLineSpace struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-LineSpace"`
		DefaultVAlign  string `json:"default-VAlign"`
		DefaultAnchor0 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-anchor0"`
		DefaultAnchor1 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-anchor1"`
		DefaultAnchor2 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-anchor2"`
		DefaultHeight struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-height"`
		DefaultRotation0 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-rotation0"`
		DefaultRotation1 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-rotation1"`
		DefaultRotation2 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-rotation2"`
		DefaultScale0 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-scale0"`
		DefaultScale1 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-scale1"`
		DefaultScale2 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-scale2"`
		DefaultShear0 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-shear0"`
		DefaultShear1 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-shear1"`
		DefaultShear2 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-shear2"`
		DefaultTextAttributes      interface{} `json:"default-text-attributes"`
		DefaultTextStrokeColorBlue struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-text-stroke-color-blue"`
		DefaultTextStrokeColorGreen struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-text-stroke-color-green"`
		DefaultTextStrokeColorRed struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-text-stroke-color-red"`
		DefaultTextStrokeOpacity struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-text-stroke-opacity"`
		DefaultTextStrokeWidth struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-text-stroke-width"`
		DefaultTranslation0 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-translation0"`
		DefaultTranslation1 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-translation1"`
		DefaultTranslation2 struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-translation2"`
		DefaultWidth struct {
			Type  string  `json:"type"`
			Value float64 `json:"value"`
		} `json:"default-width"`
	} `json:"metadata"`
	AnimationTracks struct {
		Visual []struct {
			EndTime  int    `json:"endTime"`
			Duration int    `json:"duration"`
			Range    []int  `json:"range"`
			Interp   string `json:"interp,omitempty"`
		} `json:"visual"`
	} `json:"animationTracks"`
	Src          int `json:"src,omitempty"`
	TrackNumber  int `json:"trackNumber,omitempty"`
	TrimStartSum int `json:"trimStartSum,omitempty"`
}

type rawMaybeKeyframes[T any] struct {
	// may eitheer be a static value of type T, or a keyframe of type rawKeyframe[T]
	Static      bool
	StaticValue T
	Keyframe    rawKeyframes[T]
}

type rawKeyframes[T any] struct {
	Type         string        `json:"type"`
	DefaultValue T             `json:"defaultValue"`
	Keyframes    []keyframe[T] `json:"keyframes"`
}
type keyframe[T any] struct {
	EndTime  int    `json:"endTime"`
	Time     int    `json:"time"`
	Value    T      `json:"value"`
	Interp   string `json:"interp,omitempty"`
	Duration int    `json:"duration"`
}

// create a custom unmarshal for rawMaybeKeyframe that checks if the value is T, and if so, marshal it as a static value, else, marhsal in to rawkeyframe
func (r *rawMaybeKeyframes[T]) getFirstValue() T {
	if r.Static {
		return r.StaticValue
	} else {
		return r.Keyframe.DefaultValue
	}
}

// implement marshalJSON for rawMaybeKeyframes
func (r rawMaybeKeyframes[T]) MarshalJSON() ([]byte, error) {
	if r.Static {
		return json.Marshal(r.StaticValue)
	} else {
		return json.Marshal(r.Keyframe)
	}
}

func (r *rawMaybeKeyframes[T]) UnmarshalJSON(data []byte) error {
	var staticValue T
	if err := json.Unmarshal(data, &staticValue); err == nil {
		r.Static = true
		r.StaticValue = staticValue
		return nil
	}

	var keyframe rawKeyframes[T]
	if err := json.Unmarshal(data, &keyframe); err != nil {
		return fmt.Errorf("maybeKeyframe cannot be unmarshaled as static value or keyframe: %w", err)
	}
	r.Static = false
	r.Keyframe = keyframe
	return nil
}
