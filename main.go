package main

import (
	"flag"
	"main/analysis"
	"main/challenges/cards/cards_container_list"
	"main/challenges/cards/cards_linked_list"
	"main/challenges/cards/cards_slice"
	"main/challenges/robot"
)

func main() {
	file := flag.String("file", "", "")
	flag.Parse()

	if *file == "slices" {
		analysis.Slices()
	}

	if *file == "robot" {
		robot.MoveRobot([]string{"L", "R", "P", "3", "4"})
	}

	if *file == "cards_slice" || *file == "cards" {
		cards_slice.Giveaway(1000)
	}

	if *file == "cards_linked_list" || *file == "cards" {
		cards_linked_list.Giveaway(100000)
	}

	if *file == "cards_container_list" || *file == "cards" {
		cards_container_list.Giveaway(100000)
	}
}
