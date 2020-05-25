/*
邻接矩阵：玲姐矩阵是图结构的一种表示方式。
主要思想：
	1. 在创建的时候初始化矩阵大小。
	2. 矩阵纵向横向表示的节点顺序与增加的顺序一样。
	3. 使用0表示无连接，使用1表示连接。
*/
package graphic

import (
	"anydevelop.cn/data_structure/linear"
	"fmt"
)

type AdjacencyMatrix struct {
	nodes []string
	edgeSize int
	matrix [][]int
	visit []bool
}

func NewAdjacencyMatrix(size int) *AdjacencyMatrix {
	adjacencyMatrix := new(AdjacencyMatrix)
	adjacencyMatrix.edgeSize = 0
	adjacencyMatrix.matrix = make([][]int,size)
	adjacencyMatrix.visit = make([]bool,size)
	for i:=0; i < len(adjacencyMatrix.matrix); i++ {
		adjacencyMatrix.matrix[i] = make([]int,size)
	}
	return adjacencyMatrix
}

// 增加节点
func (adjacencyMatrix *AdjacencyMatrix)AddNode(node string)  {
	adjacencyMatrix.nodes =  append(adjacencyMatrix.nodes, node)
}

// 增加边
func (adjacencyMatrix *AdjacencyMatrix)AddEdge(row,column,value int)  {
	adjacencyMatrix.matrix[row][column] = value
	adjacencyMatrix.matrix[column][row] = value
	adjacencyMatrix.edgeSize++
}

// 获取节点
func (adjacencyMatrix *AdjacencyMatrix)GetNode(index int) string {
	return adjacencyMatrix.nodes[index]
}

// 获取值
func (adjacencyMatrix *AdjacencyMatrix)GetValue(row,column int) int {
	return adjacencyMatrix.matrix[row][column]
}

// 获取节点个数
func (adjacencyMatrix *AdjacencyMatrix)GetNodeSize() int {
	return len(adjacencyMatrix.nodes)
}

// 获取边数
func (adjacencyMatrix *AdjacencyMatrix)GetEdgeSize() int {
	return adjacencyMatrix.edgeSize
}

// 打印图
func (adjacencyMatrix *AdjacencyMatrix)PrintGraphic()  {
	for _,value:= range adjacencyMatrix.matrix{
		fmt.Println(value)
	}
}

// 深度优先遍历
func (adjacencyMatrix *AdjacencyMatrix)DepthFirstTraversal(row int){
	fmt.Print(adjacencyMatrix.nodes[row],"=>")
	adjacencyMatrix.visit[row] = true
	for column :=0; column < len(adjacencyMatrix.matrix[row]); column++ {
		if adjacencyMatrix.matrix[row][column]!=0 && !adjacencyMatrix.visit[column]{
			adjacencyMatrix.DepthFirstTraversal(column)
		}
	}
}

// 深度优先全遍历
func (adjacencyMatrix *AdjacencyMatrix)DFS()  {
	for i:=0; i < len(adjacencyMatrix.nodes); i++ {
		if !adjacencyMatrix.visit[i]{
			adjacencyMatrix.DepthFirstTraversal(i)
		}
	}
}

// 清除访问
func (adjacencyMatrix *AdjacencyMatrix)CleanVisit()  {
	for i:=0; i< len(adjacencyMatrix.visit);i++ {
		adjacencyMatrix.visit[i] = false
	}
}

// 广度优先全遍历
func (adjacencyMatrix *AdjacencyMatrix)BreadthFirstTraversal(row int)  {
	fmt.Print(adjacencyMatrix.nodes[row],"=>")
	adjacencyMatrix.visit[row] = true
	queue := linear.NewArrayQueue(len(adjacencyMatrix.nodes))
	queue.Add(row)
	for !queue.IsEmpty() {
		row = queue.Pop()
		for column:=0; column< len(adjacencyMatrix.matrix[row]);column++ {
			if adjacencyMatrix.matrix[row][column]!=0 && !adjacencyMatrix.visit[column]{
				fmt.Print(adjacencyMatrix.nodes[column],"=>")
				adjacencyMatrix.visit[column] = true
				queue.Add(column)
			}
		}
	}
}

// 广度优先全遍历
func (adjacencyMatrix *AdjacencyMatrix)BFS()  {
	for i:=0; i < len(adjacencyMatrix.nodes); i++ {
		if !adjacencyMatrix.visit[i]{
			adjacencyMatrix.BreadthFirstTraversal(i)
		}
	}
}