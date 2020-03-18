package singlylinkedlist

import (
	"testing"
	"fmt"
	"errors"
	"math/rand"
)

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.Head.Next))
	l.Print()

	t.Log(l.DeleteNode(l.Head.Next.Next))
	l.Print()
}

func TestMiddle(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 6; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	fmt.Println(l.Middle().Value)

	for i := 0; i < 1; i++ {
		l.InsertToTail(7)
	}
	l.Print()
	fmt.Println(l.Middle().Value)
}

func TestReverse(t *testing.T) {
	newList := NewLinkedList()
	ary := []int{1,2,3,4,7,32}
	cur := newList.Head
	for _, v := range ary {
		newList.InsertAfterNode(cur, v)
		cur = cur.Next
	}

	newList.Print()
	newList = newList.Reverse()
	newList.Print()
}

func TestReverse1(t *testing.T) {
	fmt.Println("test reverse1 begin")
	newList := NewLinkedList()
	ary := []int{1,2,3,4,7,32}
	cur := newList.Head
	for _, v := range ary {
		newList.InsertAfterNode(cur, v)
		cur = cur.Next
	}
	newList.Print()
	newList.Reverse1()
	newList.Print()
	newList.Reverse1()
	newList.Print()
	fmt.Println("test reverse1 end")
}

func TestIsPalindrome1(t *testing.T) {
	l := NewLinkedList()
	var isPa bool
	isPa = isPalindrome1(l)
	if isPa {
		t.Fatal(errors.New("is palindrome1 error when length==0"))
	}

	l = NewLinkedList()
	ary := []string{"a","b","c","d","e","f"}
	for _, v := range ary {
		l.InsertToTail(v)
	}
	isPa = isPalindrome1(l)
	if isPa {
		t.Fatal(errors.New("is palindrome1 error"))
	}

	l = NewLinkedList()
	ary = []string{"a1","b","c","d","c","b","a1"}
	for _, v := range ary {
		l.InsertToTail(v)
	}
	isPa = isPalindrome1(l)
	if !isPa {
		t.Fatal(errors.New("is palindrome1 error"))
	}
}

func TestIsPalindrome2(t *testing.T) {
	fmt.Println("test isPalindrome2 begin")
	l := NewLinkedList()
	ary := []string{"a","b","c","d","e","f","g"}
	for _, v := range ary {
		l.InsertToTail(v)
	}
	l.Print()
	isPa := isPalindrome2(l)
	if isPa {
		t.Fatal(errors.New("isPalindrome2 error"))
	}
	l.Print()

	l = NewLinkedList()
	ary = []string{"a","b","c","c","b","a"}
	for _, v := range ary {
		l.InsertToTail(v)
	}	
	isPa = isPalindrome2(l)
	l.Print()
	if !isPa {
		t.Fatal(errors.New("isPalindrome2 error when 偶数个节点 true"))
	}

	l = NewLinkedList()
	ary = []string{"a","b","c","d","c","b","a"}
	for _, v := range ary {
		l.InsertToTail(v)
	}	
	isPa = isPalindrome2(l)
	if !isPa {
		t.Fatal(errors.New("isPalindrome2 error when 奇数个节点 true"))
	}
	l.Print()

	fmt.Println("test isPalindrome2 end")
}

func TestHasRing(t *testing.T) {
	fmt.Println("test HasRing begin")
	//没有ring的情况
	l := NewLinkedList()
	ary := []string{"a","b","c","d","e","f","g"}
	for _, v := range ary {
		l.InsertToTail(v)
	}
	entry := l.HasRing()
	if entry != nil {
		t.Fatal("没有ring的情况，判断错误")
	}

	//有ring的情况
	ll := NewLinkedList()
	// var entry *ListNode;
	for i := 0; i < 100; i ++ {
		ll.InsertToTail(i)
	}

	v := rand.Intn(99)

	entry = ll.FindByIndex(uint(v))
	last := ll.FindByIndex(99)
	last.Next = entry

	findEntry := ll.HasRing()
	fmt.Println(findEntry)
	if entry != findEntry {
		t.Log(entry)
		t.Log(findEntry)

		t.Fatal("有ring的情况，判断错误")
	}

	fmt.Println("test HasRing end")
}
