/*
公交站问题：给定一组无相连通图，要求以最少的权值连接所有的节点。
主要思想：
	1. 使用Kruskal算法。
	2. 扫描出所有的连接。
	3. 把连接按权值从小到大排序。
	4. 遍历连接，把与已选择的连接不构成回路的连接加入连通。
*/
package other

import "fmt"

const UNCONNECT = 10000

type BGraph struct {
	connect int
	nodes   []rune
	matrix  [][]int
}

// 打印图
func (bGraph *BGraph) PrintGraph() {
	for i := 0; i < len(bGraph.matrix); i++ {
		fmt.Println(bGraph.matrix[i])
	}
}

// 获取节点下标
func (bGraph *BGraph) GetNodeIndex(node rune) int {
	for i := 0; i < len(bGraph.nodes); i++ {
		if bGraph.nodes[i] == node {
			return i
		}
	}
	return -1
}

type CData struct {
	start  rune
	end    rune
	weight int
}

// 创建图
func CreateBGraph(nodes []rune, matrix [][]int) *BGraph {
	bGraph := new(BGraph)
	bGraph.nodes = make([]rune, len(nodes))
	bGraph.matrix = make([][]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		bGraph.matrix[i] = make([]int, len(nodes))
	}
	for i := 0; i < len(nodes); i++ {
		bGraph.nodes[i] = nodes[i]
		for j := 0; j < len(nodes); j++ {
			bGraph.matrix[i][j] = matrix[i][j]
		}
	}
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if matrix[i][j] > 0 && matrix[i][j] != UNCONNECT {
				bGraph.connect++
			}
		}
	}
	return bGraph
}

// kruskal
func Kruskal(graph *BGraph) {
	ends := make([]int, graph.connect)
	result := make([]CData, graph.connect)
	connects := getCDatas(graph)
	sortCData(connects)
	var index int
	for i := 0; i < len(connects); i++ {
		n1 := graph.GetNodeIndex(connects[i].start)
		n2 := graph.GetNodeIndex(connects[i].end)

		m := getEnd(ends, n1)
		n := getEnd(ends, n2)
		if m != n {
			ends[m] = n
			result[index] = connects[i]
			index++
		}
	}
	for i := 0; i < index; i++ {
		fmt.Println("<", string(result[i].start), ",", string(result[i].end), ">,weight=", result[i].weight)
	}
}

// 获取终点
func getEnd(ends []int, node int) int {
	for ends[node] != 0 {
		node = ends[node]
	}
	return node
}

// 获取连接信息
func getCDatas(graph *BGraph) []CData {
	ends := make([]CData, graph.connect)
	var index int
	for i := 0; i < len(graph.matrix); i++ {
		for j := i + 1; j < len(graph.matrix[i]); j++ {
			if graph.matrix[i][j] > 0 && graph.matrix[i][j] != UNCONNECT {
				var end CData
				end.start = graph.nodes[i]
				end.end = graph.nodes[j]
				end.weight = graph.matrix[i][j]
				ends[index] = end
				index++
			}
		}
	}
	return ends
}

// 对连接信息进行排序
func sortCData(array []CData) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j].weight > array[j+1].weight {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
