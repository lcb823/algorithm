package queue

import (
	"testing"
	"fmt"
)

func TestQueueBasedOnArray(t *testing.T) {
	queue := NewQueueBasedOnArray(10)

	item := queue.Dequeue()

	if nil != item {
		t.Fatal("error dequeue when empty")
	}

	for i:=0; i<10; i++ {
		queue.Enqueue(i)
	}

	for i:=0; i<10; i++ {
		item := queue.Dequeue()
		fmt.Println(item)
		if item.(int) != i {
			t.Fatal("error dequeue")
		}
	}
}


func TestQueueBasedOnLinkedList(t *testing.T) {
	queue := NewQueueBasedOnLinkedList()

	item := queue.Dequeue()

	if nil != item {
		t.Fatal("error dequeue when empty")
	}

	for i:=0; i<10; i++ {
		queue.Enqueue(i)
	}
	queue.items.Print()

	for i:=0; i<10; i++ {
		item := queue.Dequeue()
		fmt.Println(item)
		if item.(int) != i {
			t.Fatal("error dequeue")
		}
	}
}

func TestCircleQueue(t *testing.T) {
	fmt.Println("------------begin")
	queue := NewCircleQueue(10)

	item := queue.Dequeue()

	if nil != item {
		t.Fatal("error dequeue when empty")
	}

	for i:=0; i<10; i++ {
		queue.Enqueue(i)
	}
	fmt.Println(queue.items)

	for i:=0; i<5; i++ {
		item := queue.Dequeue()
		if i == 9 {
			if nil != item {
				t.Fatal("error dequeue when nil")
			}
		} else {
			if item.(int) != i {
				t.Fatal("error dequeue")
			}
		}
	}

	for i:=0; i<10; i++ {
		queue.Enqueue(i)
	}
	fmt.Println(queue.items)


	fmt.Println("------------end")
}

