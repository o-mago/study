package cards_slice

import (
	"fmt"
	"main/tools"
)

type Queue []int

func Giveaway(cards int) ([]int, int) {
	totalLoops := tools.PerformanceParam{
		Name:  "total loops",
		Value: 0,
	}
	defer tools.Performance("cards_slice", &totalLoops)()

	queue := Queue{}

	giveawayCards := addCardsToQueue(cards, &queue, &totalLoops)
	fmt.Println(queue)
	giveawayRemainingCards(&queue, &giveawayCards, &totalLoops)

	remainderCard := queue.dequeue()

	return giveawayCards, remainderCard
}

func (q *Queue) enqueue(value int) {
	*q = append(*q, value)
}

func (q *Queue) dequeue() int {
	queue := *q
	if len(*q) > 0 {
		card := queue[0]
		*q = queue[1:]
		return card
	}

	return 0
}

func addCardsToQueue(cards int, queue *Queue, totalLoops *tools.PerformanceParam) []int {
	giveawayCards := []int{}

	for i := 1; i <= cards; i++ {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		if i%2 == 1 {
			giveawayCards = append(giveawayCards, i)
		} else {
			queue.enqueue(i)
		}
	}

	return giveawayCards
}

func giveawayRemainingCards(queue *Queue, giveawayCards *[]int, totalLoops *tools.PerformanceParam) {
	for len(*queue) > 1 {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		givewayCard := queue.dequeue()
		*giveawayCards = append(*giveawayCards, givewayCard)
		passCard := queue.dequeue()
		queue.enqueue(passCard)
	}
}
