/*
暴力匹配算法：现在有两个字符串：𝑠1和𝑠2，现在要你输出𝑠2在𝑠1当中第一次出现的位置，
主要思想：
	1. 使用两个变量来分别记录两个字符串的下标。
	2. 如果匹配则同时前进，如果不匹配s1重置为第一次匹配的下标后一位，s2重置为0。
	3. 如果最后记录s2下标的变量值等于s2长度，说明匹配成功。
	4. 返回s1变量-s2变量的值则为第一次匹配的下标。
*/
package search

func BruteForceMatch(source,match string) int {
	i,j := len(source), len(match)
	for i<len(source) && j < len(match) {
		if source[i] == match[j]{
			i++
			j++
		}else{
			i = i - (j-1)
			j=0
		}
	}
	if j == len(match){
		j = i - j
	}
	return j
}