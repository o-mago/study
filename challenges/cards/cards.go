package cards

import (
	"log"
	"main/structures/queue"
	"main/tools"
)

func Cards(cards int, q queue.Queue[int], name string) ([]int, int) {
	totalLoops := tools.PerformanceParam{
		Name:  "total loops",
		Value: 0,
	}
	defer tools.Performance(name, &totalLoops)()

	removedCards := addCardsToQueue(cards, q, &totalLoops)
	removeRemainingCards(q, &removedCards, &totalLoops)

	remainingCard, err := q.Dequeue()
	if err != nil {
		log.Fatal(err)
	}

	return removedCards, remainingCard
}

func addCardsToQueue(cards int, queue queue.Queue[int], totalLoops *tools.PerformanceParam) []int {
	removedCards := []int{}

	for i := 1; i <= cards; i += 2 {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}
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

func removeRemainingCards(q queue.Queue[int], removedCards *[]int, totalLoops *tools.PerformanceParam) {
	for q.Len() > 1 {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		givewayCard, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}
		*removedCards = append(*removedCards, givewayCard)
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
