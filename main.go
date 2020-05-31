package main

import (
	"anydevelop.cn/algorithm/compression"
	"anydevelop.cn/algorithm/other"
	"anydevelop.cn/algorithm/search"
	"anydevelop.cn/algorithm/sort/interior/insertion"
	"anydevelop.cn/algorithm/sort/interior/merge"
	"anydevelop.cn/algorithm/sort/interior/radix"
	"anydevelop.cn/algorithm/sort/interior/selection"
	"anydevelop.cn/algorithm/sort/interior/swap"
	"anydevelop.cn/data_structure/graphic"
	"anydevelop.cn/data_structure/linear"
	"anydevelop.cn/data_structure/tree"
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
	//SelectionSortTest()
	// 插入排序测试
	//InsertionSortTest()
	// 希尔排序测试
	//ShellSortTest()
	// 快速排序测试
	//QuickSortTest()
	// 归并排序测试
	//MergeSortTest()
	// 基数排序测试
	//RadixSortTest()
	// 堆排序测试
	//HeapSortTest()
	// 线性查找测试
	//SeqSearchTest()
	// 二分查找测试
	//BinarySearchTest()
	// 插值查找测试
	//InterpolationSearchTest()
	// 斐波那契查找测试
	//FibonacciSearchTest()
	// 散列表测试
	//HashTableTest()
	// 二叉树前中后序遍历测试
	//BinaryTreeTraversalTest()
	// 二叉树前中后序搜索测试
	//BinaryTreeSearchTest()
	// 二叉树删除测试
	//BinaryTreeDeleteTest()
	// 顺序二叉树遍历测试
	//ArrayBinaryTreeTraversalTest()
	// 线索化二叉树测试
	//ThreadedBinaryTreeTest()
	// 线索化二叉树遍历测试
	//ThreadedBinaryTreeTraversalTest()
	// 哈夫曼树测试
	//HuffmanTreeTest()
	// 哈夫曼编码测试
	//HuffmanCodingTest()
	// 二叉排序树测试
	//BinarySortTreeTest()
	// ALV树测试
	//ALVTreeTest()
	// 邻接矩阵图测试
	//AdjacencyMatrixTest()
	// 汉诺塔测试
	//TowerOfHanoiTest()
	// 背包问题测试
	//KnapsackProblemTest()
	// 暴力匹配算法测试
	BruteForceMatchTest()
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
	arrayQueue := linear.NewArrayQueue(5)
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
			arrayQueue.Add(data)
		case 3:
			fmt.Println(arrayQueue.Pop())
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
	arrayCircularQueue := linear.NewArrayCircularQueue(5)
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

// 插入排序测试
func InsertionSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("插入移位排序后：")
	fmt.Println(insertion.InsertionSort(array))
}

// 希尔排序测试
func ShellSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	fmt.Println("排序前：")
	fmt.Println(array)
	/*fmt.Println("希尔交换排序后：")
	fmt.Println(insertion.ShellSwapSort(array))*/
	fmt.Println("希尔移位排序后：")
	fmt.Println(insertion.ShellMoveSort(array))
}

// 快速排序测试
func QuickSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("快速排序后：")
	swap.QuickSort(array,0, len(array)-1)
	fmt.Println(array)
}

// 归并排序测试
func MergeSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("归并排序后：")
	merge.MergeSort(array,0, len(array)-1,make([]int,20))
	fmt.Println(array)
}

// 基数排序测试
func RadixSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("基数排序后：")
	fmt.Println(radix.RadixSort(array))
}

// 堆排序测试
func HeapSortTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(500)
	}
	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("堆排序后：")
	fmt.Println(selection.HeapSort(array))
}

// 线性查找测试
func SeqSearchTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(10)
	}
	fmt.Println("切片：")
	fmt.Println(array)
	fmt.Println("线性搜索5下标")
	fmt.Println(search.SeqSearch(array,5))
}

// 线性查找测试
func BinarySearchTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(10)
	}
	swap.QuickSort(array,0, len(array)-1)
	fmt.Println("切片：")
	fmt.Println(array)
	fmt.Println("二分搜索5下标")
	fmt.Println(search.BinaryRecursiveSearch(array,5,0, len(array)-1))
	fmt.Println("二分非递归搜索5下标")
	fmt.Println(search.BinarySearch(array,5))
}

