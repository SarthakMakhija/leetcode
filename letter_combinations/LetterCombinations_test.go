package letter_combinations

import (
	"reflect"
	"testing"
)

func TestLetterCombinations1(t *testing.T) {
	combinations := letterCombinations("23")
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	if !reflect.DeepEqual(expected, combinations) {
		t.Fatalf("Expected %v, received %v", expected, combinations)
	}
}

func TestLetterCombinations2(t *testing.T) {
	combinations := letterCombinations("2")
	expected := []string{"a", "b", "c"}

	if !reflect.DeepEqual(expected, combinations) {
		t.Fatalf("Expected %v, received %v", expected, combinations)
	}
}

func TestLetterCombinations3(t *testing.T) {
	combinations := letterCombinations("234")
	expected := []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}

	if !reflect.DeepEqual(expected, combinations) {
		t.Fatalf("Expected %v, received %v", expected, combinations)
	}
}
