package file_system

import (
	"reflect"
	"testing"
)

func TestFileSystem1(t *testing.T) {
	fileSystem := Constructor()

	var creation []bool
	creation = append(creation, fileSystem.CreatePath("/leet", 1))
	creation = append(creation, fileSystem.CreatePath("/leet/code", 2))

	var get []int
	get = append(get, fileSystem.Get("/leet"))
	get = append(get, fileSystem.Get("/leet/code"))

	expectedCreation := []bool{true, true}
	expectedGet := []int{1, 2}

	if !reflect.DeepEqual(expectedCreation, creation) {
		t.Fatalf("Expected creations to be %v, received %v", expectedCreation, creation)
	}
	if !reflect.DeepEqual(expectedGet, get) {
		t.Fatalf("Expected gets to be %v, received %v", expectedGet, get)
	}
}

func TestFileSystem2(t *testing.T) {
	fileSystem := Constructor()

	var creation []bool
	creation = append(creation, fileSystem.CreatePath("/leet", 1))
	creation = append(creation, fileSystem.CreatePath("/leet/code", 2))
	creation = append(creation, fileSystem.CreatePath("/leet/code", 3))

	var get []int
	get = append(get, fileSystem.Get("/leet"))
	get = append(get, fileSystem.Get("/leet/code"))

	expectedCreation := []bool{true, true, false}
	expectedGet := []int{1, 2}

	if !reflect.DeepEqual(expectedCreation, creation) {
		t.Fatalf("Expected creations to be %v, received %v", expectedCreation, creation)
	}
	if !reflect.DeepEqual(expectedGet, get) {
		t.Fatalf("Expected gets to be %v, received %v", expectedGet, get)
	}
}

func TestFileSystem3(t *testing.T) {
	fileSystem := Constructor()

	var creation []bool
	creation = append(creation, fileSystem.CreatePath("/leet/code", 1))

	var get []int
	get = append(get, fileSystem.Get("/leet/code"))

	expectedCreation := []bool{false}
	expectedGet := []int{-1}

	if !reflect.DeepEqual(expectedCreation, creation) {
		t.Fatalf("Expected creations to be %v, received %v", expectedCreation, creation)
	}
	if !reflect.DeepEqual(expectedGet, get) {
		t.Fatalf("Expected gets to be %v, received %v", expectedGet, get)
	}
}