// 插值查找测试
func InterpolationSearchTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(10)
	}
	swap.QuickSort(array,0, len(array)-1)
	fmt.Println("切片：")
	fmt.Println(array)
	fmt.Println("插值搜索5下标")
	fmt.Println(search.InterpolationSearch(array,5,0,len(array)-1))
}

// 斐波那契查找测试
func FibonacciSearchTest()  {
	array := make([]int,20)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(array); i++ {
		array[i] = rand.Intn(10)
	}
	swap.QuickSort(array,0, len(array)-1)
	fmt.Println("切片：")
	fmt.Println(array)
	fmt.Println("斐波那契搜索5下标")
	fmt.Println(search.FibonacciSearch(array,5))
}

// 散列表测试
func HashTableTest()  {
	loop := true
	hashTable := new(linear.HashTable)
	hashTable.Init(20)
	rand.Seed(6666)
	for{
		if !loop{
			break
		}
		fmt.Println("输入1为打印散列表")
		fmt.Println("输入2为加入散列表节点")
		fmt.Println("输入3为加入散列表有序节点")
		fmt.Println("输入4为删除散列表节点")
		fmt.Println("输入5为修改散列表节点")
		fmt.Println("输入6为查找散列表节点")
		fmt.Println("输入7为退出")
		var input int
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			hashTable.PrintHashTable()
		case 2:
			var id int
			fmt.Scanf("%d",&id)
			hashTable.AddNode(&linear.HashTableNode{Id: id, Data:rand.Int()})
		case 3:
			var id int
			fmt.Scanf("%d",&id)
			hashTable.AddOrderNode(&linear.HashTableNode{Id: id, Data:rand.Int()})
		case 4:
			var id int
			fmt.Scanf("%d",&id)
			hashTable.DeleteNode(id)
		case 5:
			var id int
			fmt.Scanf("%d",&id)
			hashTable.ModifyNode(&linear.HashTableNode{Id: id, Data:rand.Int()})
		case 6:
			var id int
			fmt.Scanf("%d",&id)
			node := hashTable.GetNode(id)
			if node==nil {
				fmt.Println("没有找到这个节点")
			}else{
				fmt.Println("[",node.Id,"]=",node.Data)
			}
		case 7:
			loop = false
		}
	}
}

// 二叉树前中后序遍历测试
func BinaryTreeTraversalTest()  {
	root := &tree.BinaryTreeNode{Id: 1,Data: 100}
	node1 := &tree.BinaryTreeNode{Id: 2,Data: 100}
	node2 := &tree.BinaryTreeNode{Id: 3,Data: 100}
	node3 := &tree.BinaryTreeNode{Id: 4,Data: 100}
	node4 := &tree.BinaryTreeNode{Id: 5,Data: 100}
	node5 := &tree.BinaryTreeNode{Id: 6,Data: 100}
	root.Left = node1
	root.Right = node5
	node1.Left = node2
	node1.Right = node4
	node2.Left = node3
	binaryTree := &tree.BinaryTree{Root: root}
	fmt.Println("前序遍历：")
	binaryTree.PreorderTraversal()
	fmt.Println("中序遍历：")
	binaryTree.InorderTraversal()
	fmt.Println("后序遍历：")
	binaryTree.PostOrderTraversal()
}

// 二叉树前中后序搜索测试
func BinaryTreeSearchTest()  {
	root := &tree.BinaryTreeNode{Id: 1,Data: 100}
	node1 := &tree.BinaryTreeNode{Id: 2,Data: 100}
	node2 := &tree.BinaryTreeNode{Id: 3,Data: 100}
	node3 := &tree.BinaryTreeNode{Id: 4,Data: 100}
	node4 := &tree.BinaryTreeNode{Id: 5,Data: 100}
	node5 := &tree.BinaryTreeNode{Id: 6,Data: 100}
	root.Left = node1
	root.Right = node5
	node1.Left = node2
	node1.Right = node4
	node2.Left = node3
	binaryTree := &tree.BinaryTree{Root: root}
	fmt.Println("前序搜索：")
	fmt.Println(binaryTree.PreorderSearch(5))
	fmt.Println("中序搜索：")
	fmt.Println(binaryTree.PreorderSearch(5))
	fmt.Println("后序搜索：")
	fmt.Println(binaryTree.PreorderSearch(5))
}

