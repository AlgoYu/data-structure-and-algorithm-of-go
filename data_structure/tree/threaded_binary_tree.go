/*
线索化二叉树：线索化二叉树是一种树形结构，特点是按某一种遍历顺序把空值的左右指针指向当前节点的前置节点和后置节点。
主要思想：
	1. 创建左右两个指针节点指针。
	2. 如果左节点为空指向当前节点的前置节点，如果有节点为空指向当前节点的后置节点。
	3. 前序：父节点=>左子节点=>右子节点。
	4. 中序：左子节点=>父节点=>右子节点。
	5. 后序：左子节点=>右子节点=>父节点。
*/
package tree

import "fmt"

// 二叉树节点
type ThreadedBinaryTreeNode struct {
	Id int
	Data int
	Left *ThreadedBinaryTreeNode
	Right *ThreadedBinaryTreeNode
	IsLeftThreaded bool
	IsRightThreaded bool
}

// 删除节点
func (threadedBinaryTreeNode *ThreadedBinaryTreeNode) DeleteNode(id int)  {
	if threadedBinaryTreeNode.Left!=nil && threadedBinaryTreeNode.Left.Id == id{
		threadedBinaryTreeNode.Left = nil
		return
	}
	if threadedBinaryTreeNode.Right!=nil && threadedBinaryTreeNode.Right.Id == id{
		threadedBinaryTreeNode.Right = nil
		return
	}
	if threadedBinaryTreeNode.Left!=nil{
		threadedBinaryTreeNode.Left.DeleteNode(id)
	}
	if threadedBinaryTreeNode.Right!=nil{
		threadedBinaryTreeNode.Right.DeleteNode(id)
	}
}

// 前序遍历
func (threadedBinaryTreeNode *ThreadedBinaryTreeNode) PreorderTraversal()  {
	fmt.Print(threadedBinaryTreeNode.Id,"=>")
	if threadedBinaryTreeNode.Left != nil{
		threadedBinaryTreeNode.Left.PreorderTraversal()
	}
	if threadedBinaryTreeNode.Right != nil{
		threadedBinaryTreeNode.Right.PreorderTraversal()
	}
}

// 中序遍历
func (threadedBinaryTreeNode *ThreadedBinaryTreeNode) InorderTraversal()  {
	if threadedBinaryTreeNode.Left != nil{
		threadedBinaryTreeNode.Left.InorderTraversal()
	}
	fmt.Print(threadedBinaryTreeNode.Id,"=>")
	if threadedBinaryTreeNode.Right != nil{
		threadedBinaryTreeNode.Right.InorderTraversal()
	}
}

// 后序遍历
func (threadedBinaryTreeNode *ThreadedBinaryTreeNode) PostorderTraversal()  {
	if threadedBinaryTreeNode.Left != nil{
		threadedBinaryTreeNode.Left.PostorderTraversal()
	}
	if threadedBinaryTreeNode.Right != nil{
		threadedBinaryTreeNode.Right.PostorderTraversal()
	}
	fmt.Print(threadedBinaryTreeNode.Id,"=>")
}

// 前序搜索
func (threadedBinaryTreeNode *ThreadedBinaryTreeNode) PreorderSearch(id int) *ThreadedBinaryTreeNode {
	if threadedBinaryTreeNode.Id == id{
		return threadedBinaryTreeNode
	}
	var node *ThreadedBinaryTreeNode
	if threadedBinaryTreeNode.Left != nil{
		node = threadedBinaryTreeNode.Left.PreorderSearch(id)
	}
	if node!=nil{
		return node
	}
	if threadedBinaryTreeNode.Right != nil{
		node = threadedBinaryTreeNode.Right.PreorderSearch(id)
	}
	return node
}

// 中序搜索
func (threadedBinaryTreeNode *ThreadedBinaryTreeNode) InorderSearch(id int) *ThreadedBinaryTreeNode {
	var node *ThreadedBinaryTreeNode
	if threadedBinaryTreeNode.Left != nil{
		node = threadedBinaryTreeNode.Left.InorderSearch(id)
	}
	if node!=nil{
		return node
	}
	if threadedBinaryTreeNode.Id == id{
		return threadedBinaryTreeNode
	}
	if threadedBinaryTreeNode.Right != nil{
		node = threadedBinaryTreeNode.Right.InorderSearch(id)
	}
	return node
}

// 后序搜索
func (threadedBinaryTreeNode *ThreadedBinaryTreeNode) PostorderSearch(id int) *ThreadedBinaryTreeNode {
	var node *ThreadedBinaryTreeNode
	if threadedBinaryTreeNode.Left != nil{
		node = threadedBinaryTreeNode.Left.PostorderSearch(id)
	}
	if node!=nil{
		return node
	}
	if threadedBinaryTreeNode.Right != nil{
		node = threadedBinaryTreeNode.Right.PostorderSearch(id)
	}
	if node!=nil{
		return node
	}
	if threadedBinaryTreeNode.Id == id{
		return threadedBinaryTreeNode
	}
	return node
}

