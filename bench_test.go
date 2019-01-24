package main

import (
	"math/rand"
	"testing"
)

func BenchmarkUniqueLocations(b *testing.B) {
	b.StopTimer()
	var input []byte
	for n := 0; n < 5000; n++ {
		input = append(input, byte(randomDirection()))
	}
	b.StartTimer()

	b.Run("Find Unique Locations", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			uniqueLocations(string(input))
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
