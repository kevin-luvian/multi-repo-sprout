package display

import "fmt"

type DisplayTerminal struct {
}

func NewDisplayTerminal() *DisplayTerminal {
	return &DisplayTerminal{}
}

func (d *DisplayTerminal) Draw(board [][]Drawable) {
	fmt.Println()

	for i, row := range board {
		if i == 0 {
			fmt.Print("  ")
			for i := range len(row) {
				fmt.Printf(" %02d", i)
			}
			fmt.Println()

			fmt.Print("  ")
			for range len(row) {
				fmt.Print(" --")
			}
			fmt.Println()
		}

		for j, cell := range row {
			if j == 0 {
				fmt.Printf("%02d|", i)
			}
			value, color := cell.Draw()
			fmt.Printf("%s %s%s", color, value, COLOR_RESET)
			fmt.Print("|")
		}
		fmt.Println()

		if i == len(board)-1 {
			fmt.Print("  ")
			for range len(row) {
				fmt.Print(" --")
			}
			fmt.Println()
		}
	}

	fmt.Println()
}
