/*
01背包算法：给定一组物品，每种物品都有自己的重量和价格，在限定的总重量内，使得物品的总价格最高。
主要思想：
	1. 使用动态规划算法的思想，动态规划(dynamic programming)是运筹学的一个分支，是求解决策过程(decision process)最优化的数学方法。类似分治算法，不同的是，它将一个难以直接解决的大问题，分割成一些规模较小的重叠子问题，已经解决的子问题答案会被记录，不需要再次解决。
	2. 根据物品为行和容量为列进行二维数组建表。
	3. 逐个选择和不选择物品，进行填表。
	4. 求出此类问题的最佳组合。
*/
package other

import "fmt"

func KnapsackProblem(article [][]int,capacity int)  {
	table := make([][]int,len(article)+1)
	count := make([][]int,len(article)+1)
	for i:=0; i < len(table); i++ {
		table[i] = make([]int,capacity+1)
		count[i] = make([]int,capacity+1)
	}
	for i:=1; i < len(table); i++ {
		for j:=1; j < len(table[i]); j++ {
			if article[i-1][0] > j{
				table[i][j] = table[i-1][j]
			}else{
				if table[i-1][j] > article[i-1][1]+table[i-1][j-article[i-1][0]]{
					table[i][j] = table[i-1][j]
				}else{
					table[i][j] = article[i-1][1]+table[i-1][j-article[i-1][0]]
					count[i][j] = 1
				}
			}
		}
	}
	for i:=0; i< len(table);i++ {
		fmt.Println(table[i])
	}
	row := len(table)-1
	column := len(table[row])-1
	for row > 0 && column > 0{
		if count[row][column] == 1{
			fmt.Println("No.",row," article to knapsack")
			column -= article[row-1][0]
		}
		row--
	}
}