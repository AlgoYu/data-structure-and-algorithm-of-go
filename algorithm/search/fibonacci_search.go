/*
斐波那契查找：斐波那契查找是一种搜索算法，利用数学中的黄金分割法以及斐波那契序列来搜索查找值的下标。
主要思想：
	1. 先获取一个斐波那契数列。
	2. 使用数据长度从斐波那契数列中获取分割值下标。
	3. 填充数据以避免分割值越界。
	4. 计算出黄金分割点判断数值再次分割。
*/
package search

// 得到斐波那契数列
func GetFibonacciSequence(size int) []int {
	fibonacciSequence := make([]int,size)
	fibonacciSequence[0] = 1
	fibonacciSequence[1] = 1
	for i := 2; i < size;i++ {
		fibonacciSequence[i] = fibonacciSequence[i-1] + fibonacciSequence[i-2]
	}
	return fibonacciSequence
}

// 斐波那契查找
func FibonacciSearch(array []int,value int) int {
	low := 0
	high := len(array)-1
	// 黄金分割点
	k := 0
	// 获取斐波那契数列
	fibonacciSequence := GetFibonacciSequence(len(array))
	// 获取黄金分割点
	for high >fibonacciSequence[k] {
		k++
	}
	// 避免值越界
	temp := make([]int,fibonacciSequence[k])
	for i:=0; i < len(array);i++ {
		temp[i] = array[i]
	}
	for i:=len(array); i < len(temp); i++ {
		temp[i] = temp[i-1]
	}
	// 使用黄金分割点查找数值
	middle := 0
	for high >= low {
		middle = low + fibonacciSequence[k-1]-1
		if value > temp[middle]{
			low = middle+1
			k-=2
		}else if value < temp[middle]{
			high = middle-1
			k--
		}else{
			if high >= middle{
				return middle
			}else{
				return high
			}
		}
	}
	return -1
}