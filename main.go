package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

var FindUniqueLocations = uniqueLocationsGrid

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Input not provided")
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

func uniqueLocationsGrid(input string) int {
	var x, y, visited int
	var world = make(mapgrid)
	for _, direction := range input {
		move(direction, &x, &y)
		if world.exists(x, y) {
			continue
		}
		visited++
		world.visit(x, y)
	}
	return visited
}

type mapgrid map[int]map[int]struct{}

func (g mapgrid) exists(x, y int) bool {
	if g == nil {
		g = make(map[int]map[int]struct{})
	}
	if xx, okx := g[x]; okx {
		if xx == nil {
			g[x] = make(map[int]struct{})
		}
		_, oky := g[x][y]
		return oky
	}
	return false
}

func (g mapgrid) visit(x, y int) {
	if g == nil {
		g = make(map[int]map[int]struct{})
	}
	if _, ok := g[x]; !ok {
		g[x] = make(map[int]struct{})
	}
	g[x][y] = struct{}{}
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
