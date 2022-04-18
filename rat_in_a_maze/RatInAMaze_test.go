package rat_in_a_maze

import (
	"reflect"
	"testing"
)

func TestFindPath1(t *testing.T) {
	maze := [][]int{
		{
			1, 0, 0, 0,
		},
		{
			1, 1, 0, 1,
		},
		{
			1, 1, 0, 0,
		},
		{
			0, 1, 1, 1,
		},
	}

	paths := findPath(maze, 4)
	expected := []string{"DDRDRR", "DRDDRR"}

	if !reflect.DeepEqual(expected, paths) {
		t.Fatalf("Expected %v, received %v", expected, paths)
	}
}

func TestFindPath2(t *testing.T) {
	maze := [][]int{
		{1, 0},
		{1, 0},
	}
	paths := findPath(maze, 2)

	if len(paths) != 0 {
		t.Fatalf("Expected 0 paths but received %v", paths)
	}
}

func TestFindPath3(t *testing.T) {
	maze := [][]int{
		{1, 0, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	paths := findPath(maze, 3)
	expected := []string{"DDRURD", "DDRR", "DRDR", "DRRD"}

	if !reflect.DeepEqual(expected, paths) {
		t.Fatalf("Expected %v, received %v", expected, paths)
	}
}
