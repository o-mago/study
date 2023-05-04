package cards

import "main/tools"

func CardsRecursiveCall(cards int, perf *tools.Performance) {
	CardsRecursive(cards, "recursive", perf)
}

type action bool

const REMOVE action = true
const MOVE action = false

func CardsRecursive(cards int, name string, perf *tools.Performance) ([]int, int) {
	perf.AddToParam("total loops", 0)

	cardsSlice := addCardsToSlice(cards, perf)
	removedCards, remainingCard := removeRemainingCardsRecursively(cardsSlice, []int{}, REMOVE, perf)

	return removedCards, remainingCard
}

func addCardsToSlice(cards int, perf *tools.Performance) []int {
	cardsSlice := []int{}

	for i := 1; i <= cards; i++ {
		perf.AddToParam("total loops", 1)

		cardsSlice = append(cardsSlice, i)
	}

	return cardsSlice
}

func removeRemainingCardsRecursively(cardsSlice []int, removedCards []int, act action, perf *tools.Performance) ([]int, int) {
	if len(cardsSlice) == 1 {
		return removedCards, cardsSlice[0]
	}

	remainingCards := []int{}

	for _, card := range cardsSlice {
		perf.AddToParam("total loops", 1)

		if act == REMOVE {
			removedCards = append(removedCards, card)
		} else {
			remainingCards = append(remainingCards, card)
		}

		act = !act
	}

	return removeRemainingCardsRecursively(remainingCards, removedCards, act, perf)
}
