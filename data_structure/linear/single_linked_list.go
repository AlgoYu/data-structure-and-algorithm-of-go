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
	"math/rand"
)

func SingleLinkedListTest()  {
	loop := true
	singleLinkedList := new(SingleLinkedList)
	singleLinkedList.init()
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
			singleLinkedList.printLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.addNode(&SingleLinkedListNode{Id:id,Data:rand.Int()})
		case 3:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.addOrderNode(&SingleLinkedListNode{Id:id,Data:rand.Int()})
		case 4:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.deleteNode(id)
		case 5:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.modifyNode(&SingleLinkedListNode{Id:id,Data:rand.Int()})
		case 6:
			var id int
			fmt.Scanf("%d",&id)
			node := singleLinkedList.getNode(id)
			if node==nil {
				fmt.Println("没有找到这个节点")
			}else{
				fmt.Println("[",node.Id,"]=",node.Data)
			}
		case 7:
			fmt.Println(singleLinkedList.Length)
		case 8:
			loop = false
		}
	}
}

type SingleLinkedListNode struct {
	Id int
	Data int
	next *SingleLinkedListNode
}

type SingleLinkedList struct {
	Length uint
	next *SingleLinkedListNode
}

func (singleLinkedList *SingleLinkedList)init()  {
	singleLinkedList.next = &SingleLinkedListNode{
		Id:   0,
		Data: 0,
		next: nil,
	}
}

func (singleLinkedList *SingleLinkedList)addNode(node *SingleLinkedListNode)  {
	if singleLinkedList.getNode(node.Id) != nil{
		fmt.Println("The node already exists")
		return
	}
	temp := singleLinkedList.next
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next
	}
	temp.next = node
	singleLinkedList.Length++
}

// 增加有序节点
func (singleLinkedList *SingleLinkedList)addOrderNode(node *SingleLinkedListNode)  {
	// 判断是否已经存在
	if singleLinkedList.getNode(node.Id) != nil{
		fmt.Println("The node already exists")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := singleLinkedList.next
	// 在链表中找到比当前需要插入的节点ID小的节点
	for{
		if temp.next == nil || temp.next.Id > node.Id{
			break
		}
		temp = temp.next
	}
	// 插入节点
	node.next = temp.next
	temp.next = node
	singleLinkedList.Length++
}

// 删除节点
func (singleLinkedList *SingleLinkedList)deleteNode(id int)  {
	// 创建指针
	temp := singleLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.next == nil || temp.next.Id == id{
			break
		}
		temp = temp.next
	}
	// 判断下一个节点是否为空
	if temp.next != nil{
		temp.next = temp.next.next
		singleLinkedList.Length--
	}
}

// 修改节点
func (singleLinkedList *SingleLinkedList)modifyNode(node *SingleLinkedListNode)  {
	// 创建指针
	temp := singleLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.next == nil || temp.next.Id == node.Id{
			break
		}
		temp = temp.next
	}
	// 判断下一个节点是否为空
	if temp.next != nil{
		// 这里可以替换整个节点也可以只赋值
		node.next = temp.next.next
		temp.next = node
	}
}

// 得到节点
func (singleLinkedList *SingleLinkedList)getNode(id int) *SingleLinkedListNode {
	// 创建指针并指定到第一个空节点
	temp := singleLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.next == nil || temp.next.Id==id{
			break
		}
		temp = temp.next
	}
	return temp.next
}

// 打印链表
func (singleLinkedList *SingleLinkedList)printLinkedList(){
	// 判断链表是否为空
	if singleLinkedList.Length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := singleLinkedList.next
	// 指针反复指向下一个非空节点并打印
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next
		fmt.Print("[",temp.Id,"]=",temp.Data," ")
	}
	fmt.Println()
}