package time_based_key_value_store

import (
	"testing"
)

func TestTimeBasedKVStore1(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	value := timeMap.Get("love", 5)
	expected := ""

	if expected != value {
		t.Fatalf("Expected %v, received %v", expected, value)
	}
}

func TestTimeBasedKVStore2(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	value := timeMap.Get("love", 10)
	expected := "high"

	if expected != value {
		t.Fatalf("Expected %v, received %v", expected, value)
	}
}

func TestTimeBasedKVStore3(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	value := timeMap.Get("love", 15)
	expected := "high"

	if expected != value {
		t.Fatalf("Expected %v, received %v", expected, value)
	}
}

func TestTimeBasedKVStore4(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	value := timeMap.Get("love", 25)
	expected := "low"

	if expected != value {
		t.Fatalf("Expected %v, received %v", expected, value)
	}
}
