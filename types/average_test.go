package types

import (
	"os"
	"testing"
)

func TestAverage(t *testing.T) {
	// Mock user input
	inputs := []string{"10.00", "20.00", "30.00", ""}

	// Mock user input by redirecting Stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	os.Stdin = r
	go func() {
		defer w.Close()
		for _, input := range inputs {
			w.WriteString(input + "\n")
		}
	}()

	// Call the function and capture the output
	avg := Average()

	// Define a tolerance for floating-point comparison
	tolerance := 0.0001

	// Expected average for the given input is 20.0
	expected := 20.0

	// Check if the average matches the expected value within the tolerance
	if avg < expected-tolerance || avg > expected+tolerance {
		t.Errorf("Average was incorrect, got: %f, want: %f", avg, expected)
	}
}
