package gotasia

type rawProject struct {
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
						TrackIndex int `json:"trackIndex"`
						Medias     []struct {
							ID   int    `json:"id"`
							Type string `json:"_type"`
							Def  struct {
								Kind                 string  `json:"kind"`
								Shape                string  `json:"shape"`
								Style                string  `json:"style"`
								CornerRadius         float64 `json:"corner-radius"`
								EnableLigatures      float64 `json:"enable-ligatures"`
								Height               float64 `json:"height"`
								LineSpacing          float64 `json:"line-spacing"`
								TextStrokeAlignment  float64 `json:"text-stroke-alignment"`
								TextStrokeColorAlpha float64 `json:"text-stroke-color-alpha"`
								TextStrokeColorBlue  float64 `json:"text-stroke-color-blue"`
								TextStrokeColorGreen float64 `json:"text-stroke-color-green"`
								TextStrokeColorRed   float64 `json:"text-stroke-color-red"`
								TextStrokeWidth      float64 `json:"text-stroke-width"`
								Width                struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Interp   string  `json:"interp"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"width"`
								WordWrap            float64 `json:"word-wrap"`
								HorizontalAlignment string  `json:"horizontal-alignment"`
								ResizeBehavior      string  `json:"resize-behavior"`
								Text                string  `json:"text"`
								VerticalAlignment   string  `json:"vertical-alignment"`
								Font                struct {
									ColorBlue  float64 `json:"color-blue"`
									ColorGreen float64 `json:"color-green"`
									ColorRed   float64 `json:"color-red"`
									Size       float64 `json:"size"`
									Tracking   float64 `json:"tracking"`
									Name       string  `json:"name"`
									Weight     string  `json:"weight"`
								} `json:"font"`
								TextAttributes struct {
									Type      string `json:"type"`
									Keyframes []struct {
										EndTime int `json:"endTime"`
										Time    int `json:"time"`
										Value   []struct {
											Name       string `json:"name"`
											RangeEnd   int    `json:"rangeEnd"`
											RangeStart int    `json:"rangeStart"`
											Value      string `json:"value"`
											ValueType  string `json:"valueType"`
										} `json:"value"`
										Duration int `json:"duration"`
									} `json:"keyframes"`
								} `json:"textAttributes"`
							} `json:"def,omitempty"`
							Attributes struct {
								Ident          string `json:"ident"`
								AutoRotateText bool   `json:"autoRotateText"`
							} `json:"attributes"`
							Parameters struct {
								Translation0 struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Interp   string  `json:"interp,omitempty"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"translation0"`
								Translation1 struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Interp   string  `json:"interp,omitempty"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"translation1"`
								Translation2 struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Interp   string  `json:"interp"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"translation2"`
								Rotation1 struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Interp   string  `json:"interp"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"rotation1"`
								Shear1 struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Interp   string  `json:"interp"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"shear1"`
								Scale0 struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"scale0"`
								Scale1 struct {
									Type         string  `json:"type"`
									DefaultValue float64 `json:"defaultValue"`
									Keyframes    []struct {
										EndTime  int     `json:"endTime"`
										Time     int     `json:"time"`
										Value    float64 `json:"value"`
										Duration int     `json:"duration"`
									} `json:"keyframes"`
								} `json:"scale1"`
								GeometryCrop0 float64 `json:"geometryCrop0"`
								GeometryCrop1 float64 `json:"geometryCrop1"`
								GeometryCrop2 float64 `json:"geometryCrop2"`
								GeometryCrop3 float64 `json:"geometryCrop3"`
							} `json:"parameters"`
							Effects []struct {
								EffectName string `json:"effectName"`
								Bypassed   bool   `json:"bypassed"`
								Category   string `json:"category"`
								Parameters struct {
									Radius struct {
										Type         string  `json:"type"`
										DefaultValue float64 `json:"defaultValue"`
										Interp       string  `json:"interp"`
										UIHints      struct {
											UserInterfaceType int `json:"userInterfaceType"`
											UnitType          int `json:"unitType"`
										} `json:"uiHints"`
										Keyframes []struct {
											EndTime  int     `json:"endTime"`
											Time     int     `json:"time"`
											Value    float64 `json:"value"`
											Interp   string  `json:"interp,omitempty"`
											Duration int     `json:"duration"`
										} `json:"keyframes"`
									} `json:"radius"`
									Intensity struct {
										Type         string  `json:"type"`
										DefaultValue float64 `json:"defaultValue"`
										Interp       string  `json:"interp"`
										UIHints      struct {
											UserInterfaceType int `json:"userInterfaceType"`
											UnitType          int `json:"unitType"`
										} `json:"uiHints"`
										Keyframes []struct {
											EndTime  int     `json:"endTime"`
											Time     int     `json:"time"`
											Value    float64 `json:"value"`
											Duration int     `json:"duration"`
										} `json:"keyframes"`
									} `json:"intensity"`
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
						} `json:"medias"`
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
