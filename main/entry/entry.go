package main

import (
	"golangdemo/tree"
)

type myTreeNode struct {
	node *tree.TreeNode
}

//后续遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	//指针不能作为接收者, 需要需要定义变量来接送地址
	//myTreeNode{myNode.node.Left}.postOrder();

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.PrintValue()
}

func main() {

	//var root tree.tree.TreeNode
	//fmt.Println(root)
	//root = tree.TreeNode{
	//	Value: 3,
	//	Right: &tree.TreeNode{},
	//	Left:  &tree.TreeNode{}}
	//fmt.Println(root)
	//root.Left = &tree.TreeNode{5, nil, nil}
	//fmt.Println(root)
	//root.Right.Left = &tree.TreeNode{}
	//root.Right.Right = CreateNode(12)
	//fmt.Println(root)
	//fmt.Println(root.Left)
	//fmt.Println(root.Right)
	//nodes := []tree.TreeNode{
	//	{Value: 3},
	//	{Right: &tree.TreeNode{}},
	//	{Value: 5, Left: &tree.TreeNode{5, nil, nil}},
	//	{3, nil, new(tree.TreeNode)}}
	//fmt.Println(nodes)

	var root tree.TreeNode
	root.Value = 3
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	//root.Traverse()
	pRoot := myTreeNode{&root}
	pRoot.postOrder()

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
	//var pNode *tree.TreeNode //pNode是nil指针
	//fmt.Println(pNode)
	////fmt.Println(*pNode)
	//pNode.SetValuePointer(200)
	//pNode = &root
	//pNode.SetValuePointer(300)
	//pNode.PrintValue()
	//
	//fmt.Println("------------------------")
	//var pRoot tree.TreeNode //pRoot是tree.TreeNode的空值
	////fmt.Println(&pRoot)
	//fmt.Println(pRoot)
	//pRoot.SetValue(200)
	//pRoot = root
	//pRoot.SetValue(300)
	//pRoot.PrintValue()

}
