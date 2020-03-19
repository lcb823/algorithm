package heap

import (

)

//设计一个数据结构，用来获取一组动态数据的median值

type HeapForMedian struct {
	HeapSmall *Heap
	HeapBig *Heap
	Capacity uint//容量
	Num uint//当前个数
	CompareFun func(v, toV interface{}) int
}

func NewHeapForMedian(capacity uint, compareFun func(v, toV interface{}) int) *HeapForMedian {
	return &HeapForMedian{
		HeapSmall: NewHeap(capacity/2+uint(1), compareFunc, "asc"),//小顶堆
		HeapBig: NewHeap(capacity/2+uint(1), compareFunc, "desc"),//大顶堆
		Capacity: capacity,
		Num: 0,
		CompareFun: compareFun,
	}
}

func (this *HeapForMedian) AddItem(item interface{}) bool{
	if this.Num >= this.Capacity {
		return false
	}

	if this.Num == 0 {
		this.HeapSmall.InsertNode(item)
	} else if this.Num == 1 {
		this.HeapBig.InsertNode(item)
		if this.CompareFun(item, this.HeapSmall.Data[1]) > 0 {
			this.HeapBig.Data[1], this.HeapSmall.Data[1] = this.HeapSmall.Data[1], this.HeapBig.Data[1]
		}
	} else {
		if this.CompareFun(item, this.HeapSmall.Data[1]) >= 0 {
			this.HeapSmall.InsertNode(item)
		} else {
			this.HeapBig.InsertNode(item)
		}

		if this.HeapSmall.Num - this.HeapBig.Num == 2 {
			top := this.HeapSmall.Data[1]
			this.HeapSmall.DeleteTop()

			this.HeapBig.InsertNode(top)
		}
		if this.HeapBig.Num - this.HeapSmall.Num == 2 {
			top := this.HeapBig.Data[1]
			this.HeapBig.DeleteTop()

			this.HeapSmall.InsertNode(top)
		}
	}
	this.Num++
	return true
}

func (this *HeapForMedian) FindMedian() interface{} {
	if this.Num == 0 {
		return nil
	}
	if this.HeapBig.Num == this.HeapSmall.Num {
		return float64(this.HeapSmall.Data[1].(int) + this.HeapBig.Data[1].(int))/2
	} else if this.HeapBig.Num > this.HeapSmall.Num{
		return this.HeapBig.Data[1]
	} else {
		return this.HeapSmall.Data[1]
	}
}