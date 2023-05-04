package cards

import (
	"main/structures/queue"
	"main/tools"
)

func CardsChannel(cards int, perf *tools.Performance) {
	Cards(cards, queue.NewChannelQueue[int](cards/2), perf)
}
