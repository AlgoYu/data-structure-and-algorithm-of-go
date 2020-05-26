/*
二分查找：二分查找是一种搜索算法，前提是数据已经有序，每次比较中间的数来判断往左查找还是往右查找，找到则返回元素下标。
主要思想：
	1. 计算出中间数值的下标。
	2. 如果中间值等于要查找的元素则返回。
	3. 如果中间值不等于，则判断大小来选择往左查找还是右查找。
	4. 重复以上步骤直到找到，找不到则返回负数下标。
*/
package search

// 二分查找
func BinaryRecursiveSearch(array []int,value,left,right int) int {
	if left <= right{
		middle := (left+right)/2
		if array[middle] > value{
			return BinaryRecursiveSearch(array,value,left,middle-1)
		}else if array[middle] < value{
			return BinaryRecursiveSearch(array,value,middle+1,right)
		}else{
			return middle
		}
	}
	return -1
}

// 二分非递归查找
func BinarySearch(array []int,target int) int {
	left := 0
	right := len(array)-1
	var middle int
	for left <= right {
		middle = (left+right)/2
		if array[middle]==target{
			return middle
		}else if array[middle] > target{
			right = middle-1
		}else{
			left = middle+1
		}
	}
	return -1
}