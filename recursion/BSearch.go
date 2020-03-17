package recursion

import (

)

func BSearch(ary []interface{}, v int) int {
	num := len(ary)
	return bSearch(ary, 0, num-1, v)
}

//从由小到大排列的ary中找到值为v的元素位置
func bSearch(ary []interface{}, begin, end, v int) int {
	mid := begin + ((end-begin)>>1)
	if mid == end {
		if ary[mid].(int) == v {
			return mid
		} else {
			return -1
		}
	}

	if ary[mid].(int) == v {
		return mid
	} else if (ary[mid].(int)<v) {
		return bSearch(ary , mid + 1, end, v)
	} else {
		return bSearch(ary, begin, mid - 1, v)
	}
}