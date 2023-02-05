package robot

import (
	"strconv"
)

func robot() {
	// commands := []string{"L", "R", "P", "3", "4"}
	// expectedPosition := 0

	// position := moveRobot(commands)
	// fmt.Println(position)
	// if expectedPosition != position {
	// 	log.Fatalf("expected: %d, got: %d", expectedPosition, position)
	// }
}

func MoveRobot(commands []string) int {
	position := 0

	for index, command := range commands {
		if command != "L" && command != "R" {
			intCommand, err := strconv.Atoi(command)
			if err != nil {
				commands[index] = "E"
				continue
			}
			commands[index] = commands[intCommand-1]
		}
		movement := getMovementByCommand(commands[index])
		position += movement
	}

	return position
}

func getMovementByCommand(command string) int {
	if command == "L" {
		return -1
	} else if command == "R" {
		return 1
	}
	return 0
}
