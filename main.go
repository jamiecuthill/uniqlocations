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
	fmt.Println(uniqueLocationsTree(input))
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

func uniqueLocationsTree(input string) int {
	curr := &KDNode{X: 0, Y: 0}
	tree := &KDTree{}
	for _, direction := range input {
		curr.moveNode(direction)
		tree.Append(curr)
	}
	return tree.Len
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

func move(curr, direction position) position {
	return position{curr.x + direction.x, curr.y + direction.y}
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
