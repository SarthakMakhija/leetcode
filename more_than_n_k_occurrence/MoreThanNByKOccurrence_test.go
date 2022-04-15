package more_than_n_k_occurrence

import (
	"reflect"
	"sort"
	"testing"
)

func TestMoreThanNByK1(t *testing.T) {
	result := moreThanNByK([]int{30, 10, 20, 20, 20, 10, 40, 30, 30}, 4)
	expected := []int{20, 30}

	sort.Ints(result)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
