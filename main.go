package main

const gridsize = 8

var maxcap = 20

var FindUniqueLocations = uniqueLocationsGrid

type worldgrid [][][4]uint64

func uniqueLocationsGrid(input string) int {
	var x, y, visited int
	var world = make(worldgrid, 0, maxcap)
	var ok bool

	for _, direction := range input {
		move(direction, &x, &y)
		if world, ok = exists(world, x, y); !ok {
			visited++
		}
	}

	return visited
}

func move(direction rune, x *int, y *int) {
	switch direction {
	case 'N':
		*y++
	case 'E':
		*x++
	case 'S':
		*y--
	case 'W':
		*x--
	}
}

func exists(world worldgrid, x, y int) (worldgrid, bool) {
	d := 0
	if x < 0 {
		d++
		x = -x
	}
	if y < 0 {
		d += 2
		y = -y
	}

	xi := x >> 3
	for i := len(world); i < xi+1; i++ {
		world = append(world, make([][4]uint64, 0, maxcap))
	}

	yi := y >> 3
	for i := len(world[xi]); i < yi+1; i++ {
		world[xi] = append(world[xi], [4]uint64{})
	}

	shift := uint((x&(gridsize-1))<<3 + (y & (gridsize - 1)))

	if world[xi][yi][d]&(1<<shift) > 0 {
		return world, true
	}
	world[xi][yi][d] |= (1 << shift)
	return world, false
}

func main() {}
