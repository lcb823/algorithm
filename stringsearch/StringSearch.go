package stringsearch

import (

)

func IndexOfByBf(main string, sub string) int {
	lenMain := len(main)
	lenSub := len(sub)

	lenCompare := lenMain - lenSub + 1
	if lenCompare < 0 {
		return -1
	}
	//暴力比对
	for i := 0; i < lenCompare; i ++ {
		if sub == main[i:lenSub+i] {
			return i
		}
	}
	return -1
}


//从main中查找所有sub子串
func FindAllByBM(main string, sub string) []int {
	result := []int//存放匹配到的字符串的index

	lenMain := len(main)
	lenSub := len(sub)
	lenCompare := lenMain - lenSub + 1
	if lenCompare < 0 {
		return result
	}

	for i := 0; i < lenCompare; i ++ {
		//todo 好后缀，坏字符
	}
}