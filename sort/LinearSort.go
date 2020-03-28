package sort

import (
	"fmt"
)


//桶排序
func BuckSort(ary []interface{}) {
	//遍历元素找到最大值
	num := len(ary)
	if num <= 1 {
		return
	}
	max := ary[0].(int)
	for i := 1; i< num; i++ {
		if ary[i].(int) > max {
			max = ary[i].(int)
		}
	} 

	//声明一个二维切片,用于存储每个桶及其对应排序内容
	buckets := make([][]interface{}, num)

	//遍历ary,将元素copy到对应bucket
	for i := 0; i < num; i++ {
		//确定当前元素所在的bucket
		index := ary[i].(int)*(num-1)/max
		buckets[index] = append(buckets[index], ary[i])
	}

	//遍历 bucket,对每个bucket上的元素进行排序
	end := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			QuickSort(buckets[i])
			//将bucket 依次覆盖到ary
			copy(ary[end:], buckets[i])
			end += bucketLen
		}
	}
}

//计数排序. 比如高考分数排序(0分到700分)
func CountingSort(ary []interface{}) {
	//遍历元素找到最大值
	num := len(ary)
	if num <= 1 {
		return
	}
	max := ary[0].(int)
	for i := 1; i< num; i++ {
		if ary[i].(int) > max {
			max = ary[i].(int)
		}
	}

	//声明一个数组,下标就是元素值.从0开始
	countingAry := make([]int, max+1)
	fmt.Println(countingAry)
	//遍历待排序数组，将每个值的个数存入数组中
	for i := 0; i < num ; i ++ {
		v := ary[i].(int)
		countingAry[v] ++
	}
	for i := 1; i<= max ; i ++ {
		countingAry[i] += countingAry[i-1]
	}

	fmt.Println("countingary ----")
	fmt.Println(countingAry)
	//遍历原数组,生成排序数组
	newAry := make([]interface{}, num)
	for i := 0; i < num; i++ {
		v := ary[i]
		index := countingAry[v.(int)]
		countingAry[v.(int)]--

		newAry[index-1] = v
	}

	copy(ary, newAry)
}

//基数排序
func RadixSort(ary []interface{}) {
	
}