// 二叉树前中后序删除测试
func BinaryTreeDeleteTest()  {
	root := &tree.BinaryTreeNode{Id: 1,Data: 100}
	node1 := &tree.BinaryTreeNode{Id: 2,Data: 100}
	node2 := &tree.BinaryTreeNode{Id: 3,Data: 100}
	node3 := &tree.BinaryTreeNode{Id: 4,Data: 100}
	node4 := &tree.BinaryTreeNode{Id: 5,Data: 100}
	node5 := &tree.BinaryTreeNode{Id: 6,Data: 100}
	root.Left = node1
	root.Right = node5
	node1.Left = node2
	node1.Right = node4
	node2.Left = node3
	binaryTree := &tree.BinaryTree{Root: root}
	fmt.Println("前序遍历：")
	binaryTree.PreorderTraversal()
	binaryTree.DeleteNode(5)
	fmt.Println()
	fmt.Println("前序遍历：")
	binaryTree.PreorderTraversal()
}

// 顺序二叉树前中后序遍历测试
func ArrayBinaryTreeTraversalTest()  {
	data := []int{1,2,3,4,5,6,7}
	arrayBinary := new(tree.ArrayBinaryTree)
	arrayBinary.Init(data)
	fmt.Println("前序遍历：")
	arrayBinary.PreorderTraversal(0)
	fmt.Println()
	fmt.Println("中序遍历：")
	arrayBinary.InorderTraversal(0)
	fmt.Println()
	fmt.Println("后序遍历：")
	arrayBinary.PostOrderTraversal(0)
}

// 二叉树前中后序遍历测试
func ThreadedBinaryTreeTest()  {
	root := &tree.ThreadedBinaryTreeNode{Id: 1,Data: 100}
	node1 := &tree.ThreadedBinaryTreeNode{Id: 2,Data: 100}
	node2 := &tree.ThreadedBinaryTreeNode{Id: 3,Data: 100}
	node3 := &tree.ThreadedBinaryTreeNode{Id: 4,Data: 100}
	node4 := &tree.ThreadedBinaryTreeNode{Id: 5,Data: 100}
	node5 := &tree.ThreadedBinaryTreeNode{Id: 6,Data: 100}
	root.Left = node1
	node1.Left = node2
	node1.Right = node4
	node2.Left = node3
	root.Right = node5
	threadedBinaryTree := &tree.ThreadedBinaryTree{Root: root}
	/*fmt.Println("前序遍历：")
	threadedBinaryTree.PreorderTraversal()
	fmt.Println()
	fmt.Println("前序线索化")
	threadedBinaryTree.PreorderThreaded(root)
	fmt.Println(node3.Left.Id)
	fmt.Println(node4.Left.Id)
	fmt.Println(node5.Left.Id)
	fmt.Println(node4.Right.Id)*/
	/*fmt.Println("中序遍历：")
	threadedBinaryTree.InorderTraversal()
	fmt.Println()
	fmt.Println("中序线索化")
	threadedBinaryTree.InorderThreaded(root)
	fmt.Println(node3.Right.Id)
	fmt.Println(node1.Right.Id)
	fmt.Println(node4.Left.Id)
	fmt.Println(node4.Right.Id)
	fmt.Println(node5.Left.Id)*/
	fmt.Println("后序遍历：")
	threadedBinaryTree.PostOrderTraversal()
	fmt.Println()
	fmt.Println("后序线索化")
	threadedBinaryTree.PostOrderThreaded(root)
	fmt.Println(node3.Right.Id)
	fmt.Println(node2.Right.Id)
	fmt.Println(node4.Left.Id)
	fmt.Println(node4.Right.Id)
}

// 线索化二叉树遍历测试
func ThreadedBinaryTreeTraversalTest()  {
	root := &tree.ThreadedBinaryTreeNode{Id: 1,Data: 100}
	node1 := &tree.ThreadedBinaryTreeNode{Id: 2,Data: 100}
	node2 := &tree.ThreadedBinaryTreeNode{Id: 3,Data: 100}
	node3 := &tree.ThreadedBinaryTreeNode{Id: 4,Data: 100}
	node4 := &tree.ThreadedBinaryTreeNode{Id: 5,Data: 100}
	node5 := &tree.ThreadedBinaryTreeNode{Id: 6,Data: 100}
	root.Left = node1
	node1.Left = node2
	node1.Right = node4
	node2.Left = node3
	root.Right = node5
	threadedBinaryTree := &tree.ThreadedBinaryTree{Root: root}
	fmt.Println("中序化遍历:")
	threadedBinaryTree.InorderThreaded(root)
	threadedBinaryTree.InorderThreadedTraversal()
}

