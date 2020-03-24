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
	"math/rand"
)

func TwoWayLinkedListTest()  {
	loop := true
	twoWayLinkedList := new(TwoWayLinkedList)
	twoWayLinkedList.init()
	rand.Seed(6666)
	for{
		if !loop{
			break
		}
		fmt.Println("输入1为打印链表")
		fmt.Println("输入2为加入链表节点")
		fmt.Println("输入3为加入链表有序节点")
		fmt.Println("输入4为删除链表节点")
		fmt.Println("输入5为修改链表节点")
		fmt.Println("输入6为查找节点")
		fmt.Println("输入7为打印链表长度")
		fmt.Println("输入8为退出")
		var input int
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			twoWayLinkedList.printLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.addNode(&TwoWayLinkedListNode{id: id, data:rand.Int()})
		case 3:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.addOrderNode(&TwoWayLinkedListNode{id: id, data:rand.Int()})
		case 4:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.deleteNode(id)
		case 5:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.modifyNode(&TwoWayLinkedListNode{id: id, data:rand.Int()})
		case 6:
			var id int
			fmt.Scanf("%d",&id)
			node := twoWayLinkedList.getNode(id)
			if node==nil {
				fmt.Println("没有找到这个节点")
			}else{
				fmt.Println("[",node.id,"]=",node.data)
			}
		case 7:
			fmt.Println(twoWayLinkedList.getLength())
		case 8:
			loop = false
		}
	}
}

type TwoWayLinkedListNode struct {
	id   int
	data int
	pre  *TwoWayLinkedListNode
	next *TwoWayLinkedListNode
}

type TwoWayLinkedList struct {
	length uint
	next   *TwoWayLinkedListNode
}

// 初始化链表
func (twoWayLinkedList *TwoWayLinkedList)init()  {
	twoWayLinkedList.next = &TwoWayLinkedListNode{
		id:   0,
		data: 0,
		next: nil,
	}
}

// 得到长度
func (twoWayLinkedList *TwoWayLinkedList)getLength() uint {
	return twoWayLinkedList.length
}

// 增加节点
func (twoWayLinkedList *TwoWayLinkedList)addNode(node *TwoWayLinkedListNode)  {
	if twoWayLinkedList.getNode(node.id) != nil{
		fmt.Println("The node already exists")
		return
	}
	temp := twoWayLinkedList.next
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next
	}
	node.pre = temp
	temp.next = node
	twoWayLinkedList.length++
}

// 增加有序节点
func (twoWayLinkedList *TwoWayLinkedList)addOrderNode(node *TwoWayLinkedListNode)  {
	// 判断是否已经存在
	if twoWayLinkedList.getNode(node.id) != nil{
		fmt.Println("The node already exists")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := twoWayLinkedList.next
	// 在链表中找到比当前需要插入的节点ID小的节点
	for{
		if temp.next == nil || temp.next.id > node.id {
			break
		}
		temp = temp.next
	}
	// 插入节点
	if temp.next != nil{
		temp.next.pre = node
	}
	node.next = temp.next
	temp.next = node
	twoWayLinkedList.length++
}

// 删除节点
func (twoWayLinkedList *TwoWayLinkedList)deleteNode(id int)  {
	if node := twoWayLinkedList.getNode(id); node != nil{
		node.pre.next = node.next
		if node.next!=nil{
			node.next.pre = node.pre
		}
		twoWayLinkedList.length--
	}
}

// 修改节点
func (twoWayLinkedList *TwoWayLinkedList)modifyNode(node *TwoWayLinkedListNode)  {
	if valueNode:= twoWayLinkedList.getNode(node.id); valueNode!=nil{
		// 这里可以直接赋值也可以替换整个节点
		node.pre = valueNode.pre
		node.next = valueNode.next
		if node.next != nil{
			node.next.pre = node
		}
		node.pre.next = node
	}
}

// 得到节点
func (twoWayLinkedList *TwoWayLinkedList)getNode(id int) *TwoWayLinkedListNode {
	// 创建指针并指定到第一个空节点
	temp := twoWayLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.next == nil || temp.next.id ==id{
			break
		}
		temp = temp.next
	}
	return temp.next
}

// 打印链表
func (twoWayLinkedList *TwoWayLinkedList)printLinkedList(){
	// 判断链表是否为空
	if twoWayLinkedList.length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := twoWayLinkedList.next
	// 指针反复指向下一个非空节点并打印
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next
		fmt.Print("[",temp.id,"]=",temp.data," ")
	}
	fmt.Println()
}