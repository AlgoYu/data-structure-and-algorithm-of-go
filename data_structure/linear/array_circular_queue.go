/*
环形队列：环形队列是一个先进先出的数据结构，修复了队列空间重用的问题。
主要思想：
	1. 创建一个数组作为队列。
	2. 使用front作为前指针，使用tail最为尾指针。
	3. 前、尾指针保存的是队列的下标。
	4. 通过前、尾指针判断队列是否为空、满、增加数据、删除数据。
	5. 通过取余算法重置前、尾指针指向的下标。
*/
package linear

import (
	"fmt"
)

// 定义数组队列
type ArrayCircularQueue struct {
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
func (queue *ArrayCircularQueue) InitQueue(size int){
	queue.size = size
	queue.queue = make([]int,size)
	queue.front, queue.tail = 0,0
}

// 判断队列是否为空
func (queue *ArrayCircularQueue) IsEmpty()bool{
	return queue.front == queue.tail
}

// 判断队列是否已满
func (queue *ArrayCircularQueue) IsFull()bool{
	return (queue.tail+1) % queue.size == queue.front
}

// 加入数据到队列尾部
func (queue *ArrayCircularQueue) AddQueue(data int) {
	if queue.IsFull() {
		fmt.Println("Can't add queue because queue is full")
		return
	}
	queue.queue[queue.tail] = data
	queue.tail = (queue.tail+1) % queue.size
}

// 从队列头提取数据
func (queue *ArrayCircularQueue) OutQueue()int{
	if queue.IsEmpty(){
		panic("Can't add queue because queue is full")
	}
	data := queue.queue[queue.front]
	queue.front = (queue.front+1) % queue.size
	return data
}

// 得到当前队列长度
func (queue *ArrayCircularQueue) GetCurrentQueueLength() int{
	return (queue.tail+queue.size-queue.front) % queue.size
}

// 打印当前队列头部
func (queue *ArrayCircularQueue) PrintCurrentQueueHead()int{
	if queue.IsEmpty(){
		fmt.Println("Current not have Data")
	}
	return queue.queue[queue.front]
}

// 打印队列
func (queue *ArrayCircularQueue) PrintQueue()  {
	if queue.IsEmpty(){
		fmt.Println("Current not have Data")
	}
	for i:= queue.front; i < queue.front+queue.GetCurrentQueueLength();i++{
		fmt.Print(queue.queue[i%queue.size]," ")
	}
	fmt.Println()
}