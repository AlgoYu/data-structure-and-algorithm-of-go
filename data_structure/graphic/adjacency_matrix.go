/*
邻接矩阵：玲姐矩阵是图结构的一种表示方式。
主要思想：
	1. 在创建的时候初始化矩阵大小。
	2. 矩阵纵向横向表示的节点顺序与增加的顺序一样。
	3. 使用0表示无连接，使用1表示连接。
*/
package graphic

import "fmt"

type AdjacencyMatrix struct {
	nodes []string
	edgeSize int
	matrix [][]int
}

func NewAdjacencyMatrix(size int) *AdjacencyMatrix {
	adjacencyMatrix := new(AdjacencyMatrix)
	adjacencyMatrix.edgeSize = 0
	adjacencyMatrix.nodes = make([]string,size)
	adjacencyMatrix.matrix = make([][]int,size)
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
	return adjacencyMatrix.matrix[row][column];
}

// 获取节点个数
func (adjacencyMatrix *AdjacencyMatrix)GetNodeSize() int {
	return len(adjacencyMatrix.nodes)
}

// 获取边数
func (adjacencyMatrix *AdjacencyMatrix)GetEdgeSize() int {
	return adjacencyMatrix.edgeSize
}

// 增加节点
func (adjacencyMatrix *AdjacencyMatrix)PrintGraphic()  {
	for _,value:= range adjacencyMatrix.matrix{
		fmt.Println(value)
	}
}