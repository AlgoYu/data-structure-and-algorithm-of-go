/*
哈夫曼树：哈夫曼树是一种树形结构，该树的带权路径长度达到最小，也称为最优二叉树。
主要思想：
	1. 把数组从小到大排序。
	2. 使用左右节点指针完成前、中、后序遍历、搜索,删除。
	3. 前序：父节点=>左子节点=>右子节点。
	4. 中序：左子节点=>父节点=>右子节点。
	5. 后序：左子节点=>右子节点=>父节点。
*/
package tree

import (
	"fmt"
)

// 哈弗曼树节点
type HuffmanTreeNode struct {
	Id int
	Data int
	Left *HuffmanTreeNode
	Right *HuffmanTreeNode
}

// 前序遍历
func (huffmanTreeNode *HuffmanTreeNode) PreorderTraversal()  {
	fmt.Print(huffmanTreeNode.Data,"=>")
	if huffmanTreeNode.Left!=nil{
		huffmanTreeNode.Left.PreorderTraversal()
	}
	if huffmanTreeNode.Right!=nil{
		huffmanTreeNode.Right.PreorderTraversal()
	}
}

type HuffmanTree struct {
	Root *HuffmanTreeNode
}

// 前序遍历
func (huffmanTree *HuffmanTree)PreorderTraversal()  {
	if huffmanTree.Root == nil{
		fmt.Println("This tree is empty!")
	}else{
		huffmanTree.Root.PreorderTraversal()
	}
}

// 排序
func (huffmanTree *HuffmanTree)NodeSort(nodes []*HuffmanTreeNode)  {
	for i:=1; i < len(nodes); i++ {
		index := i-1
		temp := nodes[i]
		for index>=0 && nodes[index].Data > temp.Data {
			nodes[index+1] = nodes[index]
			index--
		}
		nodes[index+1] = temp
	}
}

// 创建哈弗曼树
func (huffmanTree *HuffmanTree)GetHuffmanTree(array []int) *HuffmanTreeNode {
	var huffmanTreeNodes []*HuffmanTreeNode
	for _,value:= range array{
		huffmanTreeNodes = append(huffmanTreeNodes,&HuffmanTreeNode{Id: value,Data: value})
	}
	for len(huffmanTreeNodes)>1 {
		huffmanTree.NodeSort(huffmanTreeNodes)
		left := huffmanTreeNodes[0]
		right := huffmanTreeNodes[1]
		parent := new(HuffmanTreeNode)
		parent.Id = left.Id+right.Id
		parent.Data = left.Data+right.Data
		parent.Left = left
		parent.Right = right
		huffmanTreeNodes = huffmanTreeNodes[2:]
		huffmanTreeNodes = append(huffmanTreeNodes, parent)
	}
	huffmanTree.Root = huffmanTreeNodes[0]
	return huffmanTree.Root
}