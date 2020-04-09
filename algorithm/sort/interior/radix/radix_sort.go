/*
基数排序：基数排序属于基数类排序的一种，以空间换时间的思想，创建多个数组作为桶，把需要排序的数据按低位到高位取出放入相应的桶中再按桶从小到大取出放回，最终使其有序。
主要思想：
	1. 创建多个数组作为捅，创建桶计数器。
	2. 取出每个数据按低位到高位递增，并放入相应的桶中。
	3. 按排序规则从桶中取出元素，放回原始变量。
	4. 反复以上步骤完成排序。
*/
package radix

import "strconv"

// 基数排序
func RadixSort(array []int) []int {
	// 获取最大数位数
	max := array[0]
	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	max = len(strconv.Itoa(max))
	// 初始化桶和其他计数
	var barrel [10][] int
	for i := 0; i < len(barrel); i++ {
		barrel[i] = make([]int, len(array))
	}
	count := new([10]int)
	divisor := 1
	var index int
	time := 0
	for time <= max {
		// 按位取出比较存入相应的捅
		for i := 0; i < len(array); i++ {
			index = array[i] / divisor % 10
			barrel[index][count[index]] = array[i]
			count[index]++
		}
		divisor *= 10
		// 取出桶数据放回原始变量中
		index = 0
		for i := 0; i < len(count); i++ {
			for j := 0; j < count[i] ; j++ {
				array[index] = barrel[i][j]
				index++
			}
			// 计数清零
			count[i] = 0
		}
		time++
	}
	return array
}