package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {

	os.Setenv("INPUT", `two1nine
  eightwothree
  abcone2threexyz
  xtwone3four
  4nineeightseven2
  zoneight234
  7pqrstsixteen`)

	var want int = 281
	got := RoundTwoProblem()

	if want != got {
		t.Fatalf(`RoundTwoProblem() - got: %q, want: %v`, got, want)
	}
}
