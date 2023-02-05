package robot

import (
	"strconv"
)

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
