package main

import "golangdemo/tree"

func main() {

	//var root tree.TreeNode
	//fmt.Println(root)
	//root = TreeNode{
	//	Value: 3,
	//	Right: &TreeNode{},
	//	Left:  &TreeNode{}}
	//fmt.Println(root)
	//root.Left = &TreeNode{5, nil, nil}
	//fmt.Println(root)
	//root.Right.Left = &TreeNode{}
	//root.Right.Right = CreateNode(12)
	//fmt.Println(root)
	//fmt.Println(root.Left)
	//fmt.Println(root.Right)
	//nodes := []TreeNode{
	//	{Value: 3},
	//	{Right: &TreeNode{}},
	//	{Value: 5, Left: &TreeNode{5, nil, nil}},
	//	{3, nil, new(TreeNode)}}
	//fmt.Println(nodes)

	var root tree.TreeNode
	root.Value = 3
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Left.Right = CreateNode(2)

	root.Traverse()

	//fmt.Println("root Value:")
	//root.PrintValue()
	//
	//root.Right.Left.SetValue(4)
	//fmt.Println("root.Right.Left  Value:")
	//root.Right.Left.PrintValue()
	//
	//root.Right.Left.SetValuePointer(4)
	//fmt.Println("root.Right.Left  Value:")
	//root.Right.Left.PrintValue()
	//
	////root.Left.Left.SetValue(12)
	//fmt.Println("------------------------")
	//var pNode *TreeNode //pNode是nil指针
	//fmt.Println(pNode)
	////fmt.Println(*pNode)
	//pNode.SetValuePointer(200)
	//pNode = &root
	//pNode.SetValuePointer(300)
	//pNode.PrintValue()
	//
	//fmt.Println("------------------------")
	//var pRoot TreeNode //pRoot是treeNode的空值
	////fmt.Println(&pRoot)
	//fmt.Println(pRoot)
	//pRoot.SetValue(200)
	//pRoot = root
	//pRoot.SetValue(300)
	//pRoot.PrintValue()

}
