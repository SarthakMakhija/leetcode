package remove_all_adjacent_duplicates_in_string_2

import (
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	result := removeDuplicates("abcd", 2)
	expected := "abcd"

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	result := removeDuplicates("deeedbbcccbdaa", 3)
	expected := "aa"

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestRemoveDuplicates3(t *testing.T) {
	result := removeDuplicates("pbbcggttciiippooaais", 2)
	expected := "ps"

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestRemoveDuplicates4(t *testing.T) {
	result := removeDuplicates("dtpdtaaaaaaaaappppppppppppppppppppaaaaaaaaaaxxxxxxxxxxxxxxsssssssssjjjjjjjjjjjjjjjjjjjjxxxxxxxxxxxxxxxxxxxxsssssssjjjjjjjjjjjjjjjjjjjjssssxxxxxxatdwvvpctpggggggggggggggggggggajagglaaaaaaaaaaaaaaaaaaaa", 20)
	expected := "dtpdttdwvvpctpajaggl"

	if expected != result {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
