package cards

import (
	"main/structures/queue"
	"testing"
)

func TestCardsSlice(t *testing.T) {
	expectedRemoved := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 2, 6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 4, 12, 20, 28, 36, 44, 52, 16, 32, 48, 24, 8}
	expectedRemaining := 40

	removedCards, remainingCard := Cards(52, queue.NewSliceQueue[int](), "cards_slice")

	if expectedRemaining != remainingCard {
		t.Errorf("expected: %d, got: %d", expectedRemaining, remainingCard)
	}

	for index := range expectedRemoved {
		if expectedRemoved[index] != removedCards[index] {
			t.Errorf("expected: %d, got: %d", expectedRemoved, removedCards)
		}
	}
}

func TestCardsLinkedList(t *testing.T) {
	expectedRemoved := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 2, 6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 4, 12, 20, 28, 36, 44, 52, 16, 32, 48, 24, 8}
	expectedRemaining := 40

	removedCards, remainingCard := Cards(52, queue.NewLinkedListQueue[int](), "cards_linked_list")

	if expectedRemaining != remainingCard {
		t.Errorf("expected: %d, got: %d", expectedRemaining, remainingCard)
	}

	for index := range expectedRemoved {
		if expectedRemoved[index] != removedCards[index] {
			t.Errorf("expected: %d, got: %d", expectedRemoved, removedCards)
		}
	}
}

func TestCardsContainerList(t *testing.T) {
	expectedRemoved := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 2, 6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 4, 12, 20, 28, 36, 44, 52, 16, 32, 48, 24, 8}
	expectedRemaining := 40

	removedCards, remainingCard := Cards(52, queue.NewContainerListQueue[int](), "cards_container_list")

	if expectedRemaining != remainingCard {
		t.Errorf("expected: %d, got: %d", expectedRemaining, remainingCard)
	}

	for index := range expectedRemoved {
		if expectedRemoved[index] != removedCards[index] {
			t.Errorf("expected: %d, got: %d", expectedRemoved, removedCards)
		}
	}
}

func TestCardsChannel(t *testing.T) {
	expectedRemoved := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 2, 6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 4, 12, 20, 28, 36, 44, 52, 16, 32, 48, 24, 8}
	expectedRemaining := 40

	removedCards, remainingCard := Cards(52, queue.NewChannelQueue[int](52/2), "cards_channel")

	if expectedRemaining != remainingCard {
		t.Errorf("expected: %d, got: %d", expectedRemaining, remainingCard)
	}

	for index := range expectedRemoved {
		if expectedRemoved[index] != removedCards[index] {
			t.Errorf("expected: %d, got: %d", expectedRemoved, removedCards)
		}
	}
}
