package sort

import (

)

func MergeSort(ary []interface{}){
	len := len(ary)
	if len <= 1 {
		return
	}
	mergeSort(ary, 0, len-1)
}

func mergeSort(ary []interface{}, start int, end int) {
	if start >= end {
		return
	}

	//找到中间节点
	mid := (start+end)/2

	mergeSort(ary, start, mid)
	mergeSort(ary, mid+1, end)

	merge(ary, start, mid, end)
}

func merge(ary []interface{}, start, mid, end int) {
	tmp := make([]interface{}, end-start+1)

	index := 0
	i := start
	j := mid + 1
	//
	for ; i <= mid && j <= end;index++ {
		if ary[j].(int) < ary[i].(int) {
			tmp[index] = ary[j]
			j++
		} else{
			tmp[index] = ary[i]
			i++
		}
	}

	for ; i <= mid; i++ {
		tmp[index] = ary[i]
		index++
	}

	for ; j <= end; j++ {
		tmp[index] = ary[j]
		index++
	}

	copy(ary[start:end+1], tmp)
}