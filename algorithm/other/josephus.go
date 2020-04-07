/*
约瑟夫问题：N个人围成一圈，从第一个开始报数，第M个将被杀掉，最后剩下一个，其余人都将被杀掉。
主要思想：
	1. 创建一个循环链表模拟人围成的圆。
	2. 数据域用来存储数据，指针域用来指向下一个数据。
	3. 最后一个数据节点的指针指向第一个数据节点。
	4. 约瑟夫问题使用一个辅助指针先指向环形指针的尾部，再跟随当前指针一起前进。
	5. 每次当前指针指向的数据节点需要出链表的时候，当前指针指向下一个，辅助指针指向当前指针。
*/
package other

import (
	"fmt"
)

// 数据节点
type People struct {
	Id   int
	Data int
	Next *People
}

// 约瑟夫环
type Josephus struct {
	length uint
	next *People
}

// 增加节点
func (josephus *Josephus) AddNode(node *People) {
	if josephus.next == nil{
		josephus.next = node
		josephus.next.Next = josephus.next
		josephus.length++
		return
	}
	temp := josephus.next
	for temp.Next != josephus.next{
		temp = temp.Next
	}
	node.Next = temp.Next
	temp.Next = node
	josephus.length++
}

// 打印节点
func (josephus *Josephus) PrintLinkedList() {
	if josephus.length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	temp := josephus.next
	count := josephus.length
	for count != 0{
		fmt.Print("[",temp.Id,"]=",temp.Data," ")
		temp = temp.Next
		count--
	}
	fmt.Println()
}

// 打印约瑟夫圆形
func (josephus *Josephus) PrintJosephusCircular(start,count int) {
	if josephus.length <= 1{
		return
	}
	// 辅助指针先指向环形指针尾部
	helper := josephus.next
	for helper.Next != josephus.next{
		helper = helper.Next
	}
	// 当前指针指向起始节点，辅助指针跟随在后。
	current := josephus.next
	for i := 0; i < start -1; i++ {
		current = current.Next
		helper = helper.Next
	}
	// 只要辅助指针不等于当前指针说明数据节点大于2
	for current != helper{
		// 当前指针指向需要出链表的节点，辅助指针跟随在后。
		for i:= 0; i < count-1; i++{
			current = current.Next
			helper = helper.Next
		}
		// 打印当前节点
		fmt.Print("[",current.Id,"]=",current.Data," ")
		// 当前指针指向下一个节点
		current = current.Next
		// 辅助指针指向当前节点
		helper.Next = current
	}
	fmt.Println()
	fmt.Println("存活：","[",current.Id,"]=",current.Data)
}