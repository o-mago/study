package main

import (
	"flag"
	"main/analysis"
	"main/challenges/cards"
	"main/challenges/robot"
	"main/tools"
	"sync"
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

	// perf.StackInuse += m2.StackInuse
	// runtime.ReadMemStats(m)

	wg := new(sync.WaitGroup)
	if *file == "cards_slice" || *file == "cards" {
		tools.MeasurePerf("cards_slice", cards.CardsSlice)
	}

	if *file == "cards_linked_list" || *file == "cards" {
		tools.MeasurePerf("cards_linked_list", cards.CardsLinkedList)
	}

	if *file == "cards_container_list" || *file == "cards" {
		tools.MeasurePerf("cards_container_list", cards.CardsContainerList)
	}

	if *file == "cards_channel" || *file == "cards" {
		tools.MeasurePerf("cards_channel", cards.CardsChannel)
	}

	if *file == "cards_recursive" || *file == "cards" {
		tools.MeasurePerf("cards_recursive", cards.CardsRecursiveCall)
	}

	if *file == "cards_loop" || *file == "cards" {
		tools.MeasurePerf("cards_loop", cards.CardsSliceLoopCall)
	}

	wg.Wait()
}
