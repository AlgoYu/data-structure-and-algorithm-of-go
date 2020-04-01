/*
递归：递归是一种算法思想，函数直接或间接调用函数本身，则该函数称为递归函数。
主要思想：
	1. 函数直接或间接调用函数本身。
	2. 函数每次调用都必须向退出递归的条件趋近。
*/
package other

import "fmt"

// 阶乘，计算一个非负整数内所有数的乘积。
func Factorial(num int)int{
	if num == 1{
		return 1
	}
	return num * Factorial(num-1)
}

// 顺序打印一个非负整数内的所有数。
func PrintAllNumOf(num int){
	if num > 1{
		PrintAllNumOf(num-1)
	}
	fmt.Println(num)
}