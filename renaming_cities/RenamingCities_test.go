package renaming_cities

import (
	"reflect"
	"testing"
)

func TestRenamingCities1(t *testing.T) {
	cities := []string{"shimla", "safari", "jammu", "delhi", "dehradun", "jammu"}
	uniqueCodes := rename(cities)
	expected := []string{"s", "sa", "j", "d", "deh", "jammu 2"}

	if !reflect.DeepEqual(expected, uniqueCodes) {
		t.Fatalf("Expected %v, received %v", expected, uniqueCodes)
	}
}

func TestRenamingCities2(t *testing.T) {
	cities := []string{"shimla", "safari", "jammu", "delhi", "dehradun", "jammu", "delhi"}
	uniqueCodes := rename(cities)
	expected := []string{"s", "sa", "j", "d", "deh", "jammu 2", "delhi 2"}

	if !reflect.DeepEqual(expected, uniqueCodes) {
		t.Fatalf("Expected %v, received %v", expected, uniqueCodes)
	}
}

func TestRenamingCities3(t *testing.T) {
	cities := []string{"shimla", "safari", "jammu", "delhi", "dehradun", "jammu", "jammu"}
	uniqueCodes := rename(cities)
	expected := []string{"s", "sa", "j", "d", "deh", "jammu 2", "jammu 3"}

	if !reflect.DeepEqual(expected, uniqueCodes) {
		t.Fatalf("Expected %v, received %v", expected, uniqueCodes)
	}
}

func TestRenamingCities4(t *testing.T) {
	cities := []string{
		"xqjjcwpdathcoxxixwrohpwlsejt",
		"frdyxpwwtspwtwqknidauzmegtovbdjydghmndpnuuaftlsrbv",
		"niposfmshchqdnnqegfhnbfrsshaarqasfdkqczshofbykbpoa",
		"umeoixvunvbkheiiaideidrqkkbsykkcsfazsujmysmoftfnek",
		"aempyjbxkfiz",
		"cqpetujyaawknkrqlgmkqesfxyeakvszdcjizueviw",
		"bfrisjqyvpuoekcmgzlaeuxyr",
		"msyaowyvmmhfyuwdolovmfzmnvkausdmaovaawkzphigka",
		"mylydwfctftyazvbmcmtbjeykoxjitgsjvjdplkzadkhcdedjm",
		"cehjdoxaygzgjfhxybztdbkncm",
		"jkdijhmuxamy",
		"dxullbpgbdyptuqpwtomwkxuaxquzmysyrtnpbvy",
		"mbjcrgesilbdtui",
		"xcibgedtvbhzdhhjaazjhqcu",
		"x",
		"ppunkvxnlnljvyuueorpgegnrphovqpzjezxjntkxmjubcwool",
		"joocimnykzeepctueezbqiviesmiwtldbcesojokskgzthburw",
		"gsyrkxhvyzlpjvgwyzzuqpolkbzicwklxtamovverzvmtbnhvb",
		"joknvcxytslwqlbkbnwjyrpxqmxrsymkeubqwsiplkmynspgiq",
		"hryakgbjzjnnoasjujhqvxjlibxzlojmkxxlsffofrxkasorby",
		"ecjbwmovtroervjuhuhlsgddpsgzebpezhiphqkcbsxwkteyvl",
		"iwemevcnjzgarbevwiihlxmyfmlqcprtsfnmcnpagoaewhohqy",
		"waetybmoezreq",
		"dtfrk",
		"zsvlcajfnfjridbhytwzhvzhwqwhtyfasvvfytnrnlnkvdminb",
		"ujomrilyludxkzerrlvjfnhecqakmiqiqsybcbvwsaydydonod",
		"nmdeglquqlsqhkyznxwbvkzhpbqkeamdptuoxcjxidfcnvlmek",
		"azxeewqryieddyqahwfafoogbhjalxdjub",
		"brjqre",
		"gkvhlrxx"}

	uniqueCodes := rename(cities)
	expected := []string{"x", "f", "n", "u", "a", "c", "b", "m", "my", "ce", "j", "d", "mb", "xc", "x", "p", "jo", "g", "jok", "h", "e", "i", "w",
		"dt", "z", "uj", "nm", "az", "br", "gk"}

	if !reflect.DeepEqual(expected, uniqueCodes) {
		t.Fatalf("Expected %v, received %v", expected, uniqueCodes)
	}
}
