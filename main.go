package main

import (
	"anydevelop.cn/algorithm/other"
	"anydevelop.cn/algorithm/sort/interior/selection"
	"anydevelop.cn/algorithm/sort/interior/swap"
	"anydevelop.cn/data_structure/linear"
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// 稀疏矩阵测试
	//SparseMatrixTest()
	// 队列测试
	//ArrayQueueTest()
	// 环形队列测试
	//ArrayCircularQueueTest()
	// 单链表测试
	//SingleLinkedListTest()
	// 双链表测试
	//TwoWayLinkedListTest()
	// 循环链表测试
	//CircularLinkedListTest()
	// 约瑟夫圆形测试
	//JosephusTest()
	// 栈测试
	//ArrayStackTest()
	// 中缀表达式计算器测试
	//CalculatorTest()
	// 后缀（逆波兰）表达式计算器测试
	//ReversePolishCalculatorTest()
	// 中缀表达式转后缀表达式
	//ExpressionConverterTest()
	// 阶乘和打印测试
	//RecursiveTest()
	// 迷宫测试
	//LabyrinthTest()
	// 八皇后测试
	//EightQueenTest()
	// 冒泡排序测试
	//BubbleSortTest()
	// 选择排序测试
	SelectionSortTest()
}

// 稀疏矩阵测试
func SparseMatrixTest()  {
	var sourceMatrix [10][10]int
	sourceMatrix[4][6] = 5
	sourceMatrix[6][3] = 2
	sourceMatrix[1][7] = 10
	arg := make([][]int,len(sourceMatrix))
	for i := range sourceMatrix {
		arg[i] = sourceMatrix[i][:]
	}
	fmt.Println("打印原始矩阵")
	linear.PrintMatrix(arg)
	sparseMatrix := linear.ConvertToSparseMatrix(arg)
	fmt.Println("打印稀疏矩阵")
	linear.PrintMatrix(sparseMatrix)
	fmt.Println("打印从稀疏矩阵中恢复的原始矩阵")
	linear.PrintMatrix(linear.RestoreMatrixFromSparseMatrix(sparseMatrix))
}

// 队列测试
func ArrayQueueTest(){
	loop := true
	arrayQueue := new(linear.ArrayQueue)
	arrayQueue.InitQueue(5)
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
			arrayQueue.PrintQueue()
		case 2:
			var data int
			fmt.Scanf("%d",&data)
			arrayQueue.AddQueue(data)
		case 3:
			fmt.Println(arrayQueue.OutQueue())
		case 4:
			fmt.Println(arrayQueue.PrintCurrentQueueHead())
		case 5:
			loop = false
		}
	}
}

// 环形队列测试
func ArrayCircularQueueTest(){
	loop := true
	arrayCircularQueue := new(linear.ArrayCircularQueue)
	arrayCircularQueue.InitQueue(5)
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
			arrayCircularQueue.PrintQueue()
		case 2:
			var data int
			fmt.Scanf("%d",&data)
			arrayCircularQueue.AddQueue(data)
		case 3:
			fmt.Println(arrayCircularQueue.OutQueue())
		case 4:
			fmt.Println(arrayCircularQueue.PrintCurrentQueueHead())
		case 5:
			loop = false
		}
	}
}

// 单链表测试
func SingleLinkedListTest()  {
	loop := true
	singleLinkedList := new(linear.SingleLinkedList)
	singleLinkedList.Init()
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
			singleLinkedList.PrintLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.AddNode(&linear.SingleLinkedListNode{Id: id, Data:rand.Int()})
		case 3:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.AddOrderNode(&linear.SingleLinkedListNode{Id: id, Data:rand.Int()})
		case 4:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.DeleteNode(id)
		case 5:
			var id int
			fmt.Scanf("%d",&id)
			singleLinkedList.ModifyNode(&linear.SingleLinkedListNode{Id: id, Data:rand.Int()})
		case 6:
			var id int
			fmt.Scanf("%d",&id)
			node := singleLinkedList.GetNode(id)
			if node==nil {
				fmt.Println("没有找到这个节点")
			}else{
				fmt.Println("[",node.Id,"]=",node.Data)
			}
		case 7:
			fmt.Println(singleLinkedList.GetLength())
		case 8:
			loop = false
		}
	}
}

// 双链表测试
func TwoWayLinkedListTest()  {
	loop := true
	twoWayLinkedList := new(linear.TwoWayLinkedList)
	twoWayLinkedList.Init()
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
			twoWayLinkedList.PrintLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.AddNode(&linear.TwoWayLinkedListNode{ID: id, Data:rand.Int()})
		case 3:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.AddOrderNode(&linear.TwoWayLinkedListNode{ID: id, Data:rand.Int()})
		case 4:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.DeleteNode(id)
		case 5:
			var id int
			fmt.Scanf("%d",&id)
			twoWayLinkedList.ModifyNode(&linear.TwoWayLinkedListNode{ID: id, Data:rand.Int()})
		case 6:
			var id int
			fmt.Scanf("%d",&id)
			node := twoWayLinkedList.GetNode(id)
			if node==nil {
				fmt.Println("没有找到这个节点")
			}else{
				fmt.Println("[",node.ID,"]=",node.Data)
			}
		case 7:
			fmt.Println(twoWayLinkedList.GetLength())
		case 8:
			loop = false
		}
	}
}

