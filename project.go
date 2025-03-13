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
	editRate               int
	id                     *idTracker
	_title                 string
	_targetLoudness        float64
	_backgroundColor       color.Color
	_width                 int
	_height                int
	_autoNormalizeLoudness bool
	_frameRate             FrameRate
	Tracks                 []*Track
	MediaBin               MediaBin
	mediaItemId            map[*MediaItem]int
}

func NewProject(width, height int) *Project {
	return &Project{
		id:                     &idTracker{},
		editRate:               705600000,
		_targetLoudness:        -18.0,
		_backgroundColor:       color.Black,
		_width:                 width,
		_height:                height,
		_autoNormalizeLoudness: true,
		_frameRate:             FrameRate60,
		Tracks:                 []*Track{},
		MediaBin:               MediaBin{},
		mediaItemId:            map[*MediaItem]int{},
	}
}

func (p *Project) Title(value string) *Project {
	p._title = value
	return p
}

func (p *Project) TargetLoudness(value float64) *Project {
	p._targetLoudness = value
	return p
}

func (p *Project) BgColor(color color.Color) *Project {
	p._backgroundColor = color
	return p
}

func (p *Project) AutoNormalizeLoudness(value bool) *Project {
	p._autoNormalizeLoudness = value
	return p
}

func (p *Project) FrameRate(value FrameRate) *Project {
	p._frameRate = value
	return p
}

func (p *Project) NewTrack(name string) *Track {
	track := &Track{
		Name:     name,
		Elements: []*Element{},
	}
	p.Tracks = append(p.Tracks, track)
	return track
}

func (p *Project) width() int {
	return p._width
}

func (p *Project) height() int {
	return p._height
}
