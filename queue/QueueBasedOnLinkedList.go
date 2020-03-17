package queue

import (
	// "fmt"
	"../linkedlist"
)

type QueueBasedOnLinkedList struct {
	items	*singlylinkedlist.LinkedList
	tail	*singlylinkedlist.ListNode
}

func NewQueueBasedOnLinkedList() *QueueBasedOnLinkedList {
	return &QueueBasedOnLinkedList {
		items:	singlylinkedlist.NewLinkedList(),
		tail:	nil,
	}
}

func (this *QueueBasedOnLinkedList) IsEmpty() bool {
	return nil == this.items.Head.Next
}

func (this *QueueBasedOnLinkedList) IsFull() bool {
	return false
}

func (this *QueueBasedOnLinkedList) Enqueue(v interface{}) bool {
	//修改尾部节点的Next
	node := singlylinkedlist.NewListNode(v)
	if nil == this.tail {
		this.items.Head.Next = node
	} else {
		this.tail.Next = node
	}
	this.tail = node

	return true
}

func (this *QueueBasedOnLinkedList) Dequeue() interface{} {
	if nil == this.tail {
		return nil
	}
	node := this.items.Head.Next
	this.items.DeleteNode(node)
	if nil == this.items.Head.Next {
		this.tail = nil
	}
	return node.Value
}