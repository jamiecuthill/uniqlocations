package main

import (
	"reflect"
	"testing"
)

var movementTests = []struct {
	curr      position
	direction position
	expected  position
}{
	{position{0, 0}, position{1, 0}, position{1, 0}},
	{position{0, 0}, position{0, 1}, position{0, 1}},
	{position{1, 1}, position{1, 0}, position{2, 1}},
	{position{1, 1}, position{0, 1}, position{1, 2}},
	{position{1, 1}, position{-1, -1}, position{0, 0}},
}

var compassTests = []struct {
	in       rune
	expected position
}{
	{'N', position{0, 1}},
	{'E', position{1, 0}},
	{'S', position{0, -1}},
	{'W', position{-1, 0}},

	{'n', position{0, 1}},
	{'e', position{1, 0}},
	{'s', position{0, -1}},
	{'w', position{-1, 0}},
}

// func TestCompassToCoordinates(t *testing.T) {
// 	for _, test := range compassTests {
// 		t.Run(string(test.in), func(t *testing.T) {
// 			actual := compassToPosition(test.in)

// 			if actual != test.expected {
// 				t.Errorf("Expected position %s but was %s", test.expected, actual)
// 			}
// 		})
// 	}
// }

func TestEmptyTree(t *testing.T) {
	tree := &KDTree{}
	node := &KDNode{X: 0, Y: 0}

	tree.Append(node)

	if !reflect.DeepEqual(tree.Root, node) {
		t.Fatalf("Root was not set to node: %s but was %s", node, tree.Root)
	}
}

func TestTreeAppendsToRoot(t *testing.T) {
	tree := &KDTree{}
	tree.Append(&KDNode{X: 0, Y: 0})
	node := &KDNode{X: 1, Y: 0}

	tree.Append(node)

	if !reflect.DeepEqual(tree.Root.Right, node) {
		t.Fatalf("Right of root was not set to node: %s but was %s", node, tree.Root.Right)
	}
}

func TestNodeAppendsLowerToLeftOfX(t *testing.T) {
	node := &KDNode{X: 0, Y: 0}
	leaf := &KDNode{X: -1, Y: 0}

	len := 0
	node.Append(leaf, true, &len)

	if !reflect.DeepEqual(node.Left, leaf) {
		t.Fatalf("Left was not set to leaf: %s but was %s", leaf, node.Left)
	}
}

func TestNodeAppendsHigherToRightOfX(t *testing.T) {
	node := &KDNode{X: 0, Y: 0}
	leaf := &KDNode{X: 1, Y: 0}

	len := 0
	node.Append(leaf, true, &len)

	if !reflect.DeepEqual(node.Right, leaf) {
		t.Fatalf("Right was not set to leaf: %s but was %s", leaf, node.Right)
	}
}

func TestNodeAppendSameNodeIsIgnored(t *testing.T) {
	node := &KDNode{X: 0, Y: 0}
	duplicate := &KDNode{X: 0, Y: 0}

	len := 0
	node.Append(duplicate, true, &len)

	if node.Left != nil {
		t.Fatalf("Node left value should be nil but was %s", node.Left)
	}

	if node.Right != nil {
		t.Fatalf("Node right value should be nil but was %s", node.Right)
	}
}

func TestNodeAppendsLowerToLeftOfY(t *testing.T) {
	node := &KDNode{X: 0, Y: 0}
	leaf := &KDNode{X: 0, Y: -1}

	len := 0
	node.Append(leaf, false, &len)

	if !reflect.DeepEqual(node.Left, leaf) {
		t.Fatalf("Left was not set to leaf: %s but was %s", leaf, node.Left)
	}
}

func TestNodeAppendsHigherToRightOfY(t *testing.T) {
	node := &KDNode{X: 0, Y: 0}
	leaf := &KDNode{X: 0, Y: 1}

	len := 0
	node.Append(leaf, false, &len)

	if !reflect.DeepEqual(node.Right, leaf) {
		t.Fatalf("Right was not set to leaf: %s but was %s", leaf, node.Right)
	}
}

