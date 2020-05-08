package main

import (
	"fmt"
	"golangdemo/tree"
)

func main() {

	var root tree.TreeNode
	root.Value = 3
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Right.Left.Value = 4
	root.Left.Right = tree.CreateNode(2)

	root.TraverseUseFunc()

	nodeCount := 0
	root.TraverseFuncf(func(node *tree.TreeNode) {
		nodeCount++
	})

	fmt.Println("nodeCount is : ", nodeCount)
}
