/*
二叉排序树：二叉排序树是一种树形结构，使用中序遍历进行排序，搜索和添加都有较好的性能。
主要思想：
	1. 创建左右子两个指针节点，左子节点比父节点小，右子节点比父节点大。
	2. 通过使用中序遍历从小到大输入有序节点。
	3. 判断当前节点比插入节点大或小判断往左递归查找还是往右递归查找完成二叉排序树的查找、增加、修改功能。
	4. 判断当前被删除节点是否为叶子节点、单子树节点、双子树节点以不同的方式完成删除功能。
*/
package tree

import "fmt"

type BinarySortTree struct {
	root *BinarySortTreeNode
}

// 增加节点
func (binarySortTree *BinarySortTree)AddNode(node *BinarySortTreeNode){
	if binarySortTree.root == nil{
		binarySortTree.root = node
	}else{
		binarySortTree.root.addNode(node)
	}
}

// 中序遍历
func (binarySortTree *BinarySortTree)InorderTraversal(){
	if binarySortTree.root!=nil{
		binarySortTree.root.inorderTraversal()
	}
}

// 搜索节点
func (binarySortTree *BinarySortTree)SearchNode(value int)*BinarySortTreeNode{
	if binarySortTree.root!=nil{
		return binarySortTree.root.searchNode(value)
	}else{
		return nil
	}
}

// 搜索父节点
func (binarySortTree *BinarySortTree)SearchParent(value int)*BinarySortTreeNode{
	if binarySortTree.root!=nil{
		return binarySortTree.root.searchParent(value)
	}else{
		return nil
	}
}

// 删除节点
func (binarySortTree *BinarySortTree)DeleteNode(value int){
	targetNode := binarySortTree.SearchNode(value)
	if targetNode!=nil{
		parentNode := binarySortTree.SearchParent(value)
		if binarySortTree.root==targetNode && targetNode.Left==nil && targetNode.Right==nil{
			binarySortTree.root = nil
		}else if targetNode.Left==nil && targetNode.Right==nil{
			if parentNode.Left!=nil && parentNode.Left == targetNode{
				parentNode.Left = nil
			}else{
				parentNode.Right = nil
			}
		}else if targetNode.Left==nil || targetNode.Right ==nil{
			if targetNode.Left != nil{
				if parentNode.Left!=nil && parentNode.Left == targetNode{
					parentNode.Left = targetNode.Left
				}else{
					parentNode.Right = targetNode.Left
				}
			}else{
				if parentNode.Left!=nil && parentNode.Left == targetNode{
					parentNode.Left = targetNode.Right
				}else{
					parentNode.Right = targetNode.Right
				}
			}
		}else{
			temp := targetNode.Right
			for temp.Left!=nil {
				temp = temp.Left
			}
			binarySortTree.DeleteNode(temp.Value)
			targetNode.Value = temp.Value
		}
	}
}

type BinarySortTreeNode struct {
	Value int
	Left *BinarySortTreeNode
	Right *BinarySortTreeNode
}

// 增加节点
func (binarySortTreeNode *BinarySortTreeNode)addNode(node *BinarySortTreeNode)  {
	if binarySortTreeNode.Value > node.Value{
		if binarySortTreeNode.Left==nil{
			binarySortTreeNode.Left = node
		}else{
			binarySortTreeNode.Left.addNode(node)
		}
	}else{
		if binarySortTreeNode.Right == nil{
			binarySortTreeNode.Right = node
		}else{
			binarySortTreeNode.Right.addNode(node)
		}
	}
}

// 查找节点
func (binarySortTreeNode *BinarySortTreeNode)searchNode(value int) *BinarySortTreeNode {
	if binarySortTreeNode.Value == value{
		return binarySortTreeNode
	}else if binarySortTreeNode.Left!=nil && binarySortTreeNode.Value > value{
		return binarySortTreeNode.Left.searchNode(value)
	}else if binarySortTreeNode.Right!=nil && binarySortTreeNode.Value < value{
		return binarySortTreeNode.Right.searchNode(value)
	}else{
		return nil
	}
}

// 查找父节点
func (binarySortTreeNode *BinarySortTreeNode)searchParent(value int) *BinarySortTreeNode {
	if binarySortTreeNode.Left!=nil && binarySortTreeNode.Left.Value == value || binarySortTreeNode.Right!=nil && binarySortTreeNode.Right.Value==value{
		return binarySortTreeNode
	}else if binarySortTreeNode.Left!=nil && binarySortTreeNode.Value > value{
		return binarySortTreeNode.Left.searchParent(value)
	}else if binarySortTreeNode.Right!=nil && binarySortTreeNode.Value < value{
		return binarySortTreeNode.Right.searchParent(value)
	}else {
		return nil
	}
}

// 中序遍历
func (binarySortTreeNode *BinarySortTreeNode)inorderTraversal()  {
	if binarySortTreeNode.Left!=nil{
		binarySortTreeNode.Left.inorderTraversal()
	}
	fmt.Print(binarySortTreeNode.Value,"=>")
	if binarySortTreeNode.Right!=nil{
		binarySortTreeNode.Right.inorderTraversal()
	}
}