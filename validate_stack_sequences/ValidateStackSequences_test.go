package validate_stack_sequences

import (
	"testing"
)

func TestValidateStackSequences1(t *testing.T) {
	validStackSequences := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
	if validStackSequences != true {
		t.Fatalf("Expected stack sequences to be valid but were not")
	}
}

func TestValidateStackSequences2(t *testing.T) {
	validStackSequences := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{3, 2, 1, 4, 5})
	if validStackSequences != true {
		t.Fatalf("Expected stack sequences to be valid but were not")
	}
}

func TestValidateStackSequences3(t *testing.T) {
	validStackSequences := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{3, 2, 1, 5, 4})
	if validStackSequences != true {
		t.Fatalf("Expected stack sequences to be valid but were not")
	}
}

func TestValidateStackSequences4(t *testing.T) {
	validStackSequences := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{3, 2, 1, 4, 5})
	if validStackSequences != true {
		t.Fatalf("Expected stack sequences to be valid but were not")
	}
}

func TestValidateStackSequences5(t *testing.T) {
	validStackSequences := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2})
	if validStackSequences != false {
		t.Fatalf("Expected stack sequences to be invalid but were valid")
	}
}
