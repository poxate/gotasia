package gotasia

type rawProject struct {
	Title                            string                 `json:"title"`
	Description                      string                 `json:"description"`
	Author                           string                 `json:"author"`
	TargetLoudness                   KeepZero               `json:"targetLoudness"`
	ShouldApplyLoudnessNormalization bool                   `json:"shouldApplyLoudnessNormalization"`
	VideoFormatFrameRate             int                    `json:"videoFormatFrameRate"`
	AudioFormatSampleRate            int                    `json:"audioFormatSampleRate"`
	AllowSubFrameEditing             bool                   `json:"allowSubFrameEditing"`
	Width                            KeepZero               `json:"width"`
	Height                           KeepZero               `json:"height"`
	Version                          string                 `json:"version"`
	EditRate                         int                    `json:"editRate"`
	AuthoringClientName              rawAuthoringClientName `json:"authoringClientName"`
	Timeline                         rawTimeline            `json:"timeline"`
	Metadata                         rawMetadata            `json:"metadata"`
}

type rawAuthoringClientName struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

type rawMetadata struct {
	AutoSaveFile        string        `json:"AutoSaveFile"`
	CanvasZoom          rawCanvasZoom `json:"CanvasZoom"`
	Date                string        `json:"Date"`
	IsAutoSave          rawIsAutoSave `json:"IsAutoSave"`
	IsStandalone        rawIsAutoSave `json:"IsStandalone"`
	Language            string        `json:"Language"`
	ProfileName         string        `json:"ProfileName"`
	ProjectThumbnail    string        `json:"ProjectThumbnail"`
	RulerGuidesX        []interface{} `json:"RulerGuidesX"`
	RulerGuidesY        []interface{} `json:"RulerGuidesY"`
	RulerShowing        rawIsAutoSave `json:"RulerShowing"`
	Title               string        `json:"Title"`
	AudioNarrationNotes string        `json:"audioNarrationNotes"`
	CalloutStyle        string        `json:"calloutStyle"`
}

type rawCanvasZoom struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type rawIsAutoSave struct {
	Type  string `json:"type"`
	Value bool   `json:"value"`
}

type rawTimeline struct {
	ID                      int                  `json:"id"`
	SceneTrack              rawSceneTrack        `json:"sceneTrack"`
	TrackAttributes         []rawTrackAttribute  `json:"trackAttributes"`
	CaptionAttributes       rawCaptionAttributes `json:"captionAttributes"`
	Gain                    KeepZero             `json:"gain"`
	LegacyAttenuateAudioMix bool                 `json:"legacyAttenuateAudioMix"`
	BackgroundColor         interface{}          `json:"backgroundColor"`
}

type rawCaptionAttributes struct {
	Enabled                  bool        `json:"enabled"`
	FontName                 string      `json:"fontName"`
	FontSize                 int         `json:"fontSize"`
	BackgroundColor          interface{} `json:"backgroundColor"`
	ForegroundColor          interface{} `json:"foregroundColor"`
	Lang                     string      `json:"lang"`
	Alignment                int         `json:"alignment"`
	DefaultFontSize          bool        `json:"defaultFontSize"`
	Opacity                  float64     `json:"opacity"`
	BackgroundEnabled        bool        `json:"backgroundEnabled"`
	BackgroundOnlyAroundText bool        `json:"backgroundOnlyAroundText"`
}

type rawSceneTrack struct {
	Scenes []rawScene `json:"scenes"`
}

type rawScene struct {
	Csml rawCsml `json:"csml"`
}

type rawCsml struct {
	Tracks []rawTrack `json:"tracks"`
}

type rawTrack struct {
	TrackIndex int        `json:"trackIndex"`
	Medias     []rawMedia `json:"medias"`
}

type rawMedia struct {
	ID   int     `json:"id"`
	Type string  `json:"_type"`
	Def  *rawDef `json:"def"`
	// Attributes      Attributes      `json:"attributes"`
	// Parameters      Parameters      `json:"parameters"`
	Effects       []interface{} `json:"effects"`
	Start         int           `json:"start"`
	Duration      int           `json:"duration"`
	MediaStart    int           `json:"mediaStart"`
	MediaDuration int           `json:"mediaDuration"`
	Scalar        int           `json:"scalar"`
	// Metadata        MediaMetadata   `json:"metadata"`
	// AnimationTracks AnimationTracks `json:"animationTracks"`
}

type rawTrackAttribute struct {
	Ident       string                    `json:"ident"`
	AudioMuted  bool                      `json:"audioMuted"`
	VideoHidden bool                      `json:"videoHidden"`
	Magnetic    bool                      `json:"magnetic"`
	Matte       int                       `json:"matte"`
	Solo        bool                      `json:"solo"`
	Metadata    rawTrackAttributeMetadata `json:"metadata"`
}

type rawTrackAttributeMetadata struct {
	IsLocked       string `json:"IsLocked"`
	WinTrackHeight string `json:"WinTrackHeight"`
}

type rawDef struct {
	Kind                 string   `json:"kind"`
	Shape                string   `json:"shape"`
	Style                string   `json:"style"`
	CornerRadius         KeepZero `json:"corner-radius"`
	EnableLigatures      KeepZero `json:"enable-ligatures"`
	Height               KeepZero `json:"height"`
	LineSpacing          KeepZero `json:"line-spacing"`
	TextStrokeAlignment  KeepZero `json:"text-stroke-alignment"`
	TextStrokeColorAlpha KeepZero `json:"text-stroke-color-alpha"`
	TextStrokeColorBlue  KeepZero `json:"text-stroke-color-blue"`
	TextStrokeColorGreen KeepZero `json:"text-stroke-color-green"`
	TextStrokeColorRed   KeepZero `json:"text-stroke-color-red"`
	TextStrokeWidth      KeepZero `json:"text-stroke-width"`
	Width                KeepZero `json:"width"`
	WordWrap             KeepZero `json:"word-wrap"`
	HorizontalAlignment  string   `json:"horizontal-alignment"`
	ResizeBehavior       string   `json:"resize-behavior"`
	Text                 string   `json:"text"`
	VerticalAlignment    string   `json:"vertical-alignment"`
	Font                 rawFont  `json:"font"`
	// TextAttributes       TextAttributes `json:"textAttributes"`
}

type rawFont struct {
	ColorBlue  KeepZero `json:"color-blue"`
	ColorGreen KeepZero `json:"color-green"`
	ColorRed   KeepZero `json:"color-red"`
	Size       KeepZero `json:"size"`
	Tracking   KeepZero `json:"tracking"`
	Name       string   `json:"name"`
	Weight     string   `json:"weight"`
}
