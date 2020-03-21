package stack

import (
	// "fmt"
	"testing"
	"errors"
)

func TestArrayStack_Top(t *testing.T) {
	stack := NewArrayStack()
	if !stack.IsEmpty() {
		t.Fatal(errors.New("check empty error"))
	}

	var url = "http://www.google.com/"
	stack.Push(url)
	if stack.IsEmpty() {
		t.Fatal(errors.New("check empty error"))
	}

	v := stack.Top()
	if v != url {
		t.Fatal(errors.New("top method error"))
	}

	v = stack.Top()
	if v != url {
		t.Fatal(errors.New("top method error"))
	}

	stack.Print()
}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack()
	var url1 = "http://www.google.com/"
	stack.Push(url1)

	var url2 = "https://www.bing.com/"
	stack.Push(url2)

	if stack.IsEmpty() {
		t.Fatal(errors.New("check empty error after push"))
	}

	v := stack.Pop()
	t.Log(v)
	if url2 != v {
		t.Fatal("pop error 1")
	}

	v = stack.Pop()
	t.Log(v)
	if url1 != v {
		t.Fatal("pop error 2")
	}

	if !stack.IsEmpty() {
		t.Fatal("check empty error after pop")
	}

	var url3 = "http://www.baidu33.com"
	stack.Push(url3)
	v = stack.Pop()
	if url3 != v {
		t.Fatal("pop error 3")
	}
	stack.Print()
}


func TestArrayStack_Flush(t *testing.T) {
	stack := NewArrayStack()
	ary := []interface{} {"google.com", "bing.com", "baidu.com"}
	for _, v := range ary {
		stack.Push(v)
	}
	stack.Print()
	stack.Flush()
	if !stack.IsEmpty() {
		t.Fatal("check empty error after flush")
	}
	stack.Print()
}