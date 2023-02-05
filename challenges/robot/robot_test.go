package robot

import (
	"testing"
)

func TestMoveRobot_WrongCommand(t *testing.T) {
	commands := []string{"L", "R", "P", "3", "4"}
	expectedPosition := 0

	position := MoveRobot(commands)

	if expectedPosition != position {
		t.Errorf("expected: %d, got: %d", expectedPosition, position)
	}
}

func TestMoveRobot_Ok(t *testing.T) {
	commands := []string{"L", "R", "2", "L", "3"}
	expectedPosition := 1

	position := MoveRobot(commands)

	if expectedPosition != position {
		t.Errorf("expected: %d, got: %d", expectedPosition, position)
	}
}
