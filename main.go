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
	if direction == 'N' {
		*y++
		return
	}
	if direction == 'E' {
		*x++
		return
	}
	if direction == 'S' {
		*y--
		return
	}
	if direction == 'W' {
		*x--
		return
	}
}

func exists(world worldgrid, x, y int) (worldgrid, bool) {
	d := 0
	if x < 0 {
		d++
		x *= -1
	}
	if y < 0 {
		d += 2
		y *= -1
	}

	xi := x / gridsize
	for i := len(world); i < xi+1; i++ {
		world = append(world, make([][4]uint64, 0, maxcap))
	}

	yi := y / gridsize
	for i := len(world[xi]); i < yi+1; i++ {
		world[xi] = append(world[xi], [4]uint64{})
	}

	shift := uint((x&(gridsize-1))*gridsize + (y & (gridsize - 1)))

	if world[xi][yi][d]&(1<<shift) > 0 {
		return world, true
	}
	world[xi][yi][d] |= (1 << shift)
	return world, false
}

func main() {}
