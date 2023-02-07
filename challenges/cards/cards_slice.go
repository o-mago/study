package cards

import (
	"main/structures/queue"
)

func CardsSlice(cards int) ([]int, int) {
	return Cards(cards, queue.NewSliceQueue[int](), "cards_slice")
}
