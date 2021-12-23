package copy_list_random_pointer

import (
	"reflect"
	"testing"
)

func TestCopy1(t *testing.T) {
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}

	node1.Next = node2
	node1.Random = nil

	node2.Next = node3
	node2.Random = node1

	node3.Next = node4
	node3.Random = node5

	node4.Next = node5
	node4.Random = node3

	node5.Next = nil
	node5.Random = node1

	copied := copyRandomList(node1)

	values := copied.all()
	expected := []ListWithRandomPointerValue{
		{Val: 7, RandomVal: -1},
		{Val: 13, RandomVal: 7},
		{Val: 11, RandomVal: 1},
		{Val: 10, RandomVal: 11},
		{Val: 1, RandomVal: 7},
	}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestCopy2(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}

	node1.Next = node2
	node1.Random = node2

	node2.Next = nil
	node2.Random = node2

	copied := copyRandomList(node1)

	values := copied.all()
	expected := []ListWithRandomPointerValue{
		{Val: 1, RandomVal: 2},
		{Val: 2, RandomVal: 2},
	}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestCopy3(t *testing.T) {
	node1 := &Node{Val: 3}
	node2 := &Node{Val: 3}
	node3 := &Node{Val: 3}

	node1.Next = node2
	node1.Random = nil

	node2.Next = node3
	node2.Random = node1

	node3.Next = nil
	node3.Random = nil

	copied := copyRandomList(node1)

	values := copied.all()
	expected := []ListWithRandomPointerValue{
		{Val: 3, RandomVal: -1},
		{Val: 3, RandomVal: 3},
		{Val: 3, RandomVal: -1},
	}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
