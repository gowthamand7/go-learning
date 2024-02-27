package hello

import "testing"

func TestSayHello(t *testing.T) {

	//can be created using a struct
	subtests := []struct {
		input    []string
		expected string
	}{
		{[]string{"Gowthaman"}, "Hello Gowthaman!"},
		{[]string{"Gowthaman", "Priya"}, "Hello Gowthaman, Priya!"},
		{[]string{}, "Hello World!"},
	}

	for _, subtest := range subtests {
		got := SayHello(subtest.input)
		if got != subtest.expected {
			t.Errorf("Expected %s, got %s", subtest.expected, got)
		}
	}
}
