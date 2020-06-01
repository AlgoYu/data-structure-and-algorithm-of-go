/*
KMP搜索算法：现在有两个字符串：s1和s2，现在要你输出s2在s1当中第一次出现的位置，先扫描出s2的前后缀匹配表，再利用匹配表快速定位s2在s1的位置。
主要思想：
	1. 先扫描出s2的前后缀匹配表，使用数组保存起来。
	2. 扫描逐步匹配两个字符串的字符。
	3. 使用一个变量存储匹配的下标（同时也是匹配表的下标）。
	4. 匹配则递增，不匹配则一直递减到不匹配的下标，已匹配的不会重新匹配；
	5. 搜索所有字符串，直到下标等于s2长度。
*/
package search

func KMPSearch(source,match string) int {
	matchTable := KMPMatchTable(match)
	for i,j:=0,0; i < len(source); i++ {
		for j > 0 && source[i] != source[j] {
			j = matchTable[j-1]
		}
		if source[i] == source[j]{
			j++
		}
		if j == len(match){
			return i-j+1
		}
	}
	return -1
}

func KMPMatchTable(str string) []int {
	matchTable := make([]int,len(str))
	for i,j:=1,0; i < len(str); i++ {
		for j > 0 && str[i] != str[j] {
			j = matchTable[j-1]
		}
		if str[i] == str[j]{
			j++
		}
	}
	return matchTable
}