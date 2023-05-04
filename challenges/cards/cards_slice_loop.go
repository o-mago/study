package cards

import "main/tools"

type actionLoop bool

const REMOVE_LOOP actionLoop = true
const MOVE_LOOP actionLoop = false

func CardsSliceLoopCall(cards int, perf *tools.Performance) {
	CardsSliceLoop(cards, "loop", perf)
}

func CardsSliceLoop(cards int, name string, perf *tools.Performance) ([]int, int) {
	perf.AddToParam("total loops", 0)

	cardsSlice := addCardsToSliceLoop(cards, perf)
	removedCards, remainingCard := removeRemainingCardsLoop(cardsSlice, []int{}, REMOVE_LOOP, perf)

	return removedCards, remainingCard
}

func addCardsToSliceLoop(cards int, perf *tools.Performance) []int {
	cardsSlice := []int{}

	for i := 1; i <= cards; i++ {
		perf.AddToParam("total loops", 1)

		cardsSlice = append(cardsSlice, i)
	}

	return cardsSlice
}

func removeRemainingCardsLoop(cardsSlice []int, removedCards []int, act actionLoop, perf *tools.Performance) ([]int, int) {
	remainingCards := []int{}

	for len(cardsSlice) > 1 {
		for _, card := range cardsSlice {
			perf.AddToParam("total loops", 1)

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
