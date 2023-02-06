package cards_container_list

import (
	"testing"
)

func TestGiveaway(t *testing.T) {
	expectedGiveaway := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 2, 6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 4, 12, 20, 28, 36, 44, 52, 16, 32, 48, 24, 8}
	expectedRemaining := 40

	givewayCards, remainingCard := Giveaway(52)

	if expectedRemaining != remainingCard {
		t.Errorf("expected: %d, got: %d", expectedRemaining, remainingCard)
	}

	for index := range expectedGiveaway {
		if expectedGiveaway[index] != givewayCards[index] {
			t.Errorf("expected: %d, got: %d", expectedGiveaway, givewayCards)
		}
	}
}
