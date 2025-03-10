package gotasia

import "time"

type Element struct {
	Node          Node
	Start         time.Duration
	Duration      time.Duration
	MediaStart    *time.Duration
	MediaDuration *time.Duration
	X             int
	Y             int
}

type Node interface {
	width() int
	height() int
}
