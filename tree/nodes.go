package tree

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) PrintValue() {
	fmt.Println(node.Value)
}

//值传递  不会更改值
func (node TreeNode) SetValue(value int) {
	if &node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	fmt.Println("pointer : ", &node)
	node.Value = value
}

//指针传递  会更改值
func (node *TreeNode) SetValuePointer(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.PrintValue()
	node.Right.Traverse()
}

func (node *TreeNode) TraverseUseFunc() {
	node.TraverseFuncf(func(n *TreeNode) {
		n.PrintValue()
	})
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node *TreeNode) TraverseFuncf(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFuncf(f)
	f(node)
	node.Right.TraverseFuncf(f)
}
