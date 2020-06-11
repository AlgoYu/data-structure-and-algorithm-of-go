/*
修路问题：给定一组无相连通图，要求以最少的权值连接所有的节点。
主要思想：
	1. 使用Prim算法，选择一个节点出发。
	2. 把当前节点添加到已连通节点中。
	3. 遍历已连通节点的所有邻接节点，把还未被连通且权值最小的一个加入连通节点。
	4. 循环以上步骤直至所有节点连通。
*/
package other

import "fmt"

type RGraph struct {
	size   int
	nodes  []rune
	matrix [][]int
}

// 打印图
func (rGraph *RGraph) PrintGraph() {
	for i := 0; i < rGraph.size; i++ {
		fmt.Println(rGraph.matrix[i])
	}
}

// 图的构造方法
func NewRGraph(size int) *RGraph {
	var mendRoad RGraph
	mendRoad.size = size
	mendRoad.matrix = make([][]int, size)
	for i := 0; i < size; i++ {
		mendRoad.matrix[i] = make([]int, size)
	}
	mendRoad.nodes = make([]rune, size)
	return &mendRoad
}

// 创建图
func CreateRGraph(size int, nodes []rune, matrix [][]int) *RGraph {
	rGraph := NewRGraph(size)
	for i := 0; i < len(matrix); i++ {
		rGraph.nodes[i] = nodes[i]
		for j := 0; j < len(matrix[i]); j++ {
			rGraph.matrix[i][j] = matrix[i][j]
		}
	}
	return rGraph
}

// 最小生成树
func MinTree(rGraph *RGraph, node int) {
	visited := make([]int, rGraph.size)
	visited[node] = 1
	var n1, n2 int
	minValue := 10000
	for i := 0; i < rGraph.size-1; i++ {
		for j := 0; j < rGraph.size; j++ {
			for k := 0; k < rGraph.size; k++ {
				if visited[j] == 1 && visited[k] == 0 && rGraph.matrix[j][k] < minValue {
					n1 = j
					n2 = k
					minValue = rGraph.matrix[j][k]
				}
			}
		}
		fmt.Println(string(rGraph.nodes[n1]), " to ", string(rGraph.nodes[n2]))
		visited[n2] = 1
		minValue = 10000
	}
}
