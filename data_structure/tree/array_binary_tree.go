/*
顺序二叉树：顺序二叉树是一种树形结构，通常只考虑完全二叉树，每个节点最多只有两个子节点，不过不用指针，使用数组来完成。
主要思想：
	1. 创建一个数组存储二叉树。
	2. 满足第n个元素的左子节点为2*n+1，右子节点为2*n+2，父节点为(n-1)/2。
	3. 前序：父节点=>左子节点=>右子节点。
	4. 中序：左子节点=>父节点=>右子节点。
	5. 后序：左子节点=>右子节点=>父节点。
	6. 完成前、中、后序遍历。
*/
package tree

import "fmt"

type ArrayBinaryTree struct {
	binaryTree []int
}

// 初始化顺序二叉树
func (arrayBinaryTree *ArrayBinaryTree) Init(tree []int)  {
	arrayBinaryTree.binaryTree = tree
}

// 前序遍历
func (arrayBinaryTree *ArrayBinaryTree) PreorderTraversal(index int)  {
	if arrayBinaryTree.binaryTree == nil || len(arrayBinaryTree.binaryTree)==0{
		fmt.Println("This tree is empty!")
		return
	}
	fmt.Print(arrayBinaryTree.binaryTree[index],"=>")
	// 向左递归
	if len(arrayBinaryTree.binaryTree) > 2*index+1{
		arrayBinaryTree.PreorderTraversal(2*index+1)
	}
	// 向右递归
	if len(arrayBinaryTree.binaryTree) > 2*index+2{
		arrayBinaryTree.PreorderTraversal(2*index+2)
	}
}

// 前序遍历
func (arrayBinaryTree *ArrayBinaryTree) InorderTraversal(index int)  {
	if arrayBinaryTree.binaryTree == nil || len(arrayBinaryTree.binaryTree)==0{
		fmt.Println("This tree is empty!")
		return
	}
	// 向左递归
	if len(arrayBinaryTree.binaryTree) > 2*index+1{
		arrayBinaryTree.InorderTraversal(2*index+1)
	}
	fmt.Print(arrayBinaryTree.binaryTree[index],"=>")
	// 向右递归
	if len(arrayBinaryTree.binaryTree) > 2*index+2{
		arrayBinaryTree.InorderTraversal(2*index+2)
	}
}

// 前序遍历
func (arrayBinaryTree *ArrayBinaryTree) PostOrderTraversal(index int)  {
	if arrayBinaryTree.binaryTree == nil || len(arrayBinaryTree.binaryTree)==0{
		fmt.Println("This tree is empty!")
		return
	}
	// 向左递归
	if len(arrayBinaryTree.binaryTree) > 2*index+1{
		arrayBinaryTree.PostOrderTraversal(2*index+1)
	}
	// 向右递归
	if len(arrayBinaryTree.binaryTree) > 2*index+2{
		arrayBinaryTree.PostOrderTraversal(2*index+2)
	}
	fmt.Print(arrayBinaryTree.binaryTree[index],"=>")
}