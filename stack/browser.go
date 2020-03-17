package stack

import (
	"fmt"
)

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
	Top() interface{}
	Flush() bool
}

type Browser struct {
	forwardStack Stack
	backStack Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack:	NewArrayStack(),
		backStack:	NewLinkedListStack(),
	}
}

func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) CanBack() bool {
	curUrl := this.backStack.Pop()
	var canBack bool = true
	if this.backStack.IsEmpty() {
		canBack = false
	}
	this.PushBack(curUrl)
	return canBack
}

func (this *Browser) Open(newUrl interface{}) {
	fmt.Println(fmt.Sprintf("open new url:%+v", newUrl))
	this.forwardStack.Flush()
	this.PushBack(newUrl)
}

func (this *Browser) PushBack(url interface{}) {
	this.backStack.Push(url)
}

func (this *Browser) Forward() {
	if this.CanForward() {
		top := this.forwardStack.Pop()
		this.backStack.Push(top)
		fmt.Println(fmt.Sprintf("forward to %+v", top))
	}
}

func (this *Browser) Back() {
	if this.CanBack() {
		top := this.backStack.Pop()
		this.forwardStack.Push(top)
		newTop := this.backStack.Top()
		fmt.Println(fmt.Sprintf("back to %+v", newTop))
	}
}