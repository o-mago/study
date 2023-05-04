package cards

import (
	"main/structures/queue"
	"main/tools"
)

func CardsLinkedList(cards int, perf *tools.Performance) {
	Cards(cards, queue.NewLinkedListQueue[int](), perf)
}
