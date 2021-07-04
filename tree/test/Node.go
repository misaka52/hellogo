package test

import (
	"fmt"
)

type Node struct {
	Left, Right *Node
	Value       int
}

// 工厂方法
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

// 私有方法
func createTreeNode() *Node {
	return &Node{Value: 0}
}

// 方法接收者，通过struct.Print()调用，可以接收nil。值传递
func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

// 指针传递
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("nil TreeNode, ignore")
		return
	}
	node.Value = value
}

// 中序遍历
func (node *Node) TraverseMiddle() {
	fmt.Print("In-order traversal:")
	node.TraverseMiddleFunc(func(n *Node) {
		fmt.Print(n.Value, " ")
	})
	fmt.Println()
}

func (node *Node) TraverseMiddleFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseMiddleFunc(f)
	f(node)
	node.Right.TraverseMiddleFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseMiddleFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
