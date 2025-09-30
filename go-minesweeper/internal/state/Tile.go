package state

import (
	"strconv"

	"github.com/kevin-luvian/multi-repo-sprout/go-minesweeper/pkg/display"
)

type Tile struct {
	IsMine     bool
	IsRevealed bool
	IsNumber   bool
	Number     int
}

func (t *Tile) Increment() {
	t.Number += 1
	t.IsNumber = true
}

func (t *Tile) Draw() (value string, color string) {
	if !t.IsRevealed {
		return "#", display.COLOR_GRAY
	}

	if t.IsMine {
		return "*", display.COLOR_RED
	}

	if t.IsNumber {
		color := display.COLOR_RED

		if t.Number == 1 {
			color = display.COLOR_GREEN
		}

		if t.Number == 2 {
			color = display.COLOR_YELLOW
		}

		return strconv.Itoa(t.Number), color
	}

	return " ", display.COLOR_WHITE
}

func (t *Tile) Reveal() (ok bool) {
	if t.IsRevealed {
		return false
	}

	t.IsRevealed = true
	return !t.IsMine
}

var _ display.Drawable = (*Tile)(nil)
