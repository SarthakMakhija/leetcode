package valid_parenthesis_string

import (
	"testing"
)

func TestValidParenthesis1(t *testing.T) {
	if checkValidString("()") != true {
		t.Fatalf("Expected () to be valid but was not")
	}
}

func TestValidParenthesis2(t *testing.T) {
	if checkValidString("(*") != true {
		t.Fatalf("Expected (* to be valid but was not")
	}
}

func TestValidParenthesis3(t *testing.T) {
	if checkValidString("****") != true {
		t.Fatalf("Expected **** to be valid but was not")
	}
}

func TestValidParenthesis4(t *testing.T) {
	if checkValidString("(*)") != true {
		t.Fatalf("Expected (*) to be valid but was not")
	}
}

func TestValidParenthesis5(t *testing.T) {
	if checkValidString("(**") != true {
		t.Fatalf("Expected (** to be valid but was not")
	}
}

func TestValidParenthesis6(t *testing.T) {
	if checkValidString(")**") != false {
		t.Fatalf("Expected )** to be invalid but was valid")
	}
}

func TestValidParenthesis7(t *testing.T) {
	if checkValidString("((*") != false {
		t.Fatalf("Expected ((* to be invalid but was valid")
	}
}

func TestValidParenthesis8(t *testing.T) {
	if checkValidString("((**") != true {
		t.Fatalf("Expected ((** to be valid but was invalid")
	}
}

func TestValidParenthesis(t *testing.T) {
	if checkValidString("(*))") != true {
		t.Fatalf("Expected (*)) to be valid but was invalid")
	}
}
