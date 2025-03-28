package gotasia

type rawProject struct {
	Title                            string                 `json:"title"`
	Description                      string                 `json:"description"`
	Author                           string                 `json:"author"`
	TargetLoudness                   keepZero               `json:"targetLoudness"`
	ShouldApplyLoudnessNormalization bool                   `json:"shouldApplyLoudnessNormalization"`
	VideoFormatFrameRate             int                    `json:"videoFormatFrameRate"`
	AudioFormatSampleRate            int                    `json:"audioFormatSampleRate"`
	AllowSubFrameEditing             bool                   `json:"allowSubFrameEditing"`
	Width                            keepZero               `json:"width"`
	Height                           keepZero               `json:"height"`
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
	Gain                    keepZero             `json:"gain"`
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
	ID              int              `json:"id"`
	Type            string           `json:"_type"`
	Def             *rawDef          `json:"def"`
	Attributes      rawAttributes    `json:"attributes"`
	Parameters      rawParameters    `json:"parameters"`
	Effects         []interface{}    `json:"effects"`
	Start           int              `json:"start"`
	Duration        int              `json:"duration"`
	MediaStart      int              `json:"mediaStart"`
	MediaDuration   int              `json:"mediaDuration"`
	Scalar          int              `json:"scalar"`
	Metadata        rawMediaMetadata `json:"metadata"`
	AnimationTracks any              `json:"animationTracks"`
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
	Kind                 string            `json:"kind"`
	Shape                string            `json:"shape"`
	Style                string            `json:"style"`
	CornerRadius         keepZero          `json:"corner-radius"`
	EnableLigatures      keepZero          `json:"enable-ligatures"`
	FillColorBlue        *keepZero         `json:"fill-color-blue,omitempty"`
	FillColorGreen       *keepZero         `json:"fill-color-green,omitempty"`
	FillColorOpacity     *keepZero         `json:"fill-color-opacity,omitempty"`
	FillColorRed         *keepZero         `json:"fill-color-red,omitempty"`
	Height               keepZero          `json:"height"`
	LineSpacing          keepZero          `json:"line-spacing"`
	StrokeColorBlue      *keepZero         `json:"stroke-color-blue,omitempty"`
	StrokeColorGreen     *keepZero         `json:"stroke-color-green,omitempty"`
	StrokeColorOpacity   *keepZero         `json:"stroke-color-opacity,omitempty"`
	StrokeColorRed       *keepZero         `json:"stroke-color-red,omitempty"`
	StrokeWidth          *keepZero         `json:"stroke-width,omitempty"`
	TailX                *keepZero         `json:"tail-x,omitempty"`
	TailY                *keepZero         `json:"tail-y,omitempty"`
	TextStrokeAlignment  keepZero          `json:"text-stroke-alignment"`
	TextStrokeColorAlpha keepZero          `json:"text-stroke-color-alpha"`
	TextStrokeColorBlue  keepZero          `json:"text-stroke-color-blue"`
	TextStrokeColorGreen keepZero          `json:"text-stroke-color-green"`
	TextStrokeColorRed   keepZero          `json:"text-stroke-color-red"`
	TextStrokeWidth      keepZero          `json:"text-stroke-width"`
	Width                keepZero          `json:"width"`
	WordWrap             keepZero          `json:"word-wrap"`
	FillStyle            *string           `json:"fill-style,omitempty"`
	HorizontalAlignment  string            `json:"horizontal-alignment"`
	ResizeBehavior       string            `json:"resize-behavior"`
	StrokeStyle          *string           `json:"stroke-style,omitempty"`
	Text                 string            `json:"text"`
	VerticalAlignment    string            `json:"vertical-alignment"`
	Font                 rawFont           `json:"font"`
	TextAttributes       rawTextAttributes `json:"textAttributes"`
}

type rawFont struct {
	ColorBlue  keepZero `json:"color-blue"`
	ColorGreen keepZero `json:"color-green"`
	ColorRed   keepZero `json:"color-red"`
	Size       keepZero `json:"size"`
	Tracking   keepZero `json:"tracking"`
	Name       string   `json:"name"`
	Weight     string   `json:"weight"`
}

type rawTextAttributes struct {
	Type      string                      `json:"type"`
	Keyframes []Keyframe[[]TextAttribute] `json:"keyframes"`
}

type Keyframe[T any] struct {
	EndTime  int `json:"endTime"`
	Time     int `json:"time"`
	Value    T   `json:"value"`
	Duration int `json:"duration"`
}

type TextAttribute struct {
	Name       string `json:"name"`
	RangeEnd   int    `json:"rangeEnd"`
	RangeStart int    `json:"rangeStart"`
	Value      any    `json:"value"`
	ValueType  string `json:"valueType"`
}

type rawAttributes struct {
	Ident          string `json:"ident"`
	AutoRotateText *bool  `json:"autoRotateText,omitempty"`
}

type rawParameters struct {
	GeometryCrop0 keepZero `json:"geometryCrop0"`
	GeometryCrop1 keepZero `json:"geometryCrop1"`
	GeometryCrop2 keepZero `json:"geometryCrop2"`
	GeometryCrop3 keepZero `json:"geometryCrop3"`
}

type rawMediaMetadata struct {
	AudiateLinkedSession        *string      `json:"audiateLinkedSession,omitempty"`
	ClipSpeedAttribute          *IsAutoSave  `json:"clipSpeedAttribute,omitempty"`
	DefaultHAlign               *string      `json:"default-HAlign,omitempty"`
	DefaultLineSpace            *CanvasZoom  `json:"default-LineSpace,omitempty"`
	DefaultVAlign               *string      `json:"default-VAlign,omitempty"`
	DefaultTextAttributes       *interface{} `json:"default-text-attributes,omitempty"`
	DefaultTextStrokeColorBlue  *CanvasZoom  `json:"default-text-stroke-color-blue,omitempty"`
	DefaultTextStrokeColorGreen *CanvasZoom  `json:"default-text-stroke-color-green,omitempty"`
	DefaultTextStrokeColorRed   *CanvasZoom  `json:"default-text-stroke-color-red,omitempty"`
	DefaultTextStrokeOpacity    *CanvasZoom  `json:"default-text-stroke-opacity,omitempty"`
	DefaultTextStrokeWidth      *CanvasZoom  `json:"default-text-stroke-width,omitempty"`
}

type CanvasZoom struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

type IsAutoSave struct {
	Type  string `json:"type"`
	Value bool   `json:"value"`
}
