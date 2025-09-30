package state

import (
	"github.com/kevin-luvian/multi-repo-sprout/go-minesweeper/pkg/display"
)

type BoardState struct {
	Size            int
	Board           [][]Tile
	IsGameOver      bool
	IsGameWon       bool
	UnrevealedTiles int
}

func NewBoardState(n int) *BoardState {
	board := make([][]Tile, n)
	for i := range board {
		board[i] = make([]Tile, n)
	}

	return &BoardState{
		Size:            n,
		IsGameOver:      false,
		IsGameWon:       false,
		Board:           board,
		UnrevealedTiles: n * n,
	}
}

func (b *BoardState) GetDrawableBoard() [][]display.Drawable {
	result := make([][]display.Drawable, len(b.Board))

	for i, row := range b.Board {
		result[i] = make([]display.Drawable, len(row))

		for j, tile := range row {
			result[i][j] = &tile
		}
	}

	return result
}

func (b *BoardState) RevealTiles(x, y int) {
	isRevealed := b.Board[x][y].Reveal()
	if isRevealed {
		b.UnrevealedTiles -= 1
	}

	if !isRevealed || b.Board[x][y].IsNumber {
		return
	}

	// reveal all adjacent tiles
	stack := make([][]int, 0)
	stack = append(stack, []int{x, y})

	for len(stack) > 0 {
		pos := stack[0]
		stack = stack[1:]

		directions := [8][2]int{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
			{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
		}
		for _, dir := range directions {
			x = pos[0] + dir[0]
			y = pos[1] + dir[1]

			// out of bounds
			if x < 0 || x >= len(b.Board) || y < 0 || y >= len(b.Board[0]) {
				continue
			}

			isRevealed := b.Board[x][y].Reveal()
			if isRevealed {
				b.UnrevealedTiles -= 1
			}

			if isRevealed && !b.Board[x][y].IsNumber {
				stack = append(stack, []int{x, y})
			}
		}
	}
}

func (b *BoardState) RevealAllMines() {
	for i, row := range b.Board {
		for j, tile := range row {
			if tile.IsMine {
				b.Board[i][j].Reveal()
				b.UnrevealedTiles -= 1
			}
		}
	}
}

func (b *BoardState) PutMine(x, y int) {
	b.Board[x][y].IsMine = true
	originX := x
	originY := y

	// increment the number of mines around the tile
	for _, dir := range [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
	} {
		x = originX + dir[0]
		y = originY + dir[1]

		if x < 0 || x >= len(b.Board) || y < 0 || y >= len(b.Board[0]) {
			continue
		}

		b.Board[x][y].Increment()
	}
}
