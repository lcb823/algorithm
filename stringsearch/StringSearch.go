package stringsearch

import (
	"errors"
	"math"
	// "fmt"
)

func IndexOfByBF(main string, sub string) int {
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

//demo,存在hash冲突
func hashForRK(a string) int {
	sum := 0
	for i := 0; i < len(a); i ++ {
		sum += int(rune(a[i]))
	}
	return sum
}

//通过设计hash算法来减少遍历
func IndexOfByRK(main string, sub string) int {
	m := len(main)
	n := len(sub)
	subHash := hashForRK(sub)
	var mainHash int
	for i := 0; i < m - n + 1; i ++ {
		if i == 0 {
			mainHash = hashForRK(main[0:n])
		} else {
			mainHash = mainHash - int(rune(main[i-1])) + int(rune(main[i+n-1]))	
		}
		if subHash == mainHash {
			if sub == main[i:n+i] {
				return i
			}		
		}
	}
	return -1
}


//从main中查找所有sub子串
func FindAllByBM(main, sub string) ([]int, error) {
	result := []int{}//存放全部匹配到的字符串的index
	m := len(main)
	n := len(sub)
	if m < n || m == 0 || n == 0{
		return nil, errors.New("not find")
	}

	//存放sub字符串中每个字符的位置，重复出现的字符记录最后一次出现的位置
	mapSub := make(map[int]int)
	for i := 0; i < n; i ++ {
		mapSub[int(sub[i])] = i
	}
	//预处理sub字符串，计算[]suffix,每个后缀对应的在子串的下标 []prefix,后缀是否是子串的前缀
	suffix, prefix := make([]int, n + 1), make([]bool, n + 1)
	for i :=0 ; i < n; i ++ {
		suffix[i] = -1
		prefix[i] = false
	}

	//初始化 suffix,prefix
	for i := 0; i < n - 1; i ++ {
		//求 sub[0:i+1] 与sub的公共后缀子串
		j := i
		k := 0
		for j >= 0 && sub[j] == sub[n-1-k] {
			k++
			suffix[k] = j
			j --
		}
		if j == -1 {
			prefix[k] = true
		}
	}

	// fmt.Println(suffix)
	// fmt.Println(prefix)

	var step int
	for i := 0; i < m - n + 1; i += step {//遍历比较,i是main中的字符下标
		var j int
		for j = n-1; j >= 0; j -- {//倒序比较
			if main[j+i] != sub[j] {
				break
			}
		}
		if j < 0 {//匹配成功
			result = append(result, i)
			step = n
			continue
		}
		x := j - mapSub[int(sub[j])]//坏字符情况下需要移动多少
		y := 0 //好后缀方法判断需要移动多少
		if j < n - 1 {//存在好后缀
			k := n - 1 - j //好后缀的长度
			if suffix[k] >= 0 {//sub中存在好后缀子串
				y = j - suffix[k] + 1
			} else {
				//如果sub中不存在好后缀子串，找到最长的后缀=sub的前缀
				isMatch := false
				for r := j + 2; r <= n - 1; r ++ {
					if prefix[n - r] {
						isMatch = true
						y = r
					}
				}
				if !isMatch {
					y = n//如果sub子串不能完整匹配到好后缀，并且好后缀中不存在能匹配sub前缀的字符，那么后移n
				}
			}
		}
		step = int(math.Max(float64(x), float64(y)))
		//可能存在step=0的情况
		if step == 0 {
			step ++
		}
	}

	if len(result) == 0 {
		return nil, errors.New("not found")
	}
	return result, nil
}