// 循环链表测试
func CircularLinkedListTest()  {
	loop := true
	circularLinkedList := new(linear.CircularLinkedList)
	circularLinkedList.Init()
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
			circularLinkedList.PrintLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.AddNode(&linear.CircularLinkedListNode{Id: id, Data:rand.Int()})
		case 3:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.AddOrderNode(&linear.CircularLinkedListNode{Id: id, Data:rand.Int()})
		case 4:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.DeleteNode(id)
		case 5:
			var id int
			fmt.Scanf("%d",&id)
			circularLinkedList.ModifyNode(&linear.CircularLinkedListNode{Id: id, Data:rand.Int()})
		case 6:
			var id int
			fmt.Scanf("%d",&id)
			node := circularLinkedList.GetNode(id)
			if node==nil {
				fmt.Println("没有找到这个节点")
			}else{
				fmt.Println("[",node.Id,"]=",node.Data)
			}
		case 7:
			fmt.Println(circularLinkedList.GetLength())
		case 8:
			loop = false
		}
	}
}

// 约瑟夫圆环测试
func JosephusTest()  {
	loop := true
	josephus := new(other.Josephus)
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
			josephus.PrintLinkedList()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			josephus.AddNode(&other.People{Id: id, Data:rand.Int()})
		case 3:
			var start,count int
			fmt.Scanf("%d", &start)
			fmt.Scanf("%d", &count)
			josephus.PrintJosephusCircular(start,count)
		case 4:
			loop = false
		}
	}
}

// 栈测试
func ArrayStackTest(){
	loop := true
	arrayStack := new(linear.ArrayStack)
	arrayStack.Init(5)
	for{
		if !loop{
			break
		}
		fmt.Println("输入1为打印栈")
		fmt.Println("输入2为数据入栈")
		fmt.Println("输入3为数据出栈")
		fmt.Println("输入4为退出")
		var input int
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			arrayStack.PrintarrayStack()
		case 2:
			var data int
			fmt.Scanf("%d",&data)
			arrayStack.Push(data)
		case 3:
			fmt.Println(arrayStack.Pop())
		case 4:
			loop = false
		}
	}
}

// 中缀表达式计算器测试
func CalculatorTest()  {
	calculator := new(other.Calculator)
	calculator.Init()
	fmt.Println(calculator.CalculateExpression("3+5-2*4/4"))
}

// 后缀（逆波兰）表达式计算器测试
func ReversePolishCalculatorTest()  {
	reversePolishCalculator := new(other.ReversePolishCalculator)
	reversePolishCalculator.Init()
	fmt.Println(reversePolishCalculator.CalculateReversePolishExpression("72 2 - 2 / 5 / 2 *"))
}

// 中缀表达式转后缀表达式
func ExpressionConverterTest()  {
	expressionConverter := new(other.ExpressionConverter)
	expressionConverter.Init()
	fmt.Println(expressionConverter.ConversionExpression("3 + 3 * ( 4 + 1 ) / 10"))
}

// 阶乘和打印测试
func RecursiveTest(){
	fmt.Println(other.Factorial(15))
	other.PrintAllNumOf(15)
}

// 迷宫测试
func LabyrinthTest()  {
	labyrinth := make([][]int,8)
	for i := range labyrinth {
		labyrinth[i] = make([]int,9)
		labyrinth[i][0] = 1
		labyrinth[i][8] = 1
	}
	for i:=0; i< len(labyrinth[0]); i++ {
		labyrinth[0][i] = 1
		labyrinth[7][i] = 1
	}
	labyrinth[3][1] = 1
	labyrinth[3][2] = 1
	labyrinth[3][3] = 1
	labyrinth[3][4] = 1
	/*labyrinth[3][5] = 1
	labyrinth[3][6] = 1
	labyrinth[3][7] = 1*/
	fmt.Println("走迷宫前：")
	other.PrintLabyrinth(labyrinth)
	fmt.Println("走迷宫后：")
	if other.DetectRoad(labyrinth,1,1,6,7){
		other.PrintLabyrinth(labyrinth)
	}else{
		fmt.Println("迷宫是死路!")
	}
}

// 八皇后测试
func EightQueenTest()  {
	eightQueen := new(other.EightQueen)
	eightQueen.Init(8,8)
	eightQueen.PutQueen(0)
}

// 冒泡排序测试
func BubbleSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	/*fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("冒泡排序后：")
	fmt.Println(swap.BubbleSort(array))*/
	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("优化冒泡排序后：")
	fmt.Println(swap.BubbleSortOptimize(array))
}

// 选择排序测试
func SelectionSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("选择排序后：")
	fmt.Println(selection.SelectionSort(array))
}