package main

import (
	"fmt"
	"hellogo/tree/test"
)

func main() {
	var root test.Node
	root = test.Node{Value: 3}
	root.Left = &test.Node{}
	root.Right = &test.Node{nil, nil, 5}
	root.Left.Left = new(test.Node)
	root.Left.Right = test.CreateTreeNode(2)

	nodes := []test.Node{
		{Value: 3},
		{nil, nil, 1},
		root,
	}
	fmt.Println(nodes)

	root2 := test.Node{}
	root2.Print()
	root2.SetValue(2)
	root.Print()

	var rootNil *test.Node
	rootNil.SetValue(2)

	root.Left.SetValue(2)
	root.TraverseMiddle()

	var c chan *test.Node
	c = root.TraverseWithChannel()

	maxValue := 0
	for node := range c {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}
	fmt.Println("Max node value is ", maxValue)

}
