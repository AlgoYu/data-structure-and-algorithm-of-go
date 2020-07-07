/*
Floyd算法：给定一组强连通图，使用Floyd算法求出各个节点到其余节点的最短路径。
主要思想：
	1. 创建一个结构体，保存节点名称，节点距离，节点前驱节点。
	2. 创建3个指针，分别为：中间节点指针、出发节点指针、目标节点指针。
	3. 三层循环遍历把每个节点当做一次中间指针、算出节点为中间指针时，累加出发节点到中间节点的距离，中间节点到目标节点的距离。
	4. 如果这个距离比已存在的距离小，则更新。
*/
package other

import "fmt"

type FloydAlgorithm struct {
	nodes []rune
	dis   [][]int
	pre   [][]int
}

func NewFloydAlgorithm(nodes []rune, dis [][]int) *FloydAlgorithm {
	floydAlgorithm := &FloydAlgorithm{nodes: nodes, dis: dis}
	floydAlgorithm.pre = make([][]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		floydAlgorithm.pre[i] = make([]int, len(nodes))
		for j := 0; j < len(nodes); j++ {
			floydAlgorithm.pre[i][j] = i
		}
	}
	return floydAlgorithm
}

func (floydAlgorithm *FloydAlgorithm) PrintGraph() {
	for i := 0; i < len(floydAlgorithm.nodes); i++ {
		fmt.Println(floydAlgorithm.dis[i])
	}
	fmt.Println()
	for i := 0; i < len(floydAlgorithm.nodes); i++ {
		fmt.Println(floydAlgorithm.pre[i])
	}
}

func (floydAlgorithm *FloydAlgorithm) Floyd() {
	var leng int
	for i := 0; i < len(floydAlgorithm.nodes); i++ {
		for j := 0; j < len(floydAlgorithm.nodes); j++ {
			for k := 0; k < len(floydAlgorithm.nodes); k++ {
				leng = floydAlgorithm.dis[j][i] + floydAlgorithm.dis[i][k]
				if leng < floydAlgorithm.dis[j][k] {
					floydAlgorithm.dis[j][k] = leng
					floydAlgorithm.dis[k][j] = leng
					floydAlgorithm.pre[j][k] = i
					floydAlgorithm.pre[k][j] = i
				}
			}
		}
	}
}
