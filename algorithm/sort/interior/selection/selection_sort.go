/*
选择排序：选择排序属于选择类排序的一种，每次逐个比较选择一个最终元素，然后与前面的元素交换，最终使其有序。
主要思想：
	1. 使用临时变量保存第一个元素和下标，逐步递增。
	2. 指针指向第二个元素，然后与临时变量比较，判断是否符合排序规则，如果不符合，临时变量更改为指针指向元素的值与下标，逐步递增。
	3. 每一趟临时变量都会选择一个最终元素。
	4. 把最终元素替换与前面的元素替换。
*/
package selection

// 选择排序
func SelectionSort(array []int) []int {
	var temp,index int
	for i:=0; i< len(array); i++ {
		temp = array[i]
		index = i
		for j:=i+1; j<len(array) ; j++ {
			if temp>array[j]{
				temp = array[j]
				index = j
			}
		}
		// 如果指向的下标发生了变化则替换值
		if index != i{
			array[index] = array[i]
			array[i] = temp
		}
	}
	return array
}