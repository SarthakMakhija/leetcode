package search_suggestions_system

import (
	"reflect"
	"testing"
)

func TestSuggestedProducts1(t *testing.T) {
	suggestedProducts := suggestedProducts(
		[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
		"mouse",
	)
	expected := [][]string{
		{"mobile", "moneypot", "monitor"},
		{"mobile", "moneypot", "monitor"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
	}

	if !reflect.DeepEqual(expected, suggestedProducts) {
		t.Fatalf("Expected %v, received %v", expected, suggestedProducts)
	}
}

func TestSuggestedProducts2(t *testing.T) {
	suggestedProducts := suggestedProducts(
		[]string{"bags", "baggage", "banner", "box", "cloths"},
		"bags",
	)
	expected := [][]string{
		{"baggage", "bags", "banner"},
		{"baggage", "bags", "banner"},
		{"baggage", "bags"},
		{"bags"},
	}

	if !reflect.DeepEqual(expected, suggestedProducts) {
		t.Fatalf("Expected %v, received %v", expected, suggestedProducts)
	}
}

func TestSuggestedProducts3(t *testing.T) {
	suggestedProducts := suggestedProducts(
		[]string{"hostel"},
		"house",
	)
	expected := [][]string{
		{"hostel"},
		{"hostel"},
		{},
		{},
		{},
	}

	if !reflect.DeepEqual(expected, suggestedProducts) {
		t.Fatalf("Expected %v, received %v", expected, suggestedProducts)
	}
}
