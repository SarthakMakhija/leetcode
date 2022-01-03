package word_search

import (
	"testing"
)

func TestWordSearch1(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exists := exist(board, "ABCCED")

	if exists != true {
		t.Fatalf("Expected ABCCED to exist but did not")
	}
}

func TestWordSearch2(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exists := exist(board, "SEE")

	if exists != true {
		t.Fatalf("Expected SEE to exist but did not")
	}
}

func TestWordSearch3(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exists := exist(board, "SAF")

	if exists != false {
		t.Fatalf("Expected SAF to not exist but it did")
	}
}

func TestWordSearch4(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exists := exist(board, "ABCB")

	if exists != false {
		t.Fatalf("Expected ABCB to not exist but it did")
	}
}

func TestWordSearch5(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'A', 'D', 'C', 'S'},
		{'S', 'F', 'E', 'E'},
	}
	exists := exist(board, "SAD")

	if exists != true {
		t.Fatalf("Expected SAD to exist but it did not")
	}
}

func TestWordSearch6(t *testing.T) {
	board := [][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}
	exists := exist(board, "aaaaaaaaaaaaa")

	if exists != false {
		t.Fatalf("Expected aaaaaaaaaaaaa to not exist but it did")
	}
}

func TestWordSearch7(t *testing.T) {
	board := [][]byte{
		{'a'},
	}
	exists := exist(board, "a")

	if exists != true {
		t.Fatalf("Expected a to exist but it did not")
	}
}

func TestWordSearch8(t *testing.T) {
	board := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	exists := exist(board, "AAB")

	if exists != true {
		t.Fatalf("Expected AAB to exist but it did not")
	}
}

func TestWordSearch9(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exists := exist(board, "ABCESEEEFS")

	if exists != true {
		t.Fatalf("Expected ABCESEEEFS to exist but it did not")
	}
}