// 哈夫曼树测试
func HuffmanTreeTest()  {
	array := []int{13,7,8,3,29,6,1}
	huffmanTree := new(tree.HuffmanTree)
	huffmanTree.GetHuffmanTree(array)
	huffmanTree.PreorderTraversal()
}

// 哈夫曼编码测试
func HuffmanCodingTest()  {
	str := "i like like like a golang do you like golang"
	var data []byte = []byte(str)
	fmt.Println("原始数据：",str)
	huffmanCoding := compression.NewHuffmanCoding()
	huffmanCompression := huffmanCoding.HuffmanCompression(data)
	fmt.Println("哈夫曼编码：",huffmanCompression)
	decompression := huffmanCoding.HuffmanDecompression(huffmanCompression)
	fmt.Println("恢复数据：",string(decompression))
	fmt.Println("哈夫曼压缩文件：","/Users/anydev/Downloads/1.jpg => /Users/anydev/Downloads/huffman.zip")
	huffmanCoding.HuffmanCompressionFile("/Users/anydev/Downloads/1.jpg","/Users/anydev/Downloads/huffman.zip")
	fmt.Println("哈夫曼解压文件：","/Users/anydev/Downloads/huffman.zip => /Users/anydev/Downloads/source.jpg")
	huffmanCoding.HuffmanDecompressionFile("/Users/anydev/Downloads/huffman.zip","/Users/anydev/Downloads/source.jpg")
}

// 二叉排序树测试
func BinarySortTreeTest()  {
	array := []int{3,5,9,1,2,7,6,9,4,11,15,23,31,-5,-1,-2,28,95,100,14}
	binarySortTree := new(tree.BinarySortTree)
	for _,value := range array{
		binarySortTree.AddNode(&tree.BinarySortTreeNode{Value: value})
	}
	fmt.Println("中序遍历：")
	binarySortTree.InorderTraversal()
	fmt.Println()
	fmt.Println("删除叶子节点：")
	binarySortTree.DeleteNode(-2)
	binarySortTree.InorderTraversal()
	fmt.Println()
	fmt.Println("删除单子树节点：")
	binarySortTree.DeleteNode(-5)
	binarySortTree.InorderTraversal()
	fmt.Println()
	fmt.Println("删除双子树节点：")
	binarySortTree.DeleteNode(3)
	binarySortTree.InorderTraversal()
}

// ALV树测试
func ALVTreeTest()  {
	array := []int{10,11,7,6,8,9}
	alvTree := new(tree.ALVTree)
	for _,value := range array{
		alvTree.AddNode(&tree.ALVTreeNode{Value: value})
	}
	fmt.Println("中序遍历：")
	alvTree.InorderTraversal()
	fmt.Println("\n获取树高度：",alvTree.GetHeight())
	fmt.Println("\n获取左子树高度：",alvTree.GetLeftHeight())
	fmt.Println("\n获取右子树高度：",alvTree.GetRightHeight())
}

// 邻接矩阵图测试
func AdjacencyMatrixTest()  {
	array := [...]string{"A","B","C","D","E"}
	adjacencyMatrix := graphic.NewAdjacencyMatrix(len(array))
	for i:=0; i < len(array);i++ {
		adjacencyMatrix.AddNode(array[i])
	}
	adjacencyMatrix.AddEdge(0,1,1)
	adjacencyMatrix.AddEdge(0,2,1)
	adjacencyMatrix.AddEdge(1,2,1)
	adjacencyMatrix.AddEdge(1,3,1)
	adjacencyMatrix.AddEdge(1,4,1)
	adjacencyMatrix.PrintGraphic()
	fmt.Println("深度优先遍历：")
	adjacencyMatrix.DFS()
	fmt.Println("\n广度优先遍历：")
	adjacencyMatrix.CleanVisit()
	adjacencyMatrix.BFS()
}

// 汉诺塔测试
func TowerOfHanoiTest() {
	other.TowerOfHanoi(3,'a','b','c')
}

// 背包问题测试
func KnapsackProblemTest()  {
	article := make([][]int,3)
	for i:=0; i< len(article); i++ {
		article[i] = make([]int,2)
	}
	article[0][0] = 1
	article[0][1] = 1500
	article[1][0] = 4
	article[1][1] = 3000
	article[2][0] = 3
	article[2][1] = 2000
	other.KnapsackProblem(article,4)
}

// 暴力匹配算法测试
func BruteForceMatchTest() {
	source := "I like Golang"
	match := "Golang"
	fmt.Println(other.BruteForceMatch(source,match))
}