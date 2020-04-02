/*
八皇后：在国际象棋上放置八个皇后，使其不能互相攻击，即任意两个皇后都不能处于同一行、同一列或同一斜线上。
主要思想：
	1. 使用递归的方法来一行一行地放置皇后。
	2. 每次判断是否与前面的皇后处于同一行、同一列、同一斜线。
	3. 由于递归的回溯特性可以把所有的组合全部找出。
*/
package other

import (
	"fmt"
	"math"
)

type EightQueen struct {
	count int
	max int
	size int
	queens []int
}

// 初始化棋盘大小和皇后数量
func (eightQueen *EightQueen) Init(max,size int){
	eightQueen.max = max
	eightQueen.size = size
	eightQueen.count = 0
	eightQueen.queens = make([]int,8)
}

// 放入皇后
func (eightQueen *EightQueen) PutQueen(n int){
	// 如果已经放满皇后则打印
	if n == eightQueen.max{
		eightQueen.count++
		eightQueen.PrintQueens()
		return
	}
	// 尝试放入每一列
	for i:=0; i<eightQueen.size ; i++ {
		eightQueen.queens[n] = i
		if eightQueen.IsValid(n){
			eightQueen.PutQueen(n+1)
		}
	}
}

// 判断当前皇后位置是否有效
func (eightQueen *EightQueen) IsValid(n int) bool{
	// 遍历皇后
	for i:=0; i<n; i++ {
		// 如果当前皇后位置与前面的皇后列相同（处在同一列）或者行列差的绝对值相同（处在一个同一斜线），则无效。
		if eightQueen.queens[i] == eightQueen.queens[n] || math.Abs(float64(n-i)) == math.Abs(float64(eightQueen.queens[i]-eightQueen.queens[n])){
			return false
		}
	}
	return true
}

// 打印皇后
func (eightQueen *EightQueen) PrintQueens(){
	fmt.Println("The ",eightQueen.count," kind:")
	for i := 0; i < len(eightQueen.queens); i++{
		fmt.Print("[",i,"]=",eightQueen.queens[i]," ")
	}
	fmt.Println()
}