package game

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/kevin-luvian/multi-repo-sprout/go-minesweeper/internal/state"
	"github.com/kevin-luvian/multi-repo-sprout/go-minesweeper/pkg/display"
)

type Game struct {
	Display *display.DisplayTerminal
	Board   *state.BoardState
	mines   int
}

type NewGameOptions struct {
	Size  int
	Mines int
}

func NewGame(opts NewGameOptions) *Game {
	fmt.Println("NewGame", opts)
	game := &Game{
		Display: display.NewDisplayTerminal(),
		Board:   state.NewBoardState(opts.Size),
		mines:   opts.Mines,
	}
	game.PutRandomMines(game.mines)

	return game
}

func (g *Game) RunCommand(cmd Command) {
	if cmd.Command == "reveal" {
		x, _ := strconv.Atoi(cmd.Args["x"])
		y, _ := strconv.Atoi(cmd.Args["y"])

		g.Board.RevealTiles(x, y)

		if g.Board.Board[x][y].IsMine {
			g.Board.IsGameOver = true
			g.Board.RevealAllMines()
			return
		}
	}

	if g.Board.UnrevealedTiles <= g.mines {
		g.Board.IsGameWon = true
	}
}

func (g *Game) RedrawBoard() {
	g.Display.Draw(g.Board.GetDrawableBoard())
}

func (g *Game) PutRandomMines(n int) {
	for range n {
		g.Board.PutMine(rand.Intn(g.Board.Size), rand.Intn(g.Board.Size))
	}
}
