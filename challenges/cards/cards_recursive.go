package cards

import (
	"main/tools"
)

type action bool

const REMOVE action = true
const MOVE action = false

func CardsRecursive(cards int, name string) ([]int, int) {
	totalLoops := tools.PerformanceParam{
		Name:  "total loops",
		Value: 0,
	}
	defer tools.Performance(name, &totalLoops)()

	cardsSlice := addCardsToSlice(cards, &totalLoops)
	removedCards, remainingCard := removeRemainingCardsRecursively(cardsSlice, []int{}, REMOVE, &totalLoops)

	return removedCards, remainingCard
}

func addCardsToSlice(cards int, totalLoops *tools.PerformanceParam) []int {
	cardsSlice := []int{}

	for i := 1; i <= cards; i++ {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		cardsSlice = append(cardsSlice, i)
	}

	return cardsSlice
}

func removeRemainingCardsRecursively(cardsSlice []int, removedCards []int, act action, totalLoops *tools.PerformanceParam) ([]int, int) {
	if len(cardsSlice) == 1 {
		return removedCards, cardsSlice[0]
	}

	remainingCards := []int{}

	for _, card := range cardsSlice {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		if act == REMOVE {
			removedCards = append(removedCards, card)
		} else {
			remainingCards = append(remainingCards, card)
		}

		act = !act
	}

	return removeRemainingCardsRecursively(remainingCards, removedCards, act, totalLoops)
}