type ThreadedBinaryTree struct {
	Root *ThreadedBinaryTreeNode
	Pre *ThreadedBinaryTreeNode
}

// 前序线索化
func (threadedBinaryTree *ThreadedBinaryTree)PreorderThreaded(node *ThreadedBinaryTreeNode)  {
	if node==nil{
		return
	}
	if node.Left == nil{
		node.Left = threadedBinaryTree.Pre
		node.IsLeftThreaded = true
	}
	if threadedBinaryTree.Pre!=nil && threadedBinaryTree.Pre.Right == nil{
		threadedBinaryTree.Pre.Right = node
		threadedBinaryTree.Pre.IsRightThreaded = true
	}
	threadedBinaryTree.Pre = node
	if !node.IsLeftThreaded{
		threadedBinaryTree.PreorderThreaded(node.Left)
	}
	if !node.IsRightThreaded{
		threadedBinaryTree.PreorderThreaded(node.Right)
	}
}

// 中序线索化
func (threadedBinaryTree *ThreadedBinaryTree)InorderThreaded(node *ThreadedBinaryTreeNode)  {
	if node==nil{
		return
	}
	if !node.IsLeftThreaded{
		threadedBinaryTree.InorderThreaded(node.Left)
	}
	if node.Left == nil{
		node.Left = threadedBinaryTree.Pre
		node.IsLeftThreaded = true
	}
	if threadedBinaryTree.Pre!=nil && threadedBinaryTree.Pre.Right == nil{
		threadedBinaryTree.Pre.Right = node
		threadedBinaryTree.Pre.IsRightThreaded = true
	}
	threadedBinaryTree.Pre = node
	if !node.IsRightThreaded{
		threadedBinaryTree.InorderThreaded(node.Right)
	}
}

// 后序线索化
func (threadedBinaryTree *ThreadedBinaryTree)PostorderThreaded(node *ThreadedBinaryTreeNode)  {
	if node==nil{
		return
	}
	if !node.IsLeftThreaded{
		threadedBinaryTree.PostorderThreaded(node.Left)
	}
	if !node.IsRightThreaded{
		threadedBinaryTree.PostorderThreaded(node.Right)
	}
	if node.Left == nil{
		node.Left = threadedBinaryTree.Pre
		node.IsLeftThreaded = true
	}
	if threadedBinaryTree.Pre!=nil && threadedBinaryTree.Pre.Right == nil{
		threadedBinaryTree.Pre.Right = node
		threadedBinaryTree.Pre.IsRightThreaded = true
	}
	threadedBinaryTree.Pre = node
}

// 删除节点
func (threadedBinaryTree *ThreadedBinaryTree) DeleteNode(id int)  {
	if threadedBinaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else if threadedBinaryTree.Root.Id == id{
		threadedBinaryTree.Root = nil
	}else{
		threadedBinaryTree.Root.DeleteNode(id)
	}
}

// 前序遍历
func (threadedBinaryTree *ThreadedBinaryTree) PreorderTraversal()  {
	if threadedBinaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		threadedBinaryTree.Root.PreorderTraversal()
	}
}

// 中序遍历
func (threadedBinaryTree *ThreadedBinaryTree) InorderTraversal()  {
	if threadedBinaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		threadedBinaryTree.Root.InorderTraversal()
	}
}

// 后序遍历
func (threadedBinaryTree *ThreadedBinaryTree) PostorderTraversal()  {
	if threadedBinaryTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		threadedBinaryTree.Root.PostorderTraversal()
	}
}

// 前序搜索
func (threadedBinaryTree *ThreadedBinaryTree) PreorderSearch(id int) *ThreadedBinaryTreeNode {
	if threadedBinaryTree.Root == nil{
		panic("This tree is empty!")
	}else{
		return threadedBinaryTree.Root.PreorderSearch(id)
	}
}

// 中序搜索
func (threadedBinaryTree *ThreadedBinaryTree) InorderSearch(id int) *ThreadedBinaryTreeNode {
	if threadedBinaryTree.Root == nil{
		panic("This tree is empty!")
	}else{
		return threadedBinaryTree.Root.InorderSearch(id)
	}
}

// 后序搜索
func (threadedBinaryTree *ThreadedBinaryTree) PostorderSearch(id int) *ThreadedBinaryTreeNode {
	if threadedBinaryTree.Root == nil{
		panic("This tree is empty!")
	}else{
		return threadedBinaryTree.Root.PostorderSearch(id)
	}
}