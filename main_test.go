package main

import (
	"fmt"
	"testing"
)

var movementTests = []struct {
	curr      position
	direction position
	expected  position
}{
	{position{0, 0}, position{1, 0}, position{1, 0}},
	{position{0, 0}, position{0, 1}, position{0, 1}},
	{position{1, 1}, position{1, 0}, position{2, 1}},
	{position{1, 1}, position{0, 1}, position{1, 2}},
	{position{1, 1}, position{-1, -1}, position{0, 0}},
}

func TestMove(t *testing.T) {
	for _, test := range movementTests {
		t.Run(fmt.Sprintf("%s + %s", test.curr, test.direction), func(t *testing.T) {
			actual := move(test.curr, test.direction)

			if actual != test.expected {
				t.Errorf("Expected position %s but was %s", test.expected, actual)
			}
		})
	}
}

var compassTests = []struct {
	in       rune
	expected position
}{
	{'N', position{0, 1}},
	{'E', position{1, 0}},
	{'S', position{0, -1}},
	{'W', position{-1, 0}},

	{'n', position{0, 1}},
	{'e', position{1, 0}},
	{'s', position{0, -1}},
	{'w', position{-1, 0}},
}

func TestCompassToCoordinates(t *testing.T) {
	for _, test := range compassTests {
		t.Run(string(test.in), func(t *testing.T) {
			actual := compassToPosition(test.in)

			if actual != test.expected {
				t.Errorf("Expected position %s but was %s", test.expected, actual)
			}
		})
	}
}
