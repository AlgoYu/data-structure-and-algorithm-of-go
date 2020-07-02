/*
Dijkstra算法：给定一组无相连通图，求出从出发节点到全部节点的最短路径。
主要思想：
	1. 创建一个结构体来保存出发节点所有可连接节点的连接距离，所有节点的访问情况，以及计算过程中的前置节点。
	2. 获取一个当前距离最短且未被访问节点访问，计算以此节点作为中转节点访问其他节点的距离，如果小于已有的访问距离，则更新新的距离访问距离。
	3. 反复以上步骤，直至所有节点被访问。
*/
package other

import (
	"fmt"
	"strconv"
)

type DijkstraGraph struct {
	Nodes  []rune
	Matrix [][]int
}

// 打印图
func (dijkstraGraph *DijkstraGraph) PrintGraph() {
	for i := 0; i < len(dijkstraGraph.Matrix); i++ {
		fmt.Println(dijkstraGraph.Matrix[i])
	}
}

type VisitedNode struct {
	Visit []int
	Pre   []int
	Dis   []int
}

func NewVisitedNode(length, index int) *VisitedNode {
	visitedNode := &VisitedNode{}
	visitedNode.Visit = make([]int, length)
	visitedNode.Dis = make([]int, length)
	n := 65535
	for i := 0; i < len(visitedNode.Dis); i++ {
		visitedNode.Dis[i] = n
	}
	visitedNode.Pre = make([]int, length)
	visitedNode.Visit[index] = 1
	visitedNode.Dis[index] = 0
	return visitedNode
}

func (visitedNode *VisitedNode) IsVisited(index int) bool {
	return visitedNode.Visit[index] == 1
}

func (visitedNode *VisitedNode) UpdatePre(current, pre int) {
	visitedNode.Pre[current] = pre
}

func (visitedNode *VisitedNode) UpdateDis(index, length int) {
	visitedNode.Dis[index] = length
}

func (visitedNode *VisitedNode) GetDistance(index int) int {
	return visitedNode.Dis[index]
}

func Dijkstra(graph *DijkstraGraph, index int) []string {
	visitedNode := NewVisitedNode(len(graph.Nodes), index)
	update(graph, visitedNode, index)
	for i := 0; i < len(graph.Nodes); i++ {
		next := getNextNode(visitedNode)
		update(graph, visitedNode, next)
	}
	return getPath(visitedNode, graph, index)
}

func getPath(visitedNode *VisitedNode, graph *DijkstraGraph, index int) []string {
	var path []string
	min := 65535
	end := 0
	for i := 0; i < len(visitedNode.Dis); i++ {
		if i != index && visitedNode.IsVisited(i) && visitedNode.GetDistance(i) < min {
			min = visitedNode.GetDistance(i)
			end = i
		}
	}
	for end != index {
		path = append(path, "<"+string(graph.Nodes[end])+","+strconv.Itoa(visitedNode.GetDistance(end))+">")
		end = visitedNode.Pre[end]
	}
	return path
}

func update(graph *DijkstraGraph, visitedNode *VisitedNode, index int) {
	var dis int
	for i := 0; i < len(graph.Matrix[index]); i++ {
		dis = graph.Matrix[index][i] + visitedNode.GetDistance(index)
		if !visitedNode.IsVisited(i) && dis < visitedNode.GetDistance(i) {
			visitedNode.UpdatePre(i, index)
			visitedNode.UpdateDis(i, dis)
		}
	}
}

func getNextNode(visitedNode *VisitedNode) int {
	min := 65535
	var index int
	for i := 0; i < len(visitedNode.Visit); i++ {
		if !visitedNode.IsVisited(i) && visitedNode.GetDistance(i) < min {
			min = visitedNode.GetDistance(i)
			index = i
		}
	}
	visitedNode.Visit[index] = 1
	return index
}
