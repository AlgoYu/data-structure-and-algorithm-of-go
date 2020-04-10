/*
线性查找：线性查找是一种搜索算法，递增比较每一个元素，把正确的元素下标返回。
主要思想：
	1. 递增比较每一个元素。
	2. 找到返回下标。
	3. 没找到返回负数。
*/
package search

// 线性查找
func SeqSearch(data []int,value int) int {
	for i:=0; i < len(data); i++ {
		if data[i] == value {
			return i
		}
	}
	return -1
}