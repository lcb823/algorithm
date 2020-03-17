package stack

import (
	"fmt"
)

type ArrayStack struct {
	data []interface{}
	top	uint //栈顶下一项. 空栈，top=0
}

var capacity uint = 100

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data:	make([]interface{}, 0, capacity),
		top: uint(0),//栈顶索引
	}
}

func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	value := this.data[this.top-1]
	this.top--
	return value
}

func (this *ArrayStack) Push(v interface{}) {
	//如果满了，返回false或者扩容
	if this.top >= uint(cap(this.data)) {
		return
	}
	if uint(len(this.data)) > this.top {
		//覆盖之前的值
		this.data[this.top] = v
	} else {
		//新加item
		this.data = append(this.data, v)
	}
	
	this.top++
}

func (this *ArrayStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top - 1]
}

func (this *ArrayStack) IsEmpty() bool {
	return this.top == uint(0)
}

func (this *ArrayStack) Flush() bool {
	this.top = 0
	return true
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		return
	}
	for i := uint(0); i < this.top; i++ {
		fmt.Println(this.data[i])
	}
}

