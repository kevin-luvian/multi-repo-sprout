package game

import (
	"testing"
)

func TestCommand_ParseCommand(t *testing.T) {
	testCases := []struct {
		name string
		args func() string
		test func(*testing.T, Command, error)
	}{
		{
			name: "Test parse start command",
			args: func() string {
				return "start 10 10"
			},
			test: func(t *testing.T, command Command, err error) {
				if command.Command != "start" {
					t.Errorf("Command should be start, got %s", command.Command)
				}

				if command.Args["size"] != "10" {
					t.Errorf("Size should be 10, got %s", command.Args["size"])
				}

				if command.Args["mines"] != "10" {
					t.Errorf("Mines should be 10, got %s", command.Args["mines"])
				}
			},
		},
		{
			name: "Test parse start command with invalid size",
			args: func() string {
				return "start 1 10"
			},
			test: func(t *testing.T, command Command, err error) {
				if command.Command != "" {
					t.Errorf("Command should be empty, got %s", command.Command)
				}

				if err == nil {
					t.Errorf("Error should be not nil")
				}
			},
		},
		{
			name: "Test parse reveal command",
			args: func() string {
				return "reveal 0 0"
			},
			test: func(t *testing.T, command Command, err error) {
				if command.Command != "reveal" {
					t.Errorf("Command should be reveal, got %s", command.Command)
				}
			},
		},
		{
			name: "Test parse reveal command with invalid format",
			args: func() string {
				return "reveal 0"
			},
			test: func(t *testing.T, command Command, err error) {
				if command.Command != "" {
					t.Errorf("Command should be empty, got %s", command.Command)
				}

				if err == nil {
					t.Errorf("Error should be not nil")
				}
			},
		},
		{
			name: "Test parse reveal command with invalid integer",
			args: func() string {
				return "reveal a b"
			},
			test: func(t *testing.T, command Command, err error) {
				if command.Command != "" {
					t.Errorf("Command should be empty, got %s", command.Command)
				}

				if err == nil {
					t.Errorf("Error should be not nil")
				}
			},
		},
		{
			name: "Test parse reveal command with out of bounds",
			args: func() string {
				return "reveal 1000 1000"
			},
			test: func(t *testing.T, command Command, err error) {
				if command.Command != "" {
					t.Errorf("Command should be empty, got %s", command.Command)
				}

				if err.Error() != "out of bounds" {
					t.Errorf("Error should be out of bounds, got %v", err)
				}
			},
		},
		{
			name: "Test parse reveal command with invalid action",
			args: func() string {
				return "no_action"
			},
			test: func(t *testing.T, command Command, err error) {
				if err == nil {
					t.Errorf("Error should be not nil")
				}
			},
		},
	}

	game := NewGame(NewGameOptions{
		Size:  10,
		Mines: 10,
	})

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			command, err := ParseCommand(testCase.args(), game)
			testCase.test(t, command, err)
		})
	}
}
