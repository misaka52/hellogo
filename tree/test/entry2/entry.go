package main

import (
	"fmt"
	"hellogo/tree/test"
)

type myTreeNode struct {
	node *test.Node
}

// 扩展后续遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func (myNode *myTreeNode) TraverseMiddle() {
	fmt.Println("this method is shadowed")
}

func main() {
	/*
				3
		       / \
		      1   5
		       \  /
		       2 4
		中序：1 2 3 4 5
	*/
	root := test.Node{Value: 3}
	root.Left = &test.Node{Value: 1}
	root.Right = &test.Node{Value: 5}
	root.Right.Left = new(test.Node)
	root.Left.Right = test.CreateTreeNode(2)
	root.Right.Left.SetValue(4)

	root.TraverseMiddle()

	fmt.Print("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	myRoot.TraverseMiddle()
	fmt.Println()
}
