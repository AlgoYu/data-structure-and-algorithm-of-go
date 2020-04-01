/*
迷宫：使用矩阵来模拟迷宫地图，0为路，1为墙，2为已走过。设置起始点和终点。
主要思想：
	1. 使用递归的方法来探索迷宫出路。
	2. 判断终点是否为2或者全部遍历为结束条件。
	3. 当前的策略为下、右、左、上的顺序。
*/
package other

import "fmt"

// 打印迷宫
func PrintLabyrinth(labyrinth [][]int){
	for i:= 0; i< len(labyrinth); i++ {
		for j:=0; j< len(labyrinth[i]); j++ {
			fmt.Print(labyrinth[i][j]," ")
		}
		fmt.Println()
	}
}

// 探测迷宫出路
func DetectRoad(labyrinth [][]int,currentRow,currentColumn,endRow,endColumn int) bool {
	if labyrinth[endRow][endColumn]==2{
		return true
	}
	// 当前的策略为下、右、左、上的顺序
	if labyrinth[currentRow][currentColumn]==0{
		labyrinth[currentRow][currentColumn] = 2
		if DetectRoad(labyrinth,currentRow+1,currentColumn,endRow,endColumn){
			return true
		}else if DetectRoad(labyrinth,currentRow,currentColumn+1,endRow,endColumn){
			return true
		}else if DetectRoad(labyrinth,currentRow,currentColumn+1,endRow,endColumn){
			return true
		}else if DetectRoad(labyrinth,currentRow,currentColumn+1,endRow,endColumn){
			return true
		}else{
			labyrinth[currentRow][currentColumn] = 3
		}
	}
	return false
}