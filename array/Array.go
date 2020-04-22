package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isArrayFull() bool {
	return this.Len() == uint(cap(this.data))
}

func (this *Array) setSize() error {
	capacity := uint(cap(this.data))
	capacity = 2 * capacity
	newData := make([]int, capacity)
	for i := uint(0); i < this.Len(); i++ {
		newData[i] = this.data[i]
	}
	this.data = newData
	return nil
}

func (this *Array) Insert(index uint, v int) error {
	// 判断array 是否full
	if this.isArrayFull() {
		// return errors.New("array is full")
		this.setSize()
	}

	if index > this.Len() {
		return errors.New("index is out of range")
	}

	for i := this.Len(); i >= index+1; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

func (this *Array) Delete(index uint) (int, error) {
	if index >= this.Len() {
		return 0, errors.New("index is out of range")
	}

	v := this.data[index]

	for i := index + 1; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

func (this *Array) Find(index uint) (int, error) {
	if index >= this.Len() {
		return 0, errors.New("index is out of range")
	}
	return this.data[index], nil
}

func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
