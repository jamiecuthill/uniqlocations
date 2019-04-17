package main

import (
	"math/rand"
	"testing"
)

func BenchmarkUniqueLocations(b *testing.B) {
	b.StopTimer()
	var input []rune
	for n := 0; n < 8574; n++ {
		input = append(input, randomDirection())
	}
	inputstr := string(input)
	b.StartTimer()

	b.Run("Using Grid", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			uniqueLocationsGrid(inputstr)
		}
	})
}

func randomDirection() rune {
	switch rand.Intn(4) {
	case 0:
		return 'N'
	case 1:
		return 'E'
	case 2:
		return 'S'
	case 3:
		return 'W'
	}
	panic("Generating invalid direction")
}
