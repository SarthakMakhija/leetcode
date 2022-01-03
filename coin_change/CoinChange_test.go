package coin_change

import (
	"testing"
)

func TestMinimumCoins1(t *testing.T) {
	coins := coinChange([]int{1, 2, 5}, 11)
	expected := 3

	if coins != expected {
		t.Fatalf("Expected %v, received %v", expected, coins)
	}
}

func TestMinimumCoins2(t *testing.T) {
	coins := coinChange([]int{1, 2, 5}, 7)
	expected := 2

	if coins != expected {
		t.Fatalf("Expected %v, received %v", expected, coins)
	}
}

func TestMinimumCoins3(t *testing.T) {
	coins := coinChange([]int{1, 2, 5, 15}, 15)
	expected := 1

	if coins != expected {
		t.Fatalf("Expected %v, received %v", expected, coins)
	}
}

func TestMinimumCoins4(t *testing.T) {
	coins := coinChange([]int{2}, 3)
	expected := -1

	if coins != expected {
		t.Fatalf("Expected %v, received %v", expected, coins)
	}
}

func TestMinimumCoins5(t *testing.T) {
	coins := coinChange([]int{186, 419, 83, 408}, 6249)
	expected := 20

	if coins != expected {
		t.Fatalf("Expected %v, received %v", expected, coins)
	}
}
