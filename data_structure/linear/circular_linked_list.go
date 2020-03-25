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
	"math/rand"
)

func CircularLinkedListTest()  {
	loop := true
	circularLinkedList := new(CircularLinkedList)
	circularLinkedList.init()
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
			circularLinkedList.printLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.addNode(&CircularLinkedListNode{id: id, data:rand.Int()})
		case 3:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.addOrderNode(&CircularLinkedListNode{id: id, data:rand.Int()})
		case 4:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.deleteNode(id)
		case 5:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.modifyNode(&CircularLinkedListNode{id: id, data:rand.Int()})
		case 6:
			var id int
			fmt.Scanf("%d",&id)
			node := circularLinkedList.getNode(id)
			if node==nil {
				fmt.Println("没有找到这个节点")
			}else{
				fmt.Println("[",node.id,"]=",node.data)
			}
		case 7:
			fmt.Println(circularLinkedList.getLength())
		case 8:
			loop = false
		}
	}
}

type CircularLinkedListNode struct {
	id   int
	data int
	next *CircularLinkedListNode
}

type CircularLinkedList struct {
	length uint
	next   *CircularLinkedListNode
}

// 初始化链表
func (circularLinkedList *CircularLinkedList)init()  {
	circularLinkedList.next = &CircularLinkedListNode{
		id:   0,
		data: 0,
		next: nil,
	}
	circularLinkedList.next.next = circularLinkedList.next
}

// 得到长度
func (circularLinkedList *CircularLinkedList)getLength() uint {
	return circularLinkedList.length
}

// 增加节点
func (circularLinkedList *CircularLinkedList)addNode(node *CircularLinkedListNode)  {
	if circularLinkedList.getNode(node.id) != nil{
		fmt.Println("The node already exists")
		return
	}
	temp := circularLinkedList.next
	for{
		if temp.next == circularLinkedList.next{
			break
		}
		temp = temp.next
	}
	node.next = temp.next
	temp.next = node
	circularLinkedList.length++
}

// 增加有序节点
func (circularLinkedList *CircularLinkedList)addOrderNode(node *CircularLinkedListNode)  {
	// 判断是否已经存在
	if circularLinkedList.getNode(node.id) != nil{
		fmt.Println("The node already exists")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := circularLinkedList.next
	// 在链表中找到比当前需要插入的节点ID小的节点
	for{
		if temp.next == circularLinkedList.next || temp.next.id > node.id {
			break
		}
		temp = temp.next
	}
	// 插入节点
	node.next = temp.next
	temp.next = node
	circularLinkedList.length++
}

// 删除节点
func (circularLinkedList *CircularLinkedList)deleteNode(id int)  {
	// 创建指针
	temp := circularLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.next == circularLinkedList.next || temp.next.id == id{
			break
		}
		temp = temp.next
	}
	// 判断下一个节点是否为空
	if temp.next != circularLinkedList.next{
		temp.next = temp.next.next
		circularLinkedList.length--
	}
}

// 修改节点
func (circularLinkedList *CircularLinkedList)modifyNode(node *CircularLinkedListNode)  {
	// 创建指针
	temp := circularLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.next == circularLinkedList.next || temp.next.id == node.id {
			break
		}
		temp = temp.next
	}
	// 判断下一个节点是否为空
	if temp.next != circularLinkedList.next{
		// 这里可以替换整个节点也可以只赋值
		node.next = temp.next.next
		temp.next = node
	}
}

// 得到节点
func (circularLinkedList *CircularLinkedList)getNode(id int) *CircularLinkedListNode {
	// 创建指针并指定到第一个空节点
	temp := circularLinkedList.next
	// 反复指向下一个非空且值不正确的节点
	for{
		if temp.next == circularLinkedList.next || temp.next.id ==id{
			break
		}
		temp = temp.next
	}
	if temp.next != circularLinkedList.next{
		return temp.next
	}
	return nil
}

// 打印链表
func (circularLinkedList *CircularLinkedList)printLinkedList(){
	// 判断链表是否为空
	if circularLinkedList.length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	// 创建指针并指定到第一个空节点
	temp := circularLinkedList.next
	// 指针反复指向下一个非空节点并打印
	for{
		if temp.next == circularLinkedList.next{
			break
		}
		temp = temp.next
		fmt.Print("[",temp.id,"]=",temp.data," ")
	}
	fmt.Println()
}