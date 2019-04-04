package main

import (
	"fmt"
	"os"
	"sort"
)

var FindUniqueLocations = uniqueLocationsGrid

func main() {
	if len(os.Args) < 2 {
		panic("Input not provided")
	}
	input := os.Args[1]
	fmt.Println(uniqueLocationsGrid(input))
}

type position struct {
	x int
	y int
}

func (p *position) move(direction rune) {
	if direction == 'N' {
		p.y++
		return
	}
	if direction == 'E' {
		p.x++
		return
	}
	if direction == 'S' {
		p.y--
		return
	}
	if direction == 'W' {
		p.x--
		return
	}
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

type positionKey position

func uniqueLocationsSet(input string) int {
	var curr = position{0, 0}
	var set = map[positionKey]struct{}{}
	for _, direction := range input {
		curr.move(direction)
		set[curr.Key()] = struct{}{}
	}

	return len(set)
}

type btree struct {
	root *node
}

type node struct {
	key   int
	value *btree
	left  *node
	right *node
}

func uniqueLocationsTree(input string) int {
	var x, y int
	// var points = make([][2]int, 0, len(input))
	var tree *Node

	for _, direction := range input {
		move(direction, &x, &y)
		if tree == nil {
			tree = &Node{location: [2]int{x, y}}
		} else {
			tree.Insert([2]int{x, y}, 0)
		}
	}

	return tree.Len()
}

var maxcap int

func uniqueLocationsGrid(input string) int {
	var x, y, visited int
	maxcap = len(input)
	var world = make([][2]uint64, 0, maxcap)

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

type grid interface {
	exists(x, y int) bool
	visit(x, y int)
}

func newgrid(gridtype string) grid {
	switch gridtype {
	case "map":
		return make(mapgrid)
	case "array":
		return make(arraygrid, 0, maxcap)
	case "compositearray":
		return &compositearray{
			pp: make(arraygrid, 0, maxcap),
			np: make(arraygrid, 0, maxcap),
			pn: make(arraygrid, 0, maxcap),
			nn: make(arraygrid, 0, maxcap),
		}
	case "bitmap":
		return make(bitmap)
	}
	panic("unknown type")
}

type mapgrid map[int]struct{}

func (g mapgrid) exists(x, y int) bool {
	_, ok := g[x*maxcap+y]
	return ok
}

func (g mapgrid) visit(x, y int) {
	g[x*maxcap+y] = struct{}{}
}

type bitshift []uint64

func exists(world [][2]uint64, x, y int) bool {
	position := x*maxcap + y
	d := 0
	if position < 0 {
		position = position * -1
		d = 1
	}

	i := (position / 62)

	if len(world) < i+1 {
		return false
	}

	return world[i][d]&(1<<uint(position%62)) > 0
}

func visit(world [][2]uint64, x, y int) [][2]uint64 {
	position := x*maxcap + y
	d := 0
	if position < 0 {
		position = position * -1
		d = 1
	}

	i := position / 62

	for j := len(world); j < i+1; j++ {
		world = append(world, [2]uint64{})
	}

	world[i][d] |= (1 << uint(position%62))
	return world
}

type arraygrid [][]bool

func (g arraygrid) exists(x, y int) bool {
	if x >= len(g) {
		return false
	}
	if y >= len(g[x]) {
		return false
	}
	return g[x][y]
}

func (g arraygrid) visit(x, y int) {
	for i := len(g); i <= x; i++ {
		g = append(g, make([]bool, 0, maxcap))
	}
	for j := len(g[x]); j <= y; j++ {
		g[x] = append(g[x], false)
	}
	g[x][y] = true
}

type compositearray struct {
	pp, pn, np, nn arraygrid
}

func (g *compositearray) exists(x, y int) bool {
	if x >= 0 && y >= 0 {
		return g.pp.exists(x, y)
	}

	if x < 0 && y < 0 {
		return g.nn.exists(-x, -y)
	}

	if x < 0 && y >= 0 {
		return g.np.exists(-x, y)
	}

	return g.pn.exists(x, -y)
}

func (g *compositearray) visit(x, y int) {
	if x >= 0 && y >= 0 {
		g.pp.visit(x, y)
		return
	}

	if x < 0 && y < 0 {
		g.nn.visit(-x, -y)
		return
	}

	if x < 0 && y >= 0 {
		g.np.visit(-x, y)
		return
	}

	g.pn.visit(x, -y)
}

type bitmap map[int]uint64

func (b bitmap) exists(x, y int) bool {
	return b[x]&(1<<uint(y)) > 0
}

func (b bitmap) visit(x, y int) {
	b[x] |= (1 << uint(y))
}

type ByAxis struct {
	axis   int
	points [][2]int
}

func (a ByAxis) Len() int           { return len(a.points) }
func (a ByAxis) Swap(i, j int)      { a.points[i], a.points[j] = a.points[j], a.points[i] }
func (a ByAxis) Less(i, j int) bool { return a.points[i][a.axis] < a.points[j][a.axis] }

func kdtree(points [][2]int, depth int) *Node {
	if len(points) == 0 {
		return nil
	}

	k := len(points[0])
	axis := depth % k

	sort.Sort(ByAxis{axis, points})
	median := len(points) / 2

	return &Node{
		location: points[median],
		left:     kdtree(points[:median], depth+1),
		right:    kdtree(points[median+1:], depth+1),
	}
}

type Node struct {
	location    [2]int
	left, right *Node
}

func (n *Node) Len() int {
	if n == nil {
		return 0
	}

	var llen, rlen int
	if n.left != nil {
		llen = n.left.Len()
	}

	if n.right != nil {
		rlen = n.right.Len()
	}

	return 1 + llen + rlen
}

func (n *Node) Insert(new [2]int, depth int) {
	if n.location[0] == new[0] && n.location[1] == new[1] {
		return
	}

	k := 2
	axis := depth % k

	if new[axis] < n.location[axis] {
		if n.left == nil {
			n.left = &Node{location: new}
		} else {
			n.left.Insert(new, depth+1)
		}
		return
	}

	if n.right == nil {
		n.right = &Node{location: new}
	} else {
		n.right.Insert(new, depth+1)
	}
}

func (n *KDNode) moveNode(direction rune) {
	if direction == 'N' {
		n.Y++
		return
	}
	if direction == 'E' {
		n.X++
		return
	}
	if direction == 'S' {
		n.Y--
		return
	}
	if direction == 'W' {
		n.X--
		return
	}
}

// String for printing in tests
func (p position) String() string {
	return fmt.Sprintf("{%d, %d}", p.x, p.y)
}

func (p position) Key() positionKey {
	return positionKey(p)
}

type KDNode struct {
	X     int
	Y     int
	Left  *KDNode
	Right *KDNode
}

func (n *KDNode) Append(node *KDNode, x bool, len *int) {
	if node.X == n.X && node.Y == n.Y {
		return
	}

	if (x && node.X < n.X) || (!x && node.Y < n.Y) {
		if n.Left == nil {
			n.Left = &KDNode{X: node.X, Y: node.Y}
			*len++
			return
		}

		n.Left.Append(node, !x, len)
		return
	}

	if n.Right == nil {
		n.Right = &KDNode{X: node.X, Y: node.Y}
		*len++
		return
	}

	n.Right.Append(node, !x, len)
	return
}

func (n *KDNode) String() string {
	return fmt.Sprintf("Node{X: %d, Y: %d, Left: %v, Right: %v}", n.X, n.Y, n.Left, n.Right)
}

type KDTree struct {
	Root *KDNode
	Len  int
}

func (t *KDTree) Append(node *KDNode) {
	if t.Root == nil {
		t.Root = &KDNode{X: node.X, Y: node.Y}
		t.Len++
		return
	}

	t.Root.Append(node, true, &t.Len)
}

func (t *KDTree) String() string {
	return fmt.Sprintf("%s", t.Root)
}
