package cards_linked_list

import (
	"main/tools"
)

func Giveaway(cards int) ([]int, int) {
	totalLoops := tools.PerformanceParam{
		Name:  "total loops",
		Value: 0,
	}
	defer tools.Performance("cards_linked_list", &totalLoops)()

	queue := Queue{}

	giveawayCards := addCardsToQueue(cards, &queue, &totalLoops)
	giveawayRemainingCards(&queue, &giveawayCards, &totalLoops)

	remainingCard := queue.dequeue()

	return giveawayCards, remainingCard
}

type node struct {
	value int
	next  *node
}

type Queue struct {
	head   *node
	tail   *node
	length int
}

func (q *Queue) enqueue(value int) {
	newNode := &node{
		value: value,
		next:  nil,
	}

	if q.tail != nil {
		q.tail.next = newNode
	} else if q.head == nil {
		q.head = newNode
	}

	q.tail = newNode
	q.length++
}

func (q *Queue) dequeue() int {
	if q.head != nil {
		if q.head == q.tail {
			q.tail = nil
		}
		oldHead := q.head
		q.head = oldHead.next
		q.length--
		return oldHead.value
	}

	return 0
}

func addCardsToQueue(cards int, queue *Queue, totalLoops *tools.PerformanceParam) []int {
	giveawayCards := []int{}

	for i := 1; i <= cards; i += 2 {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}
		giveawayCards = append(giveawayCards, i)
		if i < cards {
			queue.enqueue(i + 1)
		}
	}

	return giveawayCards
}

func giveawayRemainingCards(queue *Queue, giveawayCards *[]int, totalLoops *tools.PerformanceParam) {
	for queue.length > 1 {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		givewayCard := queue.dequeue()
		*giveawayCards = append(*giveawayCards, givewayCard)
		passCard := queue.dequeue()
		queue.enqueue(passCard)
	}
}
