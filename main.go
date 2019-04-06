package main

const gridsize = 8

var maxcap int
var FindUniqueLocations = uniqueLocationsGrid

func uniqueLocationsGrid(input string) int {
	var x, y, visited int
	// maxcap = (len(input) / gridsize) + 1
	maxcap = 20
	var world = make([][][4]uint64, 0, maxcap)

	for _, direction := range input {
		move(direction, &x, &y)
		if exists(world, x, y) {
			continue
		}
		visited++
		world = visit(world, x, y)
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

func exists(world [][][4]uint64, x, y int) bool {
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
	if len(world) < xi+1 {
		return false
	}

	yi := y / gridsize
	if len(world[xi]) < yi+1 {
		return false
	}

	shift := uint((x%gridsize)*gridsize + (y % gridsize))

	return world[xi][yi][d]&(1<<shift) > 0
}

func visit(world [][][4]uint64, x, y int) [][][4]uint64 {
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

	shift := uint((x%gridsize)*gridsize + (y % gridsize))

	world[xi][yi][d] |= (1 << shift)
	return world
}

func main() {}
