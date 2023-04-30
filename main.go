package main

import (
	"flag"
	"main/analysis"
	"main/challenges/cards"
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
		cards.CardsSlice(100000)
	}

	if *file == "cards_linked_list" || *file == "cards" {
		cards.CardsLinkedList(100000)
	}

	if *file == "cards_container_list" || *file == "cards" {
		cards.CardsContainerList(100000)
	}

	if *file == "cards_channel" || *file == "cards" {
		cards.CardsChannel(100000)
	}

	if *file == "cards_recursive" || *file == "cards" {
		cards.CardsRecursive(100000, "recursive")
	}

	if *file == "cards_loop" || *file == "cards" {
		cards.CardsSliceLoop(100000, "loop")
	}
}
