package simplify_path

import "testing"

func TestSimplifyPath1(t *testing.T) {
	canonicalPath := simplifyPath("/home/")
	expected := "/home"

	if canonicalPath != expected {
		t.Fatalf("Expected %v, received %v", expected, canonicalPath)
	}
}

func TestSimplifyPath2(t *testing.T) {
	canonicalPath := simplifyPath("/../")
	expected := "/"

	if canonicalPath != expected {
		t.Fatalf("Expected %v, received %v", expected, canonicalPath)
	}
}

func TestSimplifyPath3(t *testing.T) {
	canonicalPath := simplifyPath("/home//foo/")
	expected := "/home/foo"

	if canonicalPath != expected {
		t.Fatalf("Expected %v, received %v", expected, canonicalPath)
	}
}

func TestSimplifyPath4(t *testing.T) {
	canonicalPath := simplifyPath("/home/../foo/")
	expected := "/foo"

	if canonicalPath != expected {
		t.Fatalf("Expected %v, received %v", expected, canonicalPath)
	}
}

func TestSimplifyPath5(t *testing.T) {
	canonicalPath := simplifyPath("/home//foo/bar/../")
	expected := "/home/foo"

	if canonicalPath != expected {
		t.Fatalf("Expected %v, received %v", expected, canonicalPath)
	}
}

func TestSimplifyPath6(t *testing.T) {
	canonicalPath := simplifyPath("/home/foo/bar/room/../gate")
	expected := "/home/foo/bar/gate"

	if canonicalPath != expected {
		t.Fatalf("Expected %v, received %v", expected, canonicalPath)
	}
}
