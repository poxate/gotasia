package gotasia

type MediaType int

const (
	ImageMediaItem MediaType = iota
)

type MediaBin []*MediaItem

type MediaItem struct {
	Src    string
	Width  int
	Height int
	Type   MediaType
}
