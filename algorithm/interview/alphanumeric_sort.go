/*
字母数字排序：给定一串字符串，里面包含字母和数字的组合，按数字优先，字母其次的排序顺序进行排序。
主要思想：
	1. 截取数字使用快排进行排序。
	2. 在数字顺序不变的情况下，使用插入排序进行对字母的排序。
*/
package interview

import "strconv"

func AlphanumericSort(array []string) {
	quickSort(array, 0, len(array)-1)
	insertSort(array)
}

// 保持数字不变的情况下对字母进行排序
func insertSort(array []string) {
	var index int
	var temp string
	for i := 1; i < len(array); i++ {
		index = i - 1
		temp = array[i]
		v1, _ := strconv.Atoi(temp[1:])
		v2, _ := strconv.Atoi(array[index][1:])
		for index >= 0 && v1 == v2 && array[index][:1] > temp[:1] {
			array[index+1] = array[index]
			index--
			if index >= 0 {
				v2, _ = strconv.Atoi(array[index][1:])
			}
		}
		array[index+1] = temp
	}
}

// 对数字进行排序
func quickSort(array []string, left, right int) {
	l := left
	r := right
	middle := (left + right) >> 1
	radix, _ := strconv.Atoi(array[middle][1:])
	for l < r {
		lv, _ := strconv.Atoi(array[l][1:])
		for lv < radix && l <= right {
			l++
			if l <= right {
				lv, _ = strconv.Atoi(array[l][1:])
			}
		}
		rv, _ := strconv.Atoi(array[r][1:])
		for rv > radix && r >= left {
			r--
			if r >= left {
				rv, _ = strconv.Atoi(array[r][1:])
			}
		}
		if l >= r {
			break
		}
		array[l], array[r] = array[r], array[l]
		if value, _ := strconv.Atoi(array[r][1:]); value == radix {
			l++
		}
		if value, _ := strconv.Atoi(array[r][1:]); value == radix {
			r--
		}
	}
	if l == r {
		l++
		r--
	}
	if r > left {
		quickSort(array, left, r)
	}
	if l < right {
		quickSort(array, l, right)
	}
}
