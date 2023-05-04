package cards

import (
	"log"
	"main/structures/queue"
	"main/tools"
)

func Cards(cards int, q queue.Queue[int], perf *tools.Performance) ([]int, int) {
	perf.AddToParam("total loops", 0)

	removedCards := addCardsToQueue(cards, q, perf)
	removeRemainingCards(q, &removedCards, perf)

	remainingCard, err := q.Dequeue()
	if err != nil {
		log.Fatal(err)
	}

	return removedCards, remainingCard
}

func addCardsToQueue(cards int, queue queue.Queue[int], perf *tools.Performance) []int {
	removedCards := []int{}

	for i := 1; i <= cards; i += 2 {
		perf.AddToParam("total loops", 1)

		removedCards = append(removedCards, i)
		if i < cards {
			err := queue.Enqueue(i + 1)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return removedCards
}

func removeRemainingCards(q queue.Queue[int], removedCards *[]int, perf *tools.Performance) {
	for q.Len() > 1 {
		perf.AddToParam("total loops", 1)

		removeCard, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}
		*removedCards = append(*removedCards, removeCard)
		passCard, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}

		err = q.Enqueue(passCard)
		if err != nil {
			log.Fatal(err)
		}
	}
}
