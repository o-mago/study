package cards

import (
	"main/structures/queue"
)

func CardsLinkedList(cards int) ([]int, int) {
	return Cards(cards, queue.NewLinkedListQueue[int](), "cards_linked_list")
}
