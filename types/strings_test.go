package types

import (
	"testing"
)

func TestStings(t *testing.T) {
	var input, search, replace string

	input = "Hello World!"
	search = "World"
	replace = "Gowthaman"

	got := Strings(input, search, replace)

	expected := MyString{
		input,
		"Hello Gowthaman!",
		16,
	}

	if got != expected {
		t.Errorf("Strings was incorrect, got: %s, want: %s", got.output, expected.output)
	}

}
