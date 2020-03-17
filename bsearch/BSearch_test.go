package bsearch

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestBSearch(t *testing.T) {
	slice := make([]int, 20)
	for i := 0; i < 20; i ++ {
		slice[i] = rand.Intn(20)
	}
	//再实现一次快速排序
	fmt.Println(slice)
	QuickSort(slice)
	fmt.Println(slice)

	v := rand.Intn(25)
	fmt.Println(fmt.Sprintf("从由小到大的有序数组中找到任意等于%v的元素----- begin", v))
	index := BSearch1(slice, v)
	fmt.Println(index)
	fmt.Println("从由小到大的有序数组中找到任意等于v的元素----- end")

	v = rand.Intn(25)
	fmt.Println(fmt.Sprintf("从由小到大的有序数组中找到第一个等于%v的元素----- begin", v))
	index = BSearch2(slice, v)
	fmt.Println(index)
	fmt.Println("从由小到大的有序数组中找到第一个等于v的元素----- end")

	v = rand.Intn(25)
	fmt.Println(fmt.Sprintf("从由小到大的有序数组中找出最后一个等于%v的元素----- begin", v))
	index = BSearch3(slice, v)
	fmt.Println(index)
	fmt.Println("从由小到大的有序数组中找出最后一个等于v的元素----- end")

	v = rand.Intn(25)
	fmt.Println(fmt.Sprintf("从由小到大的有序数组中找出第一个大于等于%v的元素----- begin", v))
	index = BSearch4(slice, v)
	fmt.Println(index)
	fmt.Println("从由小到大的有序数组中找出第一个大于等于v的元素----- end")

	v = rand.Intn(25)
	fmt.Println(fmt.Sprintf("从由小到大的有序数组中找出最后一个小于等于%v的元素----- begin", v))
	index = BSearch5(slice, v)
	fmt.Println(index)
	fmt.Println("从由小到大的有序数组中找出最后一个小于等于v的元素----- end")
}


func QuickSort(slice []int) {
	num := len(slice)
	if num <= 1 {
		return
	}

	quickSort(slice, 0, num - 1)
}

func quickSort(slice []int, start, end int) {
	if start >= end {
		return
	}

	//找分区点
	pivot := partion(slice, start, end)
	quickSort(slice, start, pivot - 1)
	quickSort(slice, pivot + 1, end)
}

func partion(slice []int, start, end int) int{
	pivotValue := slice[end]

	pivotIndex := start
	for i := start; i < end; i ++ {
		if pivotValue > slice[i] {
			//交换当前元素与index位置元素
			if pivotIndex != i {
				slice[pivotIndex], slice[i] = slice[i], slice[pivotIndex]
			}
			pivotIndex ++
		}
	} 
	slice[pivotIndex], slice[end] = slice[end], slice[pivotIndex]
	return pivotIndex
}