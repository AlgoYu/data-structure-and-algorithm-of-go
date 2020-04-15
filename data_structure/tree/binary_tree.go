/*
二叉树：二叉树是一种树形结构，每个节点最多只有两个子节点。
主要思想：
	1. 创建左右两个指针节点指针。
	2. 使用左右节点指针完成前、中、后序遍历。
	3. 前序：父节点=>左子节点=>右子节点。
	4. 中序：左子节点=>父节点=>右子节点。
	5. 后序：左子节点=>右子节点=>父节点。
*/
package tree

import "fmt"

// 二叉树节点
type BinaryTreeNode struct {
	Id int
	Data int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

// 前序遍历
func (binaryTreeNode *BinaryTreeNode) PreorderTraversal()  {
	fmt.Println(binaryTreeNode.Id)
	if binaryTreeNode.Left != nil{
		binaryTreeNode.Left.PreorderTraversal()
	}
	if binaryTreeNode.Right != nil{
		binaryTreeNode.Right.PreorderTraversal()
	}
}

// 中序遍历
func (binaryTreeNode *BinaryTreeNode) InorderTraversal()  {
	if binaryTreeNode.Left != nil{
		binaryTreeNode.Left.InorderTraversal()
	}
	fmt.Println(binaryTreeNode.Id)
	if binaryTreeNode.Right != nil{
		binaryTreeNode.Right.InorderTraversal()
	}
}

// 后序遍历
func (binaryTreeNode *BinaryTreeNode) PostorderTraversal()  {
	if binaryTreeNode.Left != nil{
		binaryTreeNode.Left.PostorderTraversal()
	}
	if binaryTreeNode.Right != nil{
		binaryTreeNode.Right.PostorderTraversal()
	}
	fmt.Println(binaryTreeNode.Id)
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

// 前序遍历
func (binaryTree *BinaryTree) PreorderTraversal()  {
	if binaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		binaryTree.Root.PreorderTraversal()
	}
}

// 中序遍历
func (binaryTree *BinaryTree) InorderTraversal()  {
	if binaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		binaryTree.Root.InorderTraversal()
	}
}

// 后序遍历
func (binaryTree *BinaryTree) PostorderTraversal()  {
	if binaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		binaryTree.Root.PostorderTraversal()
	}
}