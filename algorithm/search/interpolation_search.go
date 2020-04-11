/*
插值查找：插值查找是一种搜索算法，是在二分查找上做的优化，依然前提是数据已经有序，优化了选择中间值值的计算方法，每次比较中间的数来判断往左查找还是往右查找，找到则返回元素下标。
主要思想：
	1. 计算出插值的下标，选择中间值。
	2. 如果中间值等于要查找的元素则返回。
	3. 如果中间值不等于，则判断大小来选择往左查找还是右查找。
	4. 重复以上步骤直到找到，找不到则返回负数下标。
*/
package search

// 插值查找
func InterpolationSearch(array []int,value,left,right int) int {
	if left <= right && value >= array[left] && value <= array[right]{
		middle := left + (right - left) * (value - array[left]) / (array[right] - array[left])
		if array[middle] > value{
			return InterpolationSearch(array,value,left,middle-1)
		}else if array[middle] < value{
			return InterpolationSearch(array,value,middle+1,right)
		}else{
			return middle
		}
	}
	return -1
}