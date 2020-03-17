package bsearch

import (

)

//从由小到大的有序数组中找到任意等于v的元素
func BSearch1(ary []int, v int) int {
	num := len(ary)
	start, end :=0, num - 1
	for start <= end {
		//找中间点
		mid := start + ((end-start)>>1)
		midValue := ary[mid]
		if midValue == v {
			return mid
		} else if midValue < v {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

//从由小到大的有序数组中找到第一个等于v的元素
func BSearch2(ary []int, v int) int {
	num := len(ary)
	start, end := 0, num - 1
	for start <= end {
		mid := start + ((end-start)>>1)
		midValue := ary[mid]
		if midValue == v {
			//如果当前mid=0 || 当前的元素的前一个元素!=v
			if mid == 0 || ary[mid - 1] != v {
				return mid
			} else {
				end = mid - 1
			}
		} else if midValue < v {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

//从由小到大的有序数组中找出最后一个等于v的元素
func BSearch3(ary []int, v int) int {
	num := len(ary)
	start, end := 0, num - 1
	for start <= end {
		mid := start + ((end-start)>>1)
		midValue := ary[mid]
		if midValue == v {
			if mid == num - 1 || ary[mid + 1] != v {
				return mid
			} else {
				start = mid + 1
			}
		} else if midValue < v {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

//从由小到大的有序数组中找出第一个大于等于v的元素
func BSearch4(ary []int, v int) int {
	num := len(ary)
	start, end := 0, num - 1
	for start <= end {
		mid := start + ((end-start)>>1)
		midValue := ary[mid]
		if midValue >= v {
			if mid == 0 || ary[mid - 1] < v {
				return mid
			} else {
				end = mid - 1
			}
		} else {
			start = mid + 1
		}
	}
	return -1
}

//从由小到大的有序数组中找出最后一个小于等于v的元素
func BSearch5(ary []int, v int) int {
	num := len(ary)
	start, end := 0, num - 1
	for start <= end {
		mid := start + ((end-start)>>1)
		midValue := ary[mid]
		if midValue <= v {
			if mid == num - 1 || ary[mid + 1] > v {
				return mid
			} else {
				start = mid + 1
			}
		} else {
			end = mid - 1
		}
	}
	return -1
}