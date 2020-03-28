/*
栈：栈是一种先入后出的数据结构。
主要思想：
	1. 创建一个数组当做栈空间。
	2. 创建一个指针指向栈的顶部。
	3. 使用指针完成栈的入栈出栈操作。
*/
package linear

import "fmt"

type ArrayStack struct {
	size int
	arrayStack []int
	top int
}

// 初始化栈
func (arrayStack *ArrayStack) Init(size int) {
	arrayStack.size = size
	arrayStack.arrayStack = make([]int,size)
	arrayStack.top = -1
}

// 初始化栈
func (arrayStack *ArrayStack) IsFull() bool {
	return arrayStack.top >= arrayStack.size-1
}

// 初始化栈
func (arrayStack *ArrayStack) IsEmpty() bool {
	return arrayStack.top == -1
}

// 入栈
func (arrayStack *ArrayStack) Push(data int) {
	if arrayStack.IsFull() {
		fmt.Println("arrayStack is full.")
		return
	}
	arrayStack.top++
	arrayStack.arrayStack[arrayStack.top] = data
}

// 出栈
func (arrayStack *ArrayStack) Peek()int{
	if arrayStack.IsEmpty() {
		panic("arrayStack is full.")
	}
	return arrayStack.arrayStack[arrayStack.top]
}

// 出栈
func (arrayStack *ArrayStack) Pop()int{
	if arrayStack.IsEmpty() {
		panic("arrayStack is full.")
	}
	data := arrayStack.arrayStack[arrayStack.top]
	arrayStack.top--
	return data
}

// 打印栈
func (arrayStack *ArrayStack) PrintarrayStack(){
	if arrayStack.IsEmpty() {
		fmt.Println("arrayStack is full.")
	}
	for i:=arrayStack.top; i >= 0; i--{
		fmt.Print(arrayStack.arrayStack[i]," ")
	}
	fmt.Println()
}