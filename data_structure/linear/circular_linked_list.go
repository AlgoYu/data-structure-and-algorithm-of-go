/*
循环链表：循环链表是链表的最后一个数据节点的指针指向第一个节点的闭环数据结构
主要思想：
	1. 创建一个数据结构，有数据域和指针域。
	2. 数据域用来存储数据，指针域用来指向下一个数据。
	3. 使用指针域的前后指针完成链表的增、删、改、查。
	4. 最后一个数据节点的指针指向第一个数据节点。
*/
package linear

import (
	"fmt"
)

type CircularLinkedListNode struct {
	Id   int
	Data int
	Next *CircularLinkedListNode
}

type CircularLinkedList struct {
	length uint
	next   *CircularLinkedListNode
}

// 初始化链表
func (circularLinkedList *CircularLinkedList) Init()  {
	circularLinkedList.next = &CircularLinkedListNode{
		Id:   0,
		Data: 0,
		Next: nil,
	}
	circularLinkedList.next.Next = circularLinkedList.next
}

// 得到长度
func (circularLinkedList *CircularLinkedList) GetLength() uint {
	return circularLinkedList.length
}

// 增加节点
func (circularLinkedList *CircularLinkedList) AddNode(node *CircularLinkedListNode)  {
	if circularLinkedList.GetNode(node.Id) != nil{
		fmt.Println("The node already exists")
		return
	}
	temp := circularLinkedList.next
	for temp.Next != circularLinkedList.next{
		temp = temp.Next
	}
	node.Next = temp.Next
	temp.Next = node
	circularLinkedList.length++
}

// 增加有序节点
func (circularLinkedList *CircularLinkedList) AddOrderNode(node *CircularLinkedListNode)  {
	// 判断是否已经存在
	if circularLinkedList.GetNode(node.Id) != nil{
		fmt.Println("The node already exists")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := circularLinkedList.next
	// 在链表中找到比当前需要插入的节点ID小的节点
	for temp.Next != circularLinkedList.next && temp.Next.Id < node.Id{
		temp = temp.Next
	}
	// 插入节点
	node.Next = temp.Next
	temp.Next = node
	circularLinkedList.length++
}

// 删除节点
func (circularLinkedList *CircularLinkedList) DeleteNode(id int)  {
	// 创建指针
	temp := circularLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for temp.Next != circularLinkedList.next && temp.Next.Id != id{
		temp = temp.Next
	}
	// 判断下一个节点是否为空
	if temp.Next != circularLinkedList.next{
		temp.Next = temp.Next.Next
		circularLinkedList.length--
	}
}

// 修改节点
func (circularLinkedList *CircularLinkedList) ModifyNode(node *CircularLinkedListNode)  {
	// 创建指针
	temp := circularLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for temp.Next != circularLinkedList.next && temp.Next.Id != node.Id{
		temp = temp.Next
	}
	// 判断下一个节点是否为空
	if temp.Next != circularLinkedList.next{
		// 这里可以替换整个节点也可以只赋值
		node.Next = temp.Next.Next
		temp.Next = node
	}
}

// 得到节点
func (circularLinkedList *CircularLinkedList) GetNode(id int) *CircularLinkedListNode {
	// 创建指针并指定到第一个空节点
	temp := circularLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for temp.Next != circularLinkedList.next && temp.Next.Id !=id{
		temp = temp.Next
	}
	if temp.Next != circularLinkedList.next{
		return temp.Next
	}
	return nil
}

// 打印链表
func (circularLinkedList *CircularLinkedList) PrintLinkedList(){
	// 判断链表是否为空
	if circularLinkedList.length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := circularLinkedList.next
	// 指针反复指向下一个非空节点并打印
	for temp.Next != circularLinkedList.next{
		temp = temp.Next
		fmt.Print("[",temp.Id,"]=",temp.Data," ")
	}
	fmt.Println()
}