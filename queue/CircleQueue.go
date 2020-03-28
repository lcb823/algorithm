package queue

import (

)

type CircleQueue struct {
	items		[]interface{}
	capacity	uint
	head		uint
	tail		uint
}

func NewCircleQueue(capacity uint) *CircleQueue {
	return &CircleQueue{
		items:	make([]interface{}, capacity),
		capacity:	capacity,
		head:	uint(0),
		tail:	uint(0),
	}
}

func (this *CircleQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *CircleQueue) IsFull() bool {
	if this.head == uint(0) && this.tail == this.capacity - uint(1) {
		return true
	} else {
		if this.head - uint(1) == this.tail {
			return true
		}
	}
	return false
}

func (this *CircleQueue) Enqueue(v interface{}) bool{
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.items[this.head] = v
		this.tail = (this.head + uint(1)) % this.capacity
	} else{
		this.items[(this.tail) % this.capacity] = v
		this.tail = (this.tail + uint(1)) % this.capacity
	}
	return true
}

func (this *CircleQueue) Dequeue() interface{} {
	if this.IsEmpty() {
		return nil
	}

	item := this.items[this.head]
	this.head = (this.head + uint(1)) % this.capacity
	return item
}