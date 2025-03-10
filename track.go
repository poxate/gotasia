package gotasia

type Track struct {
	Name     string
	Elements []*Element
}

func NewTrack(name string, elements ...*Element) *Track {
	return &Track{Name: name, Elements: elements}
}
