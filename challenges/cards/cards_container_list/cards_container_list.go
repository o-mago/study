package cards_container_list

import (
	"container/list"
	"main/tools"
)

func Giveaway(cards int) ([]int, int) {
	totalLoops := tools.PerformanceParam{
		Name:  "total loops",
		Value: 0,
	}
	defer tools.Performance("cards_container_list", &totalLoops)()

	queue := list.New()

	giveawayCards := addCardsToQueue(cards, queue, &totalLoops)
	giveawayRemainingCards(queue, &giveawayCards, &totalLoops)

	remainingCard := queue.Front()

	return giveawayCards, remainingCard.Value.(int)
}

func addCardsToQueue(cards int, queue *list.List, totalLoops *tools.PerformanceParam) []int {
	giveawayCards := []int{}

	for i := 1; i <= cards; i += 2 {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}
		giveawayCards = append(giveawayCards, i)
		if i < cards {
			queue.PushBack(i + 1)
		}
	}

	return giveawayCards
}

func giveawayRemainingCards(queue *list.List, giveawayCards *[]int, totalLoops *tools.PerformanceParam) {
	for queue.Len() > 1 {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		givewayCard := queue.Front()
		*giveawayCards = append(*giveawayCards, givewayCard.Value.(int))
		queue.Remove(givewayCard)

		passCard := queue.Front()
		queue.PushBack(passCard.Value)
		queue.Remove(passCard)
	}
}
