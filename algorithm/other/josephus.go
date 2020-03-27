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
	"math/rand"
)

// 数据节点
type People struct {
	id int
	data int
	next *People
}

// 约瑟夫环
type Josephus struct {
	length uint
	next *People
}

// 增加节点
func (josephus *Josephus) addNode(node *People) {
	if josephus.next == nil{
		josephus.next = node
		josephus.next.next = josephus.next
		josephus.length++
		return
	}
	temp := josephus.next
	for{
		if temp.next == josephus.next{
			break
		}
		temp = temp.next
	}
	node.next = temp.next
	temp.next = node
	josephus.length++
}

// 打印节点
func (josephus *Josephus) printLinkedList() {
	if josephus.length <= 0{
		fmt.Println("LinkedList is empty!")
		return
	}
	temp := josephus.next
	count := josephus.length
	for{
		if count == 0{
			break
		}
		fmt.Print("[",temp.id,"]=",temp.data," ")
		temp = temp.next
		count--
	}
	fmt.Println()
}

// 打印约瑟夫圆形
func (josephus *Josephus) printJosephusCircular(start,count int) {
	if josephus.length <= 1{
		return
	}
	// 辅助指针先指向环形指针尾部
	helper := josephus.next
	for{
		if helper.next == josephus.next{
			break
		}
		helper = helper.next
	}
	// 当前指针指向起始节点，辅助指针跟随在后。
	current := josephus.next
	for i := 0; i < start -1; i++ {
		current = current.next
		helper = helper.next
	}
	// 只要辅助指针不等于当前指针说明数据节点大于2
	for{
		if current == helper{
			break
		}
		// 当前指针指向需要出链表的节点，辅助指针跟随在后。
		for i:= 0; i < count-1; i++{
			current = current.next
			helper = helper.next
		}
		// 打印当前节点
		fmt.Print("[",current.id,"]=",current.data," ")
		// 当前指针指向下一个节点
		current = current.next
		// 辅助指针指向当前节点
		helper.next = current
	}
	fmt.Println()
	fmt.Println("存活：","[",current.id,"]=",current.data)
}

func JosephusTest()  {
	loop := true
	josephus := new(Josephus)
	rand.Seed(6666)
	for{
		if !loop{
			break
		}
		fmt.Println("输入1为打印链表")
		fmt.Println("输入2为加入链表节点")
		fmt.Println("输入3为打印约瑟夫圆形")
		fmt.Println("输入4为退出")
		var input int
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			josephus.printLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			josephus.addNode(&People{id: id, data:rand.Int()})
		case 3:
			var start,count int
			fmt.Scanf("%d", &start)
			fmt.Scanf("%d", &count)
			josephus.printJosephusCircular(start,count)
		case 4:
			loop = false
		}
	}
}