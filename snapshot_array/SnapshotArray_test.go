package snapshot_array

import (
	"testing"
)

func TestSnapshot1(t *testing.T) {
	snapshotArray := Constructor(1)
	snapshotArray.Set(0, 15)
	snapshotArray.Snap()
	snapshotArray.Snap()
	snapshotArray.Snap()

	result := snapshotArray.Get(0, 2)
	expected := 15
	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
	snapshotArray.Snap()
	snapshotArray.Snap()

	result = snapshotArray.Get(0, 0)
	expected = 15
	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSnapshot2(t *testing.T) {
	snapshotArray := Constructor(3)
	snapshotArray.Set(0, 5)
	snapshotArray.Snap()
	snapshotArray.Set(0, 6)

	result := snapshotArray.Get(0, 0)
	expected := 5
	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}

func TestSnapshot3(t *testing.T) {
	snapshotArray := Constructor(3)
	snapshotArray.Set(1, 18)
	snapshotArray.Set(1, 4)
	snapshotArray.Snap()

	result := snapshotArray.Get(0, 0)
	expected := 0
	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
	snapshotArray.Set(0, 20)
	snapshotArray.Snap()
	snapshotArray.Set(0, 2)
	snapshotArray.Set(1, 1)

	result = snapshotArray.Get(1, 1)
	expected = 4
	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}

	result = snapshotArray.Get(1, 0)
	expected = 4
	if result != expected {
		t.Fatalf("Expected %v, received %v", expected, result)
	}
}
