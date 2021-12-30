package word_dictionary

import (
	"testing"
)

func TestWordSearch1(t *testing.T) {
	dictionary := Constructor()
	dictionary.AddWord("bad")
	found := dictionary.Search("bad")

	if found != true {
		t.Fatalf("Expected bad to be found but was not")
	}
}

func TestWordSearch2(t *testing.T) {
	dictionary := Constructor()
	dictionary.AddWord("card")

	found := dictionary.Search("card")

	if found != true {
		t.Fatalf("Expected card to be found but was not")
	}
}

func TestWordSearch3(t *testing.T) {
	dictionary := Constructor()
	dictionary.AddWord("bad")
	dictionary.AddWord("mad")

	found := dictionary.Search(".ad")
	if found != true {
		t.Fatalf("Expected .ad to be found but was not")
	}
}

func TestWordSearch4(t *testing.T) {
	dictionary := Constructor()
	dictionary.AddWord("bad")
	dictionary.AddWord("mad")

	found := dictionary.Search("b..")
	if found != true {
		t.Fatalf("Expected b.. to be found but was not")
	}
}

func TestWordSearch5(t *testing.T) {
	dictionary := Constructor()
	dictionary.AddWord("bad")
	dictionary.AddWord("mad")

	found := dictionary.Search("b...")
	if found != false {
		t.Fatalf("Expected b... to not be found but was")
	}
}
