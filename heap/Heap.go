package heap

import (
	"fmt"
)

type Heap struct {
	Data []interface{}
	CompareFun func(v, toV interface{}) int
	Max	uint
	Num uint
	Order string//默认大顶堆
}

func (this *Heap) compareFunWrap(v, toV interface{}) int {
	if this.Order == "asc" {
		return -this.CompareFun(v, toV)
	}
	return this.CompareFun(v, toV)
}

//假定data下标从1开始
func NewHeap(capacity uint, compareFun func(v, toV interface{}) int, order string) *Heap{
	return &Heap{
		Data: make([]interface{}, capacity + uint(1)),
		CompareFun: compareFun,
		Max: capacity,
		Num: 0,
		Order: order,
	}
}


//arr 下标从1开始
func BuildHeap(arr []interface{}, compareFun func(v, toV interface{}) int, order string) *Heap {
	num := len(arr)
	heap := &Heap{
		Data: arr,
		CompareFun: compareFun,
		Max: uint(num),
		Num: 0,
		Order: order,
	}

	for heap.Num < heap.Max {
		heap.Num ++ 
		heap.heapifyDownToUp()
	}

	return heap
}

//堆化,大顶堆,由上到下
func (this *Heap) heapifyUpToDown() {
	if this.Num == 1 {
		return
	}

	cur := uint(1)
	for cur*2 <= this.Num {//当子节点存在时执行循环
		toCompare := cur * 2
		if toCompare + uint(1) <= this.Num {
			//如果存在右子节点
			if this.compareFunWrap(this.Data[toCompare + uint(1)], this.Data[toCompare]) > 0 {
				toCompare = toCompare + uint(1)
			}
		}
		if this.compareFunWrap(this.Data[cur], this.Data[toCompare]) >= 0 {
			break
		}
		this.Data[cur], this.Data[toCompare] = this.Data[toCompare], this.Data[cur]
		cur = toCompare
	}
}

//堆化,大顶堆,由下到上
func (this *Heap) heapifyDownToUp() {
	if this.Num == 1 {
		return
	}

	i := this.Num
	// fmt.Println(i)
	for i > 1 {
		parent := i / 2
		if this.compareFunWrap(this.Data[i], this.Data[parent]) < 0 {
			break
		}
		this.Data[i], this.Data[parent] = this.Data[parent], this.Data[i]

		i = i/2
	}
}

func (this *Heap) InsertNode(v interface{}) {
	if this.Num == this.Max {
		return
	}
	this.Data[this.Num + 1] = v
	this.Num ++
	this.heapifyDownToUp()
}

func (this *Heap) DeleteTop() {
	if this.Num == uint(0) {
		return
	}
	if this.Num == uint(1) {
		this.Data[1] = nil
		this.Num = 0
		return
	}
	//将最后一个元素替换堆顶元素
	this.Data[1] = this.Data[this.Num]
	this.Num --
	
	this.heapifyUpToDown()
}

//堆排序
func (this *Heap) Sort() {
	tmp := this.Num
	for this.Num > 1 {
		//模拟删除,
		this.Data[1], this.Data[this.Num] = this.Data[this.Num], this.Data[1]
		this.Num--
		this.heapifyUpToDown()
	}
	this.Num = tmp
}

func (this *Heap) Print() {
	if this.Num == 0 {
		fmt.Println(" heap num is 0")
		return
	}
	var format string
	for i := uint(0); i < this.Num; i ++ {
		format += fmt.Sprintf(" %+v", i)
	}
	fmt.Println(format)
	fmt.Println(this.Num)
}
