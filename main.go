package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Input not provided")
	}
	input := os.Args[1]
	fmt.Println(uniqueLocations(input))
}

type position struct {
	x int
	y int
}

var north = position{0, 1}
var east = position{1, 0}
var south = position{0, -1}
var west = position{-1, 0}

func uniqueLocations(input string) int {
	var curr = position{0, 0}
	// var set = map[position]*struct{}{}
	var set = map[[2]int]*struct{}{}
	var marker = &struct{}{}
	for _, movement := range input {
		curr = move(curr, compassToPosition(movement))
		set[curr.Key()] = marker
	}
	// Is there such a thing as a max map size?
	return len(set)
}

func move(curr, direction position) position {
	// curr.x += direction.x
	// curr.y += direction.y
	return position{curr.x + direction.x, curr.y + direction.y}
}

func compassToPosition(in rune) position {
	switch in {
	case 'N':
		fallthrough
	case 'n':
		return north
	case 'E':
		fallthrough
	case 'e':
		return east
	case 'S':
		fallthrough
	case 's':
		return south
	case 'W':
		fallthrough
	case 'w':
		return west
	default:
		panic("Invalid input character")
	}
}

// String for printing in tests
func (p position) String() string {
	return fmt.Sprintf("{%d, %d}", p.x, p.y)
}

func (p position) Key() [2]int {
	return [2]int{p.x, p.y}
}
