package stack

import (
	// "fmt"
	"testing"
	"errors"
)

func TestLinkListStack_Top(t *testing.T) {
	stack := NewLinkedListStack()
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

func TestLinkListStack_Pop(t *testing.T) {
	stack := NewLinkedListStack()
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

	stack.Print()
}


func TestLinkListStack_Flush(t *testing.T) {
	stack := NewLinkedListStack()
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