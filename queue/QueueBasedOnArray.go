package queue

import (
	"fmt"
)

type QueueBasedOnArray struct {
	items		[]interface{}
	capacity	uint
	head		uint//指向队列的头部索引(索引位置保存第一个元素的值,0特殊,tail==0 && head==0 时,head索引处不存在值)
	tail		uint//指向队列的尾部索引(索引位置不保存值)
}

func NewQueueBasedOnArray(capacity uint) *QueueBasedOnArray {
	return &QueueBasedOnArray{
		items:	make([]interface{}, 0, capacity),
		capacity:	capacity,
		head:	uint(0),
		tail:	uint(0),
	}
}

func (this *QueueBasedOnArray) IsFull() bool {
	return this.tail >= this.capacity && this.head == uint(0)
}

func (this *QueueBasedOnArray) IsEmpty() bool {
	return this.head == this.tail
}

func (this *QueueBasedOnArray) Enqueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}

	//需要搬移数据
	if this.tail == this.capacity {
		for i := this.head; i < this.tail; i++ {
			this.items[i - this.head] = this.items[i]
		}
		this.tail = this.tail - this.head
		this.head = uint(0)
	}

	if this.tail >= uint(len(this.items)) {
		this.items = append(this.items, v)
	} else {
		this.items[this.tail] = v
	}
	this.tail++
	return true
}

func (this *QueueBasedOnArray) Dequeue() interface{} {
	if this.IsEmpty() {
		return nil
	}
	item := this.items[this.head]
	this.head++
	if this.head == this.tail {
		this.head = uint(0)
		this.tail = uint(0)
	}
	fmt.Println(this)
	return item
}