package gotasia

import (
	"time"
)

type Element struct {
	node     Node
	gap      time.Duration
	duration time.Duration
	scale    float64

	xSet bool
	ySet bool
	x    int
	y    int
}

type Node interface {
	Dimensions
}

type Dimensions interface {
	width() int
	height() int
}

type Animation struct {
	Gap time.Duration
	X   int
	Y   int
}

func (p *Project) NewElement(node Node) *Element {
	return &Element{
		node:     node,
		duration: 5 * time.Second,
		scale:    1,
	}
}

func (e *Element) SetScale(scale float64) *Element {
	e.scale = scale
	return e
}

func (e *Element) ScaleToFit(node Dimensions) *Element {
	e.scale = min(
		float64(node.width())/float64(e.node.width()),
		float64(node.height())/float64(e.node.height()),
	)
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
