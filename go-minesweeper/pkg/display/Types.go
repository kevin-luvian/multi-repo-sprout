package display

type Drawable interface {
	Draw() (value string, color string)
}
