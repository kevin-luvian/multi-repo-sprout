package state

import (
	"testing"
)

func TestBoardState_NewBoardState(t *testing.T) {
	board := NewBoardState(3)

	if len(board.Board) != 3 {
		t.Errorf("Board should have 3 rows")
	}

	if len(board.Board[0]) != 3 {
		t.Errorf("Board should have 3 columns")
	}

	for _, row := range board.Board {
		for _, tile := range row {
			if tile.IsRevealed {
				t.Errorf("Tile should not be revealed")
			}
			if tile.IsMine {
				t.Errorf("Tile should not be a mine")
			}
			if tile.IsNumber {
				t.Errorf("Tile should not be a number")
			}
		}
	}
}

func TestBoardState_GetDrawableBoard(t *testing.T) {
	board := NewBoardState(3)
	displayBoard := board.GetDrawableBoard()

	for _, row := range displayBoard {
		for _, cell := range row {
			value, _ := cell.Draw()
			if value != "#" {
				t.Errorf("Tile should be #, got %s", value)
			}
		}
	}
}

func TestBoardState_PutMine(t *testing.T) {
	testCases := []struct {
		name string
		args func() *BoardState
		test func(*testing.T, *BoardState)
	}{
		{
			name: "Test put 1 mine",
			args: func() *BoardState {
				board := NewBoardState(1)
				board.PutMine(0, 0)
				return board
			},
			test: func(t *testing.T, board *BoardState) {
				if !board.Board[0][0].IsMine {
					t.Errorf("Tile 0,0 should be a mine")
				}

				if board.Board[0][0].Number != 0 {
					t.Errorf("Tile 0,0 should have a number of 0")
				}
			},
		},
		{
			name: "Test put number around mine",
			args: func() *BoardState {
				board := NewBoardState(3)
				board.PutMine(1, 1)
				return board
			},
			test: func(t *testing.T, board *BoardState) {
				if !board.Board[1][1].IsMine {
					t.Errorf("Tile 1,1 should be a mine")
				}

				for i := range board.Size {
					for j := range board.Size {
						if i == 1 && j == 1 {
							continue
						}

						if !board.Board[i][j].IsNumber || board.Board[i][j].Number != 1 {
							t.Errorf("Tile %d,%d should have a number of 1", i, j)
						}
					}
				}
			},
		},
		{
			name: "Test number increment around mines",
			args: func() *BoardState {
				board := NewBoardState(3)
				for i := range board.Size {
					for j := range board.Size {
						if i == 1 && j == 1 {
							continue
						}

						board.PutMine(i, j)
					}
				}
				return board
			},
			test: func(t *testing.T, board *BoardState) {
				if !board.Board[1][1].IsNumber || board.Board[1][1].Number != 8 {
					t.Errorf("Tile 1,1 should have a number of 8, got %d", board.Board[1][1].Number)
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			board := testCase.args()
			testCase.test(t, board)
		})
	}
}

func TestBoardState_RevealTiles(t *testing.T) {
	testCases := []struct {
		name string
		args func() *BoardState
		test func(*testing.T, *BoardState)
	}{
		{
			name: "Test reveal 1 tile",
			args: func() *BoardState {
				board := NewBoardState(3)
				board.PutMine(1, 1)
				return board
			},
			test: func(t *testing.T, board *BoardState) {
				board.RevealTiles(0, 1)

				if board.UnrevealedTiles != 8 {
					t.Errorf("Unrevealed tiles should be 8")
				}

				if !board.Board[0][1].IsRevealed {
					t.Errorf("Tile 0,1 should be revealed")
				}
			},
		},
		{
			name: "Test reveal all adjacent tiles",
			args: func() *BoardState {
				board := NewBoardState(5)
				board.PutMine(2, 2)
				return board
			},
			test: func(t *testing.T, board *BoardState) {
				board.RevealTiles(0, 0)

				if board.UnrevealedTiles != 1 {
					t.Errorf("Unrevealed tiles should be 1, got %d", board.UnrevealedTiles)
				}

				for i := range board.Size {
					for j := range board.Size {
						if i == 2 && j == 2 {
							continue
						}

						if !board.Board[i][j].IsRevealed {
							t.Errorf("Tile %d,%d should be revealed", i, j)
						}
					}
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			board := testCase.args()
			testCase.test(t, board)
		})
	}
}

func TestBoardState_RevealAllMines(t *testing.T) {
	testCases := []struct {
		name string
		args func() *BoardState
		test func(*testing.T, *BoardState)
	}{
		{
			name: "Test reveal all mines",
			args: func() *BoardState {
				board := NewBoardState(3)
				for i := range board.Size {
					for j := range board.Size {
						board.PutMine(i, j)
					}
				}
				board.RevealAllMines()
				return board
			},
			test: func(t *testing.T, board *BoardState) {
				if board.UnrevealedTiles != 0 {
					t.Errorf("Unrevealed tiles should be 0, got %d", board.UnrevealedTiles)
				}

				for i := range board.Size {
					for j := range board.Size {
						if !board.Board[i][j].IsMine {
							t.Errorf("Tile %d,%d should be a mine", i, j)
						}
					}
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			board := testCase.args()
			testCase.test(t, board)
		})
	}
}
