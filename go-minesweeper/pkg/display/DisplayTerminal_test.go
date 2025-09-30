package display

import (
	"testing"
)

type Tile struct {
	Value string
}

func (t *Tile) Draw() (value string, color string) {
	return t.Value, ""
}

func TestDisplayTerminal_Draw(t *testing.T) {
	terminal := DisplayTerminal{}
	terminal.Draw([][]Drawable{
		{&Tile{"1"}, &Tile{"2"}},
		{&Tile{"3"}, &Tile{"*"}},
		{&Tile{"5"}, &Tile{"#"}},
		{&Tile{"7"}, &Tile{"8"}},
	})

}
