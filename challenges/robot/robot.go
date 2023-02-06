package robot

import (
	"strconv"
)

func MoveRobot(commands []string) int {
	position := 0

	for index := range commands {
		acceptedCommand, ok := getAcceptedCommand(commands, index)
		commands[index] = acceptedCommand
		if !ok {
			continue
		}

		movement := getMovementByCommand(commands[index])
		position += movement
	}

	return position
}

func getAcceptedCommand(commands []string, index int) (string, bool) {
	if commands[index] != "L" && commands[index] != "R" {
		intCommand, err := strconv.Atoi(commands[index])
		if err != nil || intCommand > index+1 {
			return "E", false
		}
		return commands[intCommand-1], true
	}

	return commands[index], true
}

func getMovementByCommand(command string) int {
	if command == "L" {
		return -1
	} else if command == "R" {
		return 1
	}
	return 0
}
