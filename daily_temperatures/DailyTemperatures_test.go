package daily_temperatures

import (
	"reflect"
	"testing"
)

func TestWaitForWarmerDays1(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	wait := dailyTemperatures(temperatures)
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}

	if !reflect.DeepEqual(expected, wait) {
		t.Fatalf("Expecyed %v, received %v", expected, wait)
	}
}

func TestWaitForWarmerDays2(t *testing.T) {
	temperatures := []int{30, 40, 50, 60}
	wait := dailyTemperatures(temperatures)
	expected := []int{1, 1, 1, 0}

	if !reflect.DeepEqual(expected, wait) {
		t.Fatalf("Expecyed %v, received %v", expected, wait)
	}
}

func TestWaitForWarmerDays3(t *testing.T) {
	temperatures := []int{30, 60, 90}
	wait := dailyTemperatures(temperatures)
	expected := []int{1, 1, 0}

	if !reflect.DeepEqual(expected, wait) {
		t.Fatalf("Expecyed %v, received %v", expected, wait)
	}
}
