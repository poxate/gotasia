package gotasia

import (
	"time"
)

type Element struct {
	node     Node
	gap      time.Duration
	duration time.Duration
	scaleX   float64
	scaleY   float64

	xSet bool
	ySet bool
	x    int
	y    int

	Animations []*Animation
}

type Node interface {
	node()
	camType() string
	Dimensions
}

type Dimensions interface {
	width() int
	height() int
}

func (p *Project) NewElement(node Node, animations []*Animation) *Element {
	return &Element{
		node:       node,
		duration:   5 * time.Second,
		scaleX:     1,
		scaleY:     1,
		Animations: animations,
	}
}

func (e *Element) SetScaleX(scaleX float64) *Element {
	e.scaleX = scaleX
	return e
}

func (e *Element) SetScaleY(scaleY float64) *Element {
	e.scaleY = scaleY
	return e
}

func (e *Element) SetScale(scale float64) *Element {
	return e.SetScaleX(scale).SetScaleY(scale)
}

func (e *Element) ScaleToFit(node Dimensions) *Element {
	newScale := min(
		float64(node.width())/float64(e.node.width()),
		float64(node.height())/float64(e.node.height()),
	)
	e.scaleX = newScale
	e.scaleY = newScale
	return e
}

func (e *Element) SetX(x int) *Element {
	e.x = x
	e.xSet = true
	return e
}

func (e *Element) SetY(y int) *Element {
	e.y = y
	e.ySet = true
	return e
}

func (e *Element) SetXY(x, y int) *Element {
	return e.SetX(x).SetY(y)
}

func NewAnimation(gap time.Duration, dur time.Duration) *Animation {
	a := &Animation{Gap: gap, Duration: dur}
	return a
}

// Animations
type Animation struct {
	Gap      time.Duration
	Duration time.Duration
	x        *int
	y        *int
	scaleX   *float64
	scaleY   *float64
}

func (a *Animation) ToX(x int) *Animation {
	a.x = &x
	return a
}

func (a *Animation) ToY(y int) *Animation {
	a.y = &y
	return a
}

func (a *Animation) ToScaleX(scaleX float64) *Animation {
	a.scaleX = &scaleX
	return a
}

func (a *Animation) ToScaleY(scaleY float64) *Animation {
	a.scaleY = &scaleY
	return a
}

func (a *Animation) ToScale(scale float64) *Animation {
	return a.ToScaleX(scale).ToScaleY(scale)
}
