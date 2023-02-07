package cards

import (
	"main/structures/queue"
)

func CardsContainerList(cards int) ([]int, int) {
	return Cards(cards, queue.NewContainerListQueue[int](), "cards_container_list")
}
