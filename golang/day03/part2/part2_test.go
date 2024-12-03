package part2

import (
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input_small.txt")
	if err != nil {
		t.Errorf("Error readin file: %v", err)
	}

	want := 48
	got := Solve(string(input))

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}