package trie

import "testing"

func TestContains1(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	found := trie.Search("apple")
	if found != true {
		t.Fatalf("Expected apple to be found was not")
	}
}

func TestContains2(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("art")

	found := trie.Search("art")
	if found != true {
		t.Fatalf("Expected art to be found was not")
	}
}

func TestContains3(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("art")
	trie.Insert("amazing")

	found := trie.Search("amazing")
	if found != true {
		t.Fatalf("Expected amazing to be found was not")
	}
}

func TestContains4(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	found := trie.Search("apples")
	if found != false {
		t.Fatalf("Expected apples to not be found was")
	}
}

func TestStartsWith1(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("art")
	trie.Insert("amazing")

	found := trie.StartsWith("app")
	if found != true {
		t.Fatalf("Expected app to be found (starts with) was not")
	}
}

func TestStartsWith2(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("art")
	trie.Insert("amazing")

	found := trie.StartsWith("amaz")
	if found != true {
		t.Fatalf("Expected amaz to be found (starts with) was not")
	}
}

func TestStartsWith3(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("art")
	trie.Insert("amazing")

	found := trie.StartsWith("amaze")
	if found != false {
		t.Fatalf("Expected amaze to not be found (starts with) was")
	}
}
