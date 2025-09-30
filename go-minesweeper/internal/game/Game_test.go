package game

import (
	"testing"
)

func TestGame_RunCommand(t *testing.T) {
	testCases := []struct {
		name string
		args func() (*Game, Command)
		test func(*testing.T, *Game)
	}{
		{
			name: "Test run reveal command",
			args: func() (*Game, Command) {
				game := NewGame(NewGameOptions{
					Size:  10,
					Mines: 0,
				})
				game.Board.PutMine(0, 0)
				command := Command{
					Command: "reveal",
					Args:    map[string]string{"x": "5", "y": "5"},
				}
				return game, command
			},
			test: func(t *testing.T, game *Game) {
				if game.Board.IsGameOver {
					t.Errorf("Game should be not over")
				}

				if game.Board.IsGameWon {
					t.Errorf("Game should be not won")
				}

				if game.Board.UnrevealedTiles != 1 {
					t.Errorf("Unrevealed tiles should be 1, got %d", game.Board.UnrevealedTiles)
				}
			},
		},
		{
			name: "Test run reveal command with mine",
			args: func() (*Game, Command) {
				game := NewGame(NewGameOptions{
					Size:  10,
					Mines: 0,
				})
				game.Board.PutMine(5, 5)
				command := Command{
					Command: "reveal",
					Args:    map[string]string{"x": "5", "y": "5"},
				}
				return game, command
			},
			test: func(t *testing.T, game *Game) {
				if !game.Board.IsGameOver {
					t.Errorf("Game should be over")
				}

				if game.Board.IsGameWon {
					t.Errorf("Game should be not won")
				}
			},
		},
		{
			name: "Test run reveal command with all mines avoided",
			args: func() (*Game, Command) {
				game := NewGame(NewGameOptions{
					Size:  10,
					Mines: 0,
				})
				game.mines = 1
				game.Board.PutMine(5, 5)
				command := Command{
					Command: "reveal",
					Args:    map[string]string{"x": "0", "y": "0"},
				}
				return game, command
			},
			test: func(t *testing.T, game *Game) {
				if !game.Board.IsGameWon {
					t.Errorf("Game should be won")
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			game, command := testCase.args()
			game.RunCommand(command)
			testCase.test(t, game)
		})
	}
}
