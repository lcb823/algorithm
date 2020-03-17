package stack

import (
	"../linkedlist"
)

type LinkedListStack struct {
	Data	*singlylinkedlist.LinkedList
}

func NewLinkedListStack() *LinkedListStack{
	return &LinkedListStack{
		Data:	singlylinkedlist.NewLinkedList(),
	}
}

func (this *LinkedListStack) IsEmpty() bool {
	if nil == this.Data.Head.Next {
		return true
	}
	return false
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	next := this.Data.Head.Next
	v := next.Value
	this.Data.DeleteNode(next)
	return v
}

func (this *LinkedListStack) Push(v interface{}) {
	this.Data.InsertToHead(v)
}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data.Head.Next.Value
}

func (this *LinkedListStack) Flush() bool {
	this.Data.Head.Next = nil
	return true
}

func (this *LinkedListStack) Print() {
	this.Data.Print()
}