func TestNodeXAppendedToChildOfLeft(t *testing.T) {
	node := &KDNode{X: 0, Y: 0, Left: &KDNode{X: -1, Y: 0}}
	leaf := &KDNode{X: -1, Y: -1}

	len := 0
	node.Append(leaf, true, &len)

	if !reflect.DeepEqual(node.Left.Left, leaf) {
		t.Fatalf("Child of left node was not set to leaf: %s", node.Left)
	}
}

func TestNodeXAppendedToChildOfRight(t *testing.T) {
	node := &KDNode{X: 0, Y: 0, Right: &KDNode{X: 1, Y: 0}}
	leaf := &KDNode{X: 1, Y: 1}

	len := 0
	node.Append(leaf, true, &len)

	if !reflect.DeepEqual(node.Right.Right, leaf) {
		t.Fatalf("Child of right node was not set to leaf: %s", node.Right)
	}
}

func TestNodeYAppendedToChildOfLeft(t *testing.T) {
	node := &KDNode{X: 0, Y: 0, Left: &KDNode{X: 0, Y: -1}}
	leaf := &KDNode{X: -1, Y: -1}

	len := 0
	node.Append(leaf, false, &len)

	if !reflect.DeepEqual(node.Left.Left, leaf) {
		t.Fatalf("Child of left node was not set to leaf: %s", node.Left)
	}
}

func TestNodeYAppendedToChildOfRight(t *testing.T) {
	node := &KDNode{X: 0, Y: 0, Right: &KDNode{X: 0, Y: 1}}
	leaf := &KDNode{X: 1, Y: 1}

	len := 0
	node.Append(leaf, false, &len)

	if !reflect.DeepEqual(node.Right.Right, leaf) {
		t.Fatalf("Child of right node was not set to leaf: %s", node.Right)
	}
}

func TestEmptyTreeLength(t *testing.T) {
	tree := &KDTree{}

	if tree.Len != 0 {
		t.Fatalf("Tree should be empty but length was %d", tree.Len)
	}
}

func TestTreeLength(t *testing.T) {
	tree := &KDTree{}
	tree.Append(&KDNode{})

	if tree.Len != 1 {
		t.Fatalf("Tree should have length but length was %d", tree.Len)
	}
}

// func TestNodeLengthWithNoChildren(t *testing.T) {
// 	node := &KDNode{}

// 	i := 0
// 	node.Len(&i)

// 	if i != 1 {
// 		t.Fatalf("Node should have length 1 but length was %d", i)
// 	}
// }

// func TestNodeLengthWithLeftChildren(t *testing.T) {
// 	node := &KDNode{Left: &KDNode{}}

// 	i := 0
// 	node.Len(&i)
// 	if i != 2 {
// 		t.Fatalf("Node should have length 2 but length was %d", i)
// 	}
// }

// func TestNodeLengthWithRightChildren(t *testing.T) {
// 	node := &KDNode{Right: &KDNode{}}

// 	i := 0
// 	node.Len(&i)
// 	if i != 2 {
// 		t.Fatalf("Node should have length 2 but length was %d", i)
// 	}
// }

// func TestNodeLengthWithChildren(t *testing.T) {
// 	node := &KDNode{Left: &KDNode{}, Right: &KDNode{}}

// 	i := 0
// 	node.Len(&i)
// 	if i != 3 {
// 		t.Fatalf("Node should have length 3 but length was %d", i)
// 	}
// }

func TestUniqueLocationsTree(t *testing.T) {
	res := uniqueLocationsTree("NNEESSWW")
	if res != 8 {
		t.Fatalf("Expected length is 8 but was %d", res)
	}
}

func TestUniqueLocationsTree2(t *testing.T) {
	res := uniqueLocationsTree("NESSWWNNES")
	if res != 9 {
		t.Fatalf("Expected length is 9 but was %d", res)
	}
}
