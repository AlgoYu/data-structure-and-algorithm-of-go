/*
链表：链表是类似数组的一种数据结构，但链表在内存空间中非连续存储的数据结构，它通过指针来寻找下一个节点。
主要思想：
	1. 创建一个数据结构，有数据域和指针域。
	2. 数据域用来存储数据，指针域用来指向下一个数据。
	3. 使用指针域指向下一个数据节点，通过指针完成链表的增、删、改、查。
*/
package linear

import (
	"fmt"
)

type SingleLinkedListNode struct {
	Id   int
	Data int
	Next *SingleLinkedListNode
}

type SingleLinkedList struct {
	length uint
	next   *SingleLinkedListNode
}

// 初始化链表
func (singleLinkedList *SingleLinkedList) Init()  {
	singleLinkedList.next = &SingleLinkedListNode{
		Id:   0,
		Data: 0,
		Next: nil,
	}
}

// 得到长度
func (singleLinkedList *SingleLinkedList) GetLength() uint {
	return singleLinkedList.length
}

// 增加节点
func (singleLinkedList *SingleLinkedList) AddNode(node *SingleLinkedListNode)  {
	if singleLinkedList.GetNode(node.Id) != nil{
		fmt.Println("The node already exists")
		return
	}
	temp := singleLinkedList.next
	for{
		if temp.Next == nil{
			break
		}
		temp = temp.Next
	}
	temp.Next = node
	singleLinkedList.length++
}

// 增加有序节点
func (singleLinkedList *SingleLinkedList) AddOrderNode(node *SingleLinkedListNode)  {
	// 判断是否已经存在
	if singleLinkedList.GetNode(node.Id) != nil{
		fmt.Println("The node already exists")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := singleLinkedList.next
	// 在链表中找到比当前需要插入的节点ID小的节点
	for{
		if temp.Next == nil || temp.Next.Id > node.Id {
			break
		}
		temp = temp.Next
	}
	// 插入节点
	node.Next = temp.Next
	temp.Next = node
	singleLinkedList.length++
}

// 删除节点
func (singleLinkedList *SingleLinkedList) DeleteNode(id int)  {
	// 创建指针
	temp := singleLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.Next == nil || temp.Next.Id == id{
			break
		}
		temp = temp.Next
	}
	// 判断下一个节点是否为空
	if temp.Next != nil{
		temp.Next = temp.Next.Next
		singleLinkedList.length--
	}
}

// 修改节点
func (singleLinkedList *SingleLinkedList) ModifyNode(node *SingleLinkedListNode)  {
	// 创建指针
	temp := singleLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.Next == nil || temp.Next.Id == node.Id {
			break
		}
		temp = temp.Next
	}
	// 判断下一个节点是否为空
	if temp.Next != nil{
		// 这里可以替换整个节点也可以只赋值
		node.Next = temp.Next.Next
		temp.Next = node
	}
}

// 得到节点
func (singleLinkedList *SingleLinkedList) GetNode(id int) *SingleLinkedListNode {
	// 创建指针并指定到第一个空节点
	temp := singleLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.Next == nil || temp.Next.Id ==id{
			break
		}
		temp = temp.Next
	}
	return temp.Next
}

// 打印链表
func (singleLinkedList *SingleLinkedList) PrintLinkedList(){
	// 判断链表是否为空
	if singleLinkedList.length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := singleLinkedList.next
	// 指针反复指向下一个非空节点并打印
	for{
		if temp.Next == nil{
			break
		}
		temp = temp.Next
		fmt.Print("[",temp.Id,"]=",temp.Data," ")
	}
	fmt.Println()
}