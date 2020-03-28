/*
双向链表：双向链表在单链表的基础上为每个结点增加了指向前一个结点的指针。
主要思想：
	1. 创建一个数据结构，有数据域和指针域。
	2. 数据域用来存储数据，指针域有两个，分别用来指向前一个数据和下一个数据。
	3. 使用指针域的前后指针完成链表的增、删、改、查。
*/
package linear

import (
	"fmt"
)

type TwoWayLinkedListNode struct {
	ID   int
	Data int
	Pre  *TwoWayLinkedListNode
	Next *TwoWayLinkedListNode
}

type TwoWayLinkedList struct {
	length uint
	next   *TwoWayLinkedListNode
}

// 初始化链表
func (twoWayLinkedList *TwoWayLinkedList) Init()  {
	twoWayLinkedList.next = &TwoWayLinkedListNode{
		ID:   0,
		Data: 0,
		Next: nil,
	}
}

// 得到长度
func (twoWayLinkedList *TwoWayLinkedList) GetLength() uint {
	return twoWayLinkedList.length
}

// 增加节点
func (twoWayLinkedList *TwoWayLinkedList) AddNode(node *TwoWayLinkedListNode)  {
	if twoWayLinkedList.GetNode(node.ID) != nil{
		fmt.Println("The node already exists")
		return
	}
	temp := twoWayLinkedList.next
	for{
		if temp.Next == nil{
			break
		}
		temp = temp.Next
	}
	node.Pre = temp
	temp.Next = node
	twoWayLinkedList.length++
}

// 增加有序节点
func (twoWayLinkedList *TwoWayLinkedList) AddOrderNode(node *TwoWayLinkedListNode)  {
	// 判断是否已经存在
	if twoWayLinkedList.GetNode(node.ID) != nil{
		fmt.Println("The node already exists")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := twoWayLinkedList.next
	// 在链表中找到比当前需要插入的节点ID小的节点
	for{
		if temp.Next == nil || temp.Next.ID > node.ID {
			break
		}
		temp = temp.Next
	}
	// 插入节点
	if temp.Next != nil{
		temp.Next.Pre = node
	}
	node.Next = temp.Next
	node.Pre = temp
	temp.Next = node
	twoWayLinkedList.length++
}

// 删除节点
func (twoWayLinkedList *TwoWayLinkedList) DeleteNode(id int)  {
	if node := twoWayLinkedList.GetNode(id); node != nil{
		node.Pre.Next = node.Next
		if node.Next !=nil{
			node.Next.Pre = node.Pre
		}
		twoWayLinkedList.length--
	}
}

// 修改节点
func (twoWayLinkedList *TwoWayLinkedList) ModifyNode(node *TwoWayLinkedListNode)  {
	if valueNode:= twoWayLinkedList.GetNode(node.ID); valueNode!=nil{
		// 这里可以直接赋值也可以替换整个节点
		node.Pre = valueNode.Pre
		node.Next = valueNode.Next
		if node.Next != nil{
			node.Next.Pre = node
		}
		node.Pre.Next = node
	}
}

// 得到节点
func (twoWayLinkedList *TwoWayLinkedList) GetNode(id int) *TwoWayLinkedListNode {
	// 创建指针并指定到第一个空节点
	temp := twoWayLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.Next == nil || temp.Next.ID ==id{
			break
		}
		temp = temp.Next
	}
	return temp.Next
}

// 打印链表
func (twoWayLinkedList *TwoWayLinkedList) PrintLinkedList(){
	// 判断链表是否为空
	if twoWayLinkedList.length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := twoWayLinkedList.next
	// 指针反复指向下一个非空节点并打印
	for{
		if temp.Next == nil{
			break
		}
		temp = temp.Next
		fmt.Print("[",temp.ID,"]=",temp.Data," ")
	}
	fmt.Println()
}