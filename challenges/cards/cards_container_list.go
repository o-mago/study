package cards

import (
	"main/structures/queue"
	"main/tools"
)

func CardsContainerList(cards int, perf *tools.Performance) {
	Cards(cards, queue.NewContainerListQueue[int](), perf)
}
