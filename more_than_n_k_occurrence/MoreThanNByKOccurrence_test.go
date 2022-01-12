package more_than_n_k_occurrence

import (
	"reflect"
	"testing"
)

func TestMoreThanNByK1(t *testing.T) {
	result := moreThanNByK([]int{30, 10, 20, 20, 20, 10, 40, 30, 30}, 4)
	expected := []int{30, 20}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
