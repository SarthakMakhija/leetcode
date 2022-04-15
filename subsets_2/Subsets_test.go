package subsets_2

import (
	"fmt"
	"testing"
)

func TestSubsets1(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2}))
	fmt.Println(subsetsWithDup([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{0}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}
