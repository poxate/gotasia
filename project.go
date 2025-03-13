package gotasia

import (
	"image/color"
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
