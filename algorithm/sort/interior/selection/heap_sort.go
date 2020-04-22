/*
堆排序：堆排序属于选择类排序的一种，使用堆这个数据结构的特性，反复选择子节点与父节点比较，交换，最终使其有序。
主要思想：
	1. 先把数组依据排序规则转换为一个最大（最小）堆。
	2. 把根节点换替换到后面的位置。
	3. 重复以上的步骤。
*/
package selection

// 转换为最大堆
func ConvertToMaxHeap(array []int,index,length int)  {
	temp := array[index]
	for i:=index*2+1; i < length; i++ {
		if i+1<length && array[i] < array[i+1]{
			i++
		}
		if array[i] > temp{
			array[index] = array[i]
			index = i
		}else{
			break
		}
	}
	array[index] = temp
}

// 转换为最小堆
func ConvertToMinHeap(array []int,index,length int)  {
	temp := array[index]
	for i:=index*2+1; i < length; i++ {
		if i+1<length && array[i] > array[i+1]{
			i++
		}
		if array[i] < temp{
			array[index] = array[i]
			index = i
		}else{
			break
		}
	}
	array[index] = temp
}

// 堆排序
func HeapSort(array []int) []int {
	for i:= len(array)/2-1; i >= 0; i-- {
		ConvertToMaxHeap(array,i, len(array))
	}
	var temp int
	for i:= len(array)-1; i > 0; i-- {
		temp = array[i]
		array[i] = array[0]
		array[0] = temp
		ConvertToMaxHeap(array,0,i)
	}
	return array
}