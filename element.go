package gotasia

import "time"

type Element struct {
	Node     Node
	Gap      time.Duration
	Duration time.Duration
	X        int
	Y        int
}

type Node interface {
	width() int
	height() int
}
