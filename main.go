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

var north = position{0, 1}
var east = position{1, 0}
var south = position{0, -1}
var west = position{-1, 0}

type positionKey position

func uniqueLocationsSet(input string) int {
	var curr = position{0, 0}
	var set = map[positionKey]struct{}{}
	for _, movement := range input {
		curr = move(curr, compassToPosition(movement))
		set[curr.Key()] = struct{}{}
	}

	return len(set)
}

func uniqueLocationsTree(input string) int {
	curr := &KDNode{X: 0, Y: 0}
	tree := &KDTree{}
	for _, direction := range input {
		curr = moveNode(*curr, direction)
		// TODO If we know that we only ever change direction by a single value
		// it might be possible to traverse up the tree by allowing current to point to the node
		// as it is placed in the tree.
		tree.Append(curr)
	}
	return tree.Len()
}

func moveNode(curr KDNode, direction rune) *KDNode {
	switch direction {
	case 'N':
		fallthrough
	case 'n':
		return &KDNode{X: curr.X, Y: curr.Y + 1}
	case 'E':
		fallthrough
	case 'e':
		return &KDNode{X: curr.X + 1, Y: curr.Y}
	case 'S':
		fallthrough
	case 's':
		return &KDNode{X: curr.X, Y: curr.Y - 1}
	case 'W':
		fallthrough
	case 'w':
		return &KDNode{X: curr.X - 1, Y: curr.Y}
	default:
		panic("Invalid direction character")
	}
}

func move(curr, direction position) position {
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

func (p position) Key() positionKey {
	return positionKey(p)
}

type KDNode struct {
	X     int
	Y     int
	Left  *KDNode
	Right *KDNode

	Parent *KDNode
}

func (n *KDNode) Len() int {
	var left, right int
	if n.Left != nil {
		left = n.Left.Len()
	}
	if n.Right != nil {
		right = n.Right.Len()
	}

	return 1 + left + right
}

func (n *KDNode) Append(node *KDNode, depth int) int {
	if node.X == n.X && node.Y == n.Y {
		log.Printf("Found duplicate node %s", n)
		node = n
		return depth
	}

	var nodeValue, thisValue int
	if depth%2 == 0 {
		nodeValue = node.X
		thisValue = n.X

	} else {
		nodeValue = node.Y
		thisValue = n.Y
	}
	if nodeValue < thisValue {
		if n.Left == nil {
			n.Left = node
			node.Parent = n
			return depth
		}

		return n.Left.Append(node, depth+1)
	}

	if n.Right == nil {
		n.Right = node
		node.Parent = n
		return depth
	}

	return n.Right.Append(node, depth+1)
}

func (n *KDNode) String() string {
	return fmt.Sprintf("Node{X: %d, Y: %d, Left: %v, Right: %v}", n.X, n.Y, n.Left, n.Right)
}

type KDTree struct {
	Root  *KDNode
	next  *KDNode
	depth int
}

func (t *KDTree) Append(node *KDNode) {
	if t.Root == nil {
		log.Printf("Root is now %s", node)
		t.Root = node
		t.next = node
		return
	}

	insertedDepth := t.next.Append(node, t.depth)
	if node.Parent != nil && node.Parent.Parent != nil {
		log.Printf("Using parent node for next append %s at depth %d", node.Parent.Parent, insertedDepth-1)
		t.next = node.Parent.Parent
		t.depth = insertedDepth - 1
		return
	}

	log.Printf("Back at root node %s", t.Root)
	t.next = t.Root
	t.depth = 0
}

func (t *KDTree) Len() int {
	if t.Root == nil {
		return 0
	}

	return t.Root.Len()
}

func (t *KDTree) String() string {
	return fmt.Sprintf("%s", t.Root)
}
