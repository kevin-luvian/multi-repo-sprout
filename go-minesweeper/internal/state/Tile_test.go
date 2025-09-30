package state

import (
	"testing"

	"github.com/kevin-luvian/multi-repo-sprout/go-minesweeper/pkg/display"
)

func TestTile_Draw(t *testing.T) {
	testCases := []struct {
		name string
		args func() *Tile
		test func(t *testing.T, value string, color string)
	}{
		{
			name: "Test draw unrevealed tile",
			args: func() *Tile {
				return &Tile{}
			},
			test: func(t *testing.T, value string, color string) {
				if value != "#" {
					t.Errorf("Tile should be #, got %s", value)
				}

				if color != display.COLOR_GRAY {
					t.Errorf("Tile should have no color, got %s", color)
				}
			},
		},
		{
			name: "Test draw revealed empty tile",
			args: func() *Tile {
				return &Tile{
					IsRevealed: true,
				}
			},
			test: func(t *testing.T, value string, color string) {
				if value != " " {
					t.Errorf("Tile should be ' ', got %s", value)
				}
			},
		},
		{
			name: "Test draw revealed number tile",
			args: func() *Tile {
				return &Tile{
					IsMine:     false,
					IsRevealed: true,
					IsNumber:   true,
					Number:     1,
				}
			},
			test: func(t *testing.T, value string, color string) {
				if value != "1" {
					t.Errorf("Tile should be 1, got %s", value)
				}
			},
		},
		{
			name: "Test draw revealed mine tile",
			args: func() *Tile {
				return &Tile{
					IsMine:     true,
					IsRevealed: true,
				}
			},
			test: func(t *testing.T, value string, color string) {
				if value != "*" {
					t.Errorf("Tile should be *, got %s", value)
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tile := testCase.args()
			value, color := tile.Draw()
			testCase.test(t, value, color)
		})
	}

}
