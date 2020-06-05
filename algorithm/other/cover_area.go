/*
覆盖地区：给定一组需要覆盖的地区和不同的频道信号，没个不同的频道信号覆盖不同的地区，求出最优解或者近似解。
主要思想：
	1. 使用贪心算法，每次比较频道覆盖需要地区个数，选择每次的最优解。
*/
package other

func CoverArea(allArea []string,channel map[string][]string) []string {
	selects := make([]string,0)
	var temp []string
	var maxKey string
	maxSize := 0
	for len(allArea)>0 {
		maxSize = 0
		for key,value := range channel{
			temp = temp[:0]
			temp = contain(allArea,value)
			if len(temp) > maxSize{
				maxKey = key
				maxSize = len(temp)
			}
		}
		if maxSize>0{
			selects = append(selects, maxKey)
			allArea = removeContain(allArea,channel[maxKey])
		}
	}
	return selects
}

func contain(all,part []string) []string {
	match := make([]string,0)
	for i:=0; i< len(part); i++ {
		for j:=0; j<len(all); j++ {
			if part[i] == all[j]{
				match = append(match, part[i])
			}
		}
	}
	return match
}

func removeContain(all,part []string) []string {
	for i:=0; i< len(part); i++ {
		for j:=0; j<len(all); j++ {
			if part[i] == all[j]{
				all = append(all[:j],all[j+1:]...)
				break
			}
		}
	}
	return all
}