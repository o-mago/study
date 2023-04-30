package cards

import (
	"main/tools"
)

type actionLoop bool

const REMOVE_LOOP actionLoop = true
const MOVE_LOOP actionLoop = false

func CardsSliceLoop(cards int, name string) ([]int, int) {
	totalLoops := tools.PerformanceParam{
		Name:  "total loops",
		Value: 0,
	}
	defer tools.Performance(name, &totalLoops)()

	cardsSlice := addCardsToSliceLoop(cards, &totalLoops)
	removedCards, remainingCard := removeRemainingCardsLoop(cardsSlice, []int{}, REMOVE_LOOP, &totalLoops)

	return removedCards, remainingCard
}

func addCardsToSliceLoop(cards int, totalLoops *tools.PerformanceParam) []int {
	cardsSlice := []int{}

	for i := 1; i <= cards; i++ {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}
		cardsSlice = append(cardsSlice, i)
	}

	return cardsSlice
}

func removeRemainingCardsLoop(cardsSlice []int, removedCards []int, act actionLoop, totalLoops *tools.PerformanceParam) ([]int, int) {
	remainingCards := []int{}

	for len(cardsSlice) > 1 {
		for _, card := range cardsSlice {
			if totalLoopsInt, ok := totalLoops.Value.(int); ok {
				totalLoops.Value = totalLoopsInt + 1
			}

			if act == REMOVE_LOOP {
				removedCards = append(removedCards, card)
			} else {
				remainingCards = append(remainingCards, card)
			}

			act = !act
		}

		cardsSlice = remainingCards
		remainingCards = []int{}
	}

	return removedCards, cardsSlice[0]
}
