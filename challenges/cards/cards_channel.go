package cards

import (
	"main/structures/queue"
)

func CardsChannel(cards int) ([]int, int) {
	return Cards(cards, queue.NewChannelQueue[int](cards), "cards_channel")
}
