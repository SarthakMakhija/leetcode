package camel_case

import (
	"reflect"
	"sort"
	"testing"
)

func TestAllCamelCaseWords1(t *testing.T) {
	dictionary := []string{"WeTl", "Wol"}
	matchingWords := findAllWords(dictionary, "WT")
	expected := []string{"WeTl"}

	if !reflect.DeepEqual(expected, matchingWords) {
		t.Fatalf("Expected %v, received %v", expected, matchingWords)
	}
}

func TestAllCamelCaseWords2(t *testing.T) {
	dictionary := []string{"Alg", "AlgH"}
	matchingWords := findAllWords(dictionary, "AH")
	expected := []string{"AlgH"}

	if !reflect.DeepEqual(expected, matchingWords) {
		t.Fatalf("Expected %v, received %v", expected, matchingWords)
	}
}

func TestAllCamelCaseWords3(t *testing.T) {
	dictionary := []string{"WelcomeGeek", "WelcomeToGeeksForGeeks", "GeeksForGeeks"}
	matchingWords := findAllWords(dictionary, "WTG")
	expected := []string{"WelcomeToGeeksForGeeks"}

	if !reflect.DeepEqual(expected, matchingWords) {
		t.Fatalf("Expected %v, received %v", expected, matchingWords)
	}
}

func TestAllCamelCaseWords4(t *testing.T) {
	dictionary := []string{"WelcomeGeek", "WelcomeToGeeksForGeeks", "GeeksForGeeks", "WelcomeToGeeksForGeeksAgain"}
	matchingWords := findAllWords(dictionary, "WTG")
	expected := []string{"WelcomeToGeeksForGeeks", "WelcomeToGeeksForGeeksAgain"}

	if !reflect.DeepEqual(expected, matchingWords) {
		t.Fatalf("Expected %v, received %v", expected, matchingWords)
	}
}
func TestAllCamelCaseWords5(t *testing.T) {
	dictionary := []string{"Hi", "Hello", "HelloWorld", "HiTech", "HiGeek", "HiTechWorld", "HiTechCity", "HiTechLab"}
	matchingWords := findAllWords(dictionary, "HA")
	var expected []string

	if !reflect.DeepEqual(expected, matchingWords) {
		t.Fatalf("Expected %v, received %v", expected, matchingWords)
	}
}

func TestAllCamelCaseWords6(t *testing.T) {
	dictionary := []string{"Hi", "Hello", "HelloWorld", "HiTech", "HiGeek", "HiTechWorld", "HiTechCity", "HiTechLab"}
	matchingWords := findAllWords(dictionary, "H")

	sort.Strings(matchingWords)
	expected := []string{"Hello", "HelloWorld", "Hi", "HiGeek", "HiTech", "HiTechCity", "HiTechLab", "HiTechWorld"}

	if !reflect.DeepEqual(expected, matchingWords) {
		t.Fatalf("Expected %v, received %v", expected, matchingWords)
	}
}

func TestAllCamelCaseWords7(t *testing.T) {
	dictionary := []string{"Hello", "Hello", "Hello", "Hello"}
	matchingWords := findAllWords(dictionary, "H")

	sort.Strings(matchingWords)
	expected := []string{"Hello", "Hello", "Hello", "Hello"}

	if !reflect.DeepEqual(expected, matchingWords) {
		t.Fatalf("Expected %v, received %v", expected, matchingWords)
	}
}
