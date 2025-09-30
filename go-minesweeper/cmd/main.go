package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kevin-luvian/multi-repo-sprout/go-minesweeper/internal/game"
)

func main() {
	var minesweeper *game.Game = nil

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nEnter a command: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		command, err := game.ParseCommand(input, minesweeper)
		if err != nil {
			fmt.Println("Invalid command: ", err)
			fmt.Println("Please try again!")
			continue
		}

		if command.Command == "start" && minesweeper == nil {
			size, _ := strconv.Atoi(command.Args["size"])
			mines, _ := strconv.Atoi(command.Args["mines"])
			minesweeper = game.NewGame(game.NewGameOptions{
				Size:  size,
				Mines: mines,
			})
			minesweeper.RedrawBoard()
			continue
		} else if command.Command == "start" {
			fmt.Println("Game already started!")
			continue
		}

		if minesweeper == nil {
			fmt.Println("Please start a game first!")
			continue
		}

		minesweeper.RunCommand(command)
		minesweeper.RedrawBoard()

		if minesweeper.Board.IsGameOver {
			fmt.Println("Game over!")
			break
		}

		if minesweeper.Board.IsGameWon {
			fmt.Println("Game won!")
			break
		}
	}
}
