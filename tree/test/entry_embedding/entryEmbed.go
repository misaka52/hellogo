package main

import (
	"hellogo/tree/test"
)

type myTreeNode struct {
	// 响应与定义了名为Node的Node节点，不可重名
	*test.Node
}

// 扩展后续遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}

	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func main() {

}
