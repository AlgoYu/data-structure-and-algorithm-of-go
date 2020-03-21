/*
队列：队列是一个先进先出的数据结构，大部分情况下用来模拟排队的情况。
主要思想：
	1. 创建一个数组作为队列。
	2. 使用front作为前指针，使用tail最为尾指针。
	3. 前、尾指针保存的是队列的下标。
	4. 通过前、尾指针判断队列是否为空、满、增加数据、删除数据。
*/
package linear

import (
	"fmt"
)

func ArrayQueueTest(){
	loop := true
	arrayQueue := new(ArrayQueue)
	arrayQueue.initQueue(5)
	for{
		if !loop{
			break
		}
		fmt.Println("输入1为打印队列")
		fmt.Println("输入2为加入数据")
		fmt.Println("输入3为取出数据")
		fmt.Println("输入4为显示当前队列头部")
		fmt.Println("输入5为退出")
		var input int
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			arrayQueue.printQueue()
		case 2:
			var data int
			fmt.Scanf("%d",&data)
			arrayQueue.addQueue(data)
		case 3:
			fmt.Println(arrayQueue.outQueue())
		case 4:
			fmt.Println(arrayQueue.printCurrentQueueHead())
		case 5:
			loop = false
		}
	}
}

// 定义数组队列
type ArrayQueue struct {
	// 队列
	queue []int
	// 队列大小
	size int
	// 前指针
	front int
	// 尾指针
	tail int
}

// 初始化队列
func (queue *ArrayQueue) initQueue(size int){
	queue.size = size
	queue.queue = make([]int,size)
	queue.front, queue.tail = -1,-1
}

// 判断队列是否为空
func (queue *ArrayQueue) isEmpty()bool{
	return queue.front == queue.tail
}

// 判断队列是否已满
func (queue *ArrayQueue) isFull()bool{
	return queue.tail+1 >= queue.size
}

// 加入数据到队列尾部
func (queue *ArrayQueue) addQueue(data int) {
	if queue.isFull() {
		fmt.Println("Can't add queue because queue is full")
		return
	}
	queue.tail++
	queue.queue[queue.tail] = data
}

// 从队列头提取数据
func (queue *ArrayQueue) outQueue()int{
	if queue.isEmpty(){
		panic("Can't add queue because queue is full")
	}
	queue.front++
	return queue.queue[queue.front]
}

// 打印当前队列头部
func (queue *ArrayQueue) printCurrentQueueHead()int{
	fmt.Println("当前front",queue.front)
	if queue.isEmpty()||queue.front==queue.size-1{
		panic("Current not have data")
	}
	return queue.queue[queue.front+1]
}

// 打印队列
func (queue *ArrayQueue) printQueue()  {
	fmt.Println(queue.queue)
}