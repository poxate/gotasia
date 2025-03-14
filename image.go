package gotasia

type ImageFile struct {
	Src   *MediaItem
	Ident string
}

func (i *ImageFile) width() int {
	return i.Src.Width
}
func (i *ImageFile) height() int {
	return i.Src.Height
}

func (i *ImageFile) node()           {}
func (i *ImageFile) camType() string { return "IMFile" }
