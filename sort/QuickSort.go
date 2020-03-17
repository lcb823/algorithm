package sort

import (
	// "fmt"
)

func QuickSort(ary []interface{}) {
	len := len(ary)

	quickSort(ary, 0, len-1)
}

func quickSort(ary []interface{}, start, end int) {
	if start >= end {
		return
	}

	pivot := partion(ary, start, end)

	quickSort(ary, start, pivot-1)
	quickSort(ary, pivot+1, end)
}

func partion(ary []interface{}, start, end int) int {
	//默认选取end为分区点
	pivot := ary[end]

	index := start
	for i := start; i < end; i++ {
		if ary[i].(int) < pivot.(int) {//如果当前元素小于分区值,分区点的位置需要+1
			if i != index {//
				ary[i], ary[index] = ary[index], ary[i]	
			}
			index++
		}
	}

	ary[end], ary[index] = ary[index], ary[end]
	return index
}

//找到一个数组中第K大的元素
func FindK (ary []interface{}, k int) interface{}{
	len := len(ary)

	index := findK(ary, 0, len-1, k)

	return ary[index]
}

func findK(ary []interface{}, start, end, k int) int{
	pivot := find(ary, start, end)
	if pivot == k - 1 {
		return pivot
	}
	if pivot > k - 1 {
		return findK(ary, start, pivot-1, k)
	} else {
		return findK(ary, pivot +1, end, k)
	}
}

func find(ary []interface{}, start, end int) int {
	//默认选取end为分区点
	pivot := ary[end]

	index := start
	for i := start; i < end; i++ {
		if ary[i].(int) > pivot.(int) {//如果当前元素大于分区值,分区点的位置需要+1,由大到小排序
			if i != index {//
				ary[i], ary[index] = ary[index], ary[i]	
			}
			index++
		}
	}
	ary[end], ary[index] = ary[index], ary[end]

	return index
}