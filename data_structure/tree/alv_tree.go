/*
ALV树：ALV树是一种树形结构，在二叉排序树上进行优化，通过节点左右旋转使其左右子树保持高度平衡（差值小于等于1）。
主要思想：
	1. 创建左右子两个指针节点，左子节点比父节点小，右子节点比父节点大。
	2. 通过使用中序遍历从小到大输入有序节点。
	3. 判断当前节点比插入节点大或小判断往左递归查找还是往右递归查找完成二叉排序树的查找、增加、修改功能。
	4. 判断当前被删除节点是否为叶子节点、单子树节点、双子树节点以不同的方式完成删除功能。
	5. 通过左右旋转保持ALV树左右子树高度差值不大于1。
*/
package tree

import (
	"fmt"
	"math"
)

type ALVTree struct {
	root *ALVTreeNode
}

// 获取树高度
func (alvTree *ALVTree)GetHeight()int{
	if alvTree.root!=nil{
		return alvTree.root.GetHeight()
	}
	return 0
}

// 获取左子树高度
func (alvTree *ALVTree)GetLeftHeight()int{
	if alvTree.root!=nil{
		return alvTree.root.GetLeftHeight()
	}
	return 0
}

// 获取右子树高度
func (alvTree *ALVTree)GetRightHeight()int{
	if alvTree.root!=nil{
		return alvTree.root.GetRightHeight()
	}
	return 0
}

// 增加节点
func (alvTree *ALVTree)AddNode(node *ALVTreeNode){
	if alvTree.root == nil{
		alvTree.root = node
	}else{
		alvTree.root.AddNode(node)
	}
	if alvTree.root.GetLeftHeight() - alvTree.root.GetRightHeight() > 1{
		if alvTree.root.Left.GetRightHeight() > alvTree.root.Left.GetLeftHeight(){
			alvTree.root.Left.LeftRotate()
		}
		alvTree.root.RightRotate()
	}
	if alvTree.root.GetRightHeight() - alvTree.root.GetLeftHeight() > 1{
		if alvTree.root.Right.GetLeftHeight() > alvTree.root.Right.GetRightHeight(){
			alvTree.root.Right.RightRotate()
		}
		alvTree.root.LeftRotate()
	}
}

// 中序遍历
func (alvTree *ALVTree)InorderTraversal(){
	if alvTree.root!=nil{
		alvTree.root.InorderTraversal()
	}
}

// 搜索节点
func (alvTree *ALVTree)SearchNode(value int)*ALVTreeNode{
	if alvTree.root!=nil{
		return alvTree.root.SearchNode(value)
	}else{
		return nil
	}
}

// 搜索父节点
func (alvTree *ALVTree)SearchParent(value int)*ALVTreeNode{
	if alvTree.root!=nil{
		return alvTree.root.SearchParent(value)
	}else{
		return nil
	}
}

// 删除节点
func (alvTree *ALVTree)DeleteNode(value int){
	targetNode := alvTree.SearchNode(value)
	if targetNode!=nil{
		parentNode := alvTree.SearchParent(value)
		if alvTree.root==targetNode && targetNode.Left==nil && targetNode.Right==nil{
			alvTree.root = nil
		}else if targetNode.Left==nil && targetNode.Right==nil{
			if parentNode.Left!=nil && parentNode.Left == targetNode{
				parentNode.Left = nil
			}else{
				parentNode.Right = nil
			}
		}else if targetNode.Left==nil || targetNode.Right ==nil{
			if parentNode==nil{
				if targetNode.Left!=nil{
					alvTree.root = targetNode.Left
				}else{
					alvTree.root = targetNode.Right
				}
			}else if targetNode.Left != nil{
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
			alvTree.DeleteNode(temp.Value)
			targetNode.Value = temp.Value
		}
	}
}

type ALVTreeNode struct {
	Value int
	Left *ALVTreeNode
	Right *ALVTreeNode
}

// 左旋
func (alvTreeNode *ALVTreeNode)LeftRotate() {
	temp := &ALVTreeNode{Value: alvTreeNode.Value}
	temp.Left = alvTreeNode.Left
	temp.Right = alvTreeNode.Right.Left
	alvTreeNode.Value = alvTreeNode.Right.Value
	alvTreeNode.Left = temp
	alvTreeNode.Right = alvTreeNode.Right.Right
}

// 右旋
func (alvTreeNode *ALVTreeNode)RightRotate() {
	temp := &ALVTreeNode{Value: alvTreeNode.Value}
	temp.Right = alvTreeNode.Right
	temp.Left = alvTreeNode.Left.Right
	alvTreeNode.Value = alvTreeNode.Left.Value
	alvTreeNode.Right = temp
	alvTreeNode.Left = alvTreeNode.Left.Left
}

// 获取树的高度
func (alvTreeNode *ALVTreeNode)GetHeight() int {
	var leftHeight,rightHeight int
	if alvTreeNode.Left!=nil{
		leftHeight = alvTreeNode.Left.GetHeight()
	}
	if alvTreeNode.Right!=nil{
		rightHeight = alvTreeNode.Right.GetHeight()
	}
	return int(math.Max(float64(leftHeight),float64(rightHeight)))+1
}

// 获取左子树的高度
func (alvTreeNode *ALVTreeNode)GetLeftHeight() int {
	if alvTreeNode.Left!=nil{
		return alvTreeNode.Left.GetHeight()
	}
	return 0
}

// 获取右子树的高度
func (alvTreeNode *ALVTreeNode)GetRightHeight() int {
	if alvTreeNode.Right!=nil{
		return alvTreeNode.Right.GetHeight()
	}
	return 0
}

// 增加节点
func (alvTreeNode *ALVTreeNode) AddNode(node *ALVTreeNode)  {
	if alvTreeNode.Value > node.Value{
		if alvTreeNode.Left==nil{
			alvTreeNode.Left = node
		}else{
			alvTreeNode.Left.AddNode(node)
		}
	}else{
		if alvTreeNode.Right == nil{
			alvTreeNode.Right = node
		}else{
			alvTreeNode.Right.AddNode(node)
		}
	}
}

// 查找节点
func (alvTreeNode *ALVTreeNode) SearchNode(value int) *ALVTreeNode {
	if alvTreeNode.Value == value{
		return alvTreeNode
	}else if alvTreeNode.Left!=nil && alvTreeNode.Value > value{
		return alvTreeNode.Left.SearchNode(value)
	}else if alvTreeNode.Right!=nil && alvTreeNode.Value < value{
		return alvTreeNode.Right.SearchNode(value)
	}else{
		return nil
	}
}

// 查找父节点
func (alvTreeNode *ALVTreeNode) SearchParent(value int) *ALVTreeNode {
	if alvTreeNode.Left!=nil && alvTreeNode.Left.Value == value || alvTreeNode.Right!=nil && alvTreeNode.Right.Value==value{
		return alvTreeNode
	}else if alvTreeNode.Left!=nil && alvTreeNode.Value > value{
		return alvTreeNode.Left.SearchParent(value)
	}else if alvTreeNode.Right!=nil && alvTreeNode.Value < value{
		return alvTreeNode.Right.SearchParent(value)
	}else {
		return nil
	}
}

// 中序遍历
func (alvTreeNode *ALVTreeNode) InorderTraversal()  {
	if alvTreeNode.Left!=nil{
		alvTreeNode.Left.InorderTraversal()
	}
	fmt.Print(alvTreeNode.Value,"=>")
	if alvTreeNode.Right!=nil{
		alvTreeNode.Right.InorderTraversal()
	}
}