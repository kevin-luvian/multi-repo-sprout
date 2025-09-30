# Go Minesweeper

A simple terminal-based Minesweeper game written in Go.

## Features

- Play classic Minesweeper in your terminal
- Configurable board size and mine count
- Keyboard controls for navigation and marking
- Clear and simple interface

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.24.1 or newer installed

### Installation

Clone the repository:

```bash
git clone <repository-url>
cd go-minesweeper
```

## Running the Game

### Option 1: Run with Go

```bash
go run cmd/main.go
```

### Option 2: Build and Run Binary

Build the executable:

```bash
go build -o minesweeper cmd/main.go
```

Then run the binary:

```bash
./minesweeper
```

Or if you have the pre-built binary:

```bash
./main
```

## How to Play

1. Start a new game with: `start size=10 mines=15`
2. Navigate with: `move up/down/left/right`
3. Reveal tiles with: `reveal`
4. Mark/unmark mines with: `mark`
5. Quit with: `quit`

## Game Commands

- `start size=<number> mines=<number>` - Start a new game
- `move <direction>` - Move cursor (up, down, left, right)
- `reveal` - Reveal the current tile
- `mark` - Mark/unmark current tile as mine
- `quit` - Exit the game
