/*
稀疏矩阵：是当一个矩阵中存在着大量重复且无意义的值，可以使用稀疏数组来存储这个矩阵，以减少占用的空间。
主要思想：
	1. 稀疏矩阵第一行[0][0]记录矩阵的行数，[0][1]记录矩阵的列数，[0][2]记录有效数据个数。
	2. 稀疏矩阵第二行开始记录有效数据，例如[1][0]记录有效数据的行数，[1][1]记录有效数据的列数，[1][2]记录有效数据数值。
	3. 重复第第二步。
*/
package main

import (
	"fmt"
)

func main()  {
	var sourceMatrix [10][10]int
	sourceMatrix[4][6] = 5
	sourceMatrix[6][3] = 2
	sourceMatrix[1][7] = 10
	arg := make([][]int,len(sourceMatrix))
	for i := range sourceMatrix {
		arg[i] = sourceMatrix[i][:]
	}
	fmt.Println("打印原始矩阵")
	printMatrix(arg)
	sparseMatrix := convertToSparseMatrix(arg)
	fmt.Println("打印稀疏矩阵")
	printMatrix(sparseMatrix)
	fmt.Println("打印从稀疏矩阵中恢复的原始矩阵")
	printMatrix(restoreMatrixFromSparseMatrix(sparseMatrix))
}

/* 把原始矩阵转换为稀疏矩阵 */
func convertToSparseMatrix(sourceMatrix [][]int) [][]int  {
	// validData记录矩阵的有效数据个数
	var validData int
	// 遍历矩阵获取有效数据的个数
	for _,r := range sourceMatrix{
		for _,c := range r{
			if c!=0{
				validData++
			}
		}
	}
	// 初始化行数为validData+1，列为3的矩阵。
	sparseMatrix := make([][]int,validData+1)
	// 初始化整行
	for i := range sparseMatrix{
		sparseMatrix[i] = make([]int,3)
	}
	// 记录行数
	sparseMatrix[0][0] = len(sourceMatrix)
	// 记录列数
	sparseMatrix[0][1] = len(sourceMatrix[0])
	// 记录有效数据个数
	sparseMatrix[0][2] = validData
	// 往稀疏矩阵存入有效值的行和列及值
	row := 0
	for i,r := range sourceMatrix{
		for j,c := range r{
			if c!=0{
				row++
				// 记录有效数据行数
				sparseMatrix[row][0] = i
				// 记录有效数据列数
				sparseMatrix[row][1] = j
				// 记录有效数据值
				sparseMatrix[row][2] = sourceMatrix[i][j]
			}
		}
	}
	return sparseMatrix
}

/* 把稀疏矩阵恢复为原始矩阵 */
func restoreMatrixFromSparseMatrix(sparseMatrix [][]int)[][]int{
	// 从稀疏矩阵中获取原始矩阵行数并初始化
	sourceMatrix := make([][]int,sparseMatrix[0][0])
	// 初始化每一行
	for i := range sourceMatrix{
		sourceMatrix[i] = make([]int,sparseMatrix[0][1])
	}
	// 从稀疏矩阵中获取有效数据行和列并赋值到原始矩阵
	for i := 0; i < sparseMatrix[0][2]; {
		i++
		sourceMatrix[sparseMatrix[i][0]][sparseMatrix[i][1]] = sparseMatrix[i][2]
	}
	return sourceMatrix
}

/* 打印矩阵 */
func printMatrix(matrix [][]int)  {
	for _,v:= range matrix{
		fmt.Println(v)
	}
}