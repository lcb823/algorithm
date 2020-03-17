package sort

import (
	// "fmt"
)

//将数组由小到大排列

//冒泡排序
//实现思路，从左至右，依次比较相邻的两个元素，如果不是由小到大，交换这两个元素
func Sort1(ary []interface{}) []interface{}{
	len := len(ary)
	
	for i := 0; i < len - 1; i++ {//冒泡次数
		isChange := false
		for j:= 0; j < len - i - 1; j++ {//第i次冒泡需比较的次数
			if ary[j].(int) > ary[j+1].(int) {
				ary[j], ary[j+1] = ary[j+1], ary[j]
				isChange = true
			}	
		}
		if !isChange {
			break
		}
	}
	return ary
}

//插入排序
//将数组区分为有序区和未排序区，从未排序区依次将元素插入有序区,初始状态有序区包含第一个元素
func Sort2(ary []interface{}) []interface{} {
	len := len(ary)

	for i := 1; i < len ; i++ {//遍历未排序区元素 i 是有序区和待排序区的分界点
		cur := ary[i] //保存当前值

		var j int
		for j = i-1; j >= 0;j-- {
			if cur.(int) < ary[j].(int) {
				ary[j+1] = ary[j]
			} else {
				break
			}
		}
		ary[j+1] = cur
	}

	return ary
}

//选择排序
//将数组分为有序区和未排序区，从未排序区找到最小的元素和未排序区第一个元素交换,初始状态都是未排序区
func Sort3(ary []interface{}) []interface{} {
	len := len(ary)
	for i := 0; i < len; i++ {//遍历未排序区元素
		cur := i
		for j := i+1; j< len; j++ {//找到最小的元素
			if ary[j].(int) < ary[cur].(int) {
				cur = j
			}
		}
		//交换
		if i != cur {
			ary[i], ary[cur] = ary[cur], ary[i]
		}
	}
	return ary
}

//希尔排序
func Sort4(ary []interface{}) []interface{} {
	len := len(ary)
	for gap := len/2; gap>0; gap = gap/2 {
		//对分组进行插入排序
		for i:= gap; i < len ; i++ {
			cur := ary[i]
			var j int
			for j = i-gap; j >= 0; j -= gap {
				if cur.(int) < ary[j].(int) {
					ary[j+gap] = ary[j]
				} else{
					break
				}
			}
			ary[j+gap] = cur
		}
	}
	return ary
}


//希尔排序
func Sort5(ary []interface{}) []interface{} {
	len := len(ary)

	for gap := len / 2; gap > 0; gap /= 2 {
		for i := gap; i < len; i++ {
			for j := i - gap; j >= 0 && ary[j].(int) > ary[j + gap].(int); j -= gap {
				ary[j], ary[j+gap] = ary[j+gap], ary[j]
			}
		}
	}

	return ary
}
