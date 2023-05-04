package cards

import (
	"main/structures/queue"
	"main/tools"
)

func CardsSlice(cards int, perf *tools.Performance) {
	Cards(cards, queue.NewSliceQueue[int](), perf)
}
