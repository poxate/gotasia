package gotasia

type Project struct {
	Width  int
	Height int
}

func NewProject(width, height int) *Project {
	return &Project{
		Width:  width,
		Height: height,
	}
}
