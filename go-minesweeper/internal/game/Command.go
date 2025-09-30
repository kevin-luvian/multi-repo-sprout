package game

import (
	"errors"
	"strconv"
	"strings"
)

type Command struct {
	Command string
	Args    map[string]string
}

func ParseCommand(command string, game *Game) (Command, error) {
	parts := strings.Split(command, " ")
	action := parts[0]
	parts = parts[1:]

	if action == "start" {
		if len(parts) != 2 {
			return Command{}, errors.New("invalid command format")
		}

		return NewStartCommand(parts[0], parts[1])
	}

	if action == "reveal" {
		if len(parts) != 2 {
			return Command{}, errors.New("invalid command format")
		}

		return NewRevealCommand(parts[0], parts[1], game.Board.Size)
	}

	return Command{}, errors.New("unknown command")
}

func NewStartCommand(size, mines string) (Command, error) {
	sizeInt, err := strconv.Atoi(strings.TrimSpace(size))
	if err != nil {
		return Command{}, err
	}

	minesInt, err := strconv.Atoi(strings.TrimSpace(mines))
	if err != nil {
		return Command{}, err
	}

	if sizeInt < 0 || minesInt <= 0 || minesInt >= sizeInt*sizeInt {
		return Command{}, errors.New("invalid start command format")
	}

	return Command{
		Command: "start",
		Args:    map[string]string{"size": strconv.Itoa(sizeInt), "mines": strconv.Itoa(minesInt)},
	}, nil
}

func NewRevealCommand(x, y string, boardSize int) (Command, error) {
	xInt, err := strconv.Atoi(strings.TrimSpace(x))
	if err != nil {
		return Command{}, err
	}

	yInt, err := strconv.Atoi(strings.TrimSpace(y))
	if err != nil {
		return Command{}, err
	}

	if xInt < 0 || xInt >= boardSize || yInt < 0 || yInt >= boardSize {
		return Command{}, errors.New("out of bounds")
	}

	return Command{
		Command: "reveal",
		Args:    map[string]string{"x": strconv.Itoa(xInt), "y": strconv.Itoa(yInt)},
	}, nil
}
