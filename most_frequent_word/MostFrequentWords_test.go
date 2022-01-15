package most_frequent_word

import "testing"

func TestMostFrequentOccurringWord1(t *testing.T) {
	words := []string{"geeks", "for", "geeks"}
	word := mostFrequentWord(words)
	expected := "geeks"

	if expected != word {
		t.Fatalf("Expected %v, received %v", expected, word)
	}
}

func TestMostFrequentOccurringWord2(t *testing.T) {
	words := []string{"hello", "world"}
	word := mostFrequentWord(words)
	expected := "world"

	if expected != word {
		t.Fatalf("Expected %v, received %v", expected, word)
	}
}

func TestMostFrequentOccurringWord3(t *testing.T) {
	words := []string{"hello", "hello", "hello", "hello", "how"}
	word := mostFrequentWord(words)
	expected := "hello"

	if expected != word {
		t.Fatalf("Expected %v, received %v", expected, word)
	}
}

func TestMostFrequentOccurringWord4(t *testing.T) {
	words := []string{"hello", "bye", "hello", "hello", "bye", "bye"}
	word := mostFrequentWord(words)
	expected := "bye"

	if expected != word {
		t.Fatalf("Expected %v, received %v", expected, word)
	}
}

func TestMostFrequentOccurringWord5(t *testing.T) {
	words := []string{"bye", "hello", "bye", "hello", "bye", "hello"}
	word := mostFrequentWord(words)
	expected := "hello"

	if expected != word {
		t.Fatalf("Expected %v, received %v", expected, word)
	}
}

func TestMostFrequentOccurringWord6(t *testing.T) {
	words := []string{"hello", "bye", "hello", "bye"}
	word := mostFrequentWord(words)
	expected := "bye"

	if expected != word {
		t.Fatalf("Expected %v, received %v", expected, word)
	}
}
