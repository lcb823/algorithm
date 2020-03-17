package singlylinkedlist

import (
	"fmt"
	// "math"
)

type ListNode struct {
	Next	*ListNode
	Value	interface{}
}

type LinkedList struct {
	Head	*ListNode
	Length	uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:	NewListNode(0),
		Length:	uint(0),
	}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfterNode(p *ListNode, v interface{}) bool {
	//遍历找到该节点
	cur := this.Head

	for nil != cur {
		if cur == p {
			//在p后面插入新节点
			newNode := NewListNode(v)
			newNode.Next = p.Next
			p.Next = newNode
			this.Length++
			return true
		}
		cur = cur.Next
	}
	return false
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBeforeNode(p *ListNode, v interface{}) bool {
	cur := this.Head.Next
	pre := this.Head

	if nil == cur {
		return false
	}
	for nil != cur {
		if cur == p {
			//在pre后面插入新节点
			newNode := NewListNode(v)
			newNode.Next = cur
			pre.Next = newNode
			this.Length++
			return true
		}
		pre = cur
		cur = cur.Next
	}
	return false
}

//在链表头部插入节点
func(this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfterNode(this.Head, v)
}
//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.Head
	for nil != cur {
		if nil == cur.Next {
			break
		}
		cur = cur.Next
	}
	return this.InsertAfterNode(cur, v)
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == this.Head.Next || nil == p {
		return false
	}
	cur := this.Head.Next
	pre := this.Head
	for nil != cur {
		if cur == p {
			//删除节点
			pre.Next = cur.Next
			p = nil
			this.Length--
			return true
		}
		pre = cur
		cur = cur.Next
	}
	return false
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.Length {
		return nil
	}
	cur := this.Head.Next
	for i := uint(0); i<=index; i++ {
		if i == index {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

//单链表反转
func (this *LinkedList) Reverse() *LinkedList {

	if this.Length <= uint(0) {
		return this
	}
	newList := NewLinkedList()
	cur := this.Head.Next
	for nil != cur {
		newList.InsertToHead(cur.Value)
		cur = cur.Next
	}
	return newList
}

//单链表反转,原地反转
func (this *LinkedList) Reverse1() {
	if this.Length == uint(0) || this.Length == uint(1) {
		return
	}
	var pre *ListNode = nil
	cur := this.Head.Next

	for nil != cur {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	this.Head.Next = pre
	// if nil == this.Head || nil == this.Head.Next || nil == this.Head.Next.Next {
	// 	return
	// }

	// var pre *ListNode = nil
	// cur := this.Head.Next
	// for nil != cur {
	// 	tmp := cur.Next
	// 	cur.Next = pre
	// 	pre = cur
	// 	cur = tmp
	// }

	// this.Head.Next = pre
}

//返回链表中间节点,如果有两个，则返回第二个
func (this *LinkedList) Middle() *ListNode {
	if this.Length == uint(0) {
		return nil
	}
	//计算中间节点的index
	var index uint
	if this.Length%2 == 0 {
		index = this.Length/2
	}else {
		index = this.Length/2
	}
	return this.FindByIndex(index)
}

//打印链表
func (this *LinkedList) Print() {
	var format string
	cur := this.Head.Next
	for nil != cur {
		format += fmt.Sprintf("|%+v", cur.Value)
		cur = cur.Next
	}
	fmt.Println(format)
}

//单链表回文实现,时间复杂度O(n),空间复杂度O(n)
func isPalindrome1(l *LinkedList) bool {
	lLen := l.Length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	step := lLen / 2
	cur := l.Head
	s := make([]string, 0, step)

	j := uint(0)
	for i := uint(0); i <= lLen - uint(1); i++ {
		cur = cur.Next

		//如果是前半部分，进入栈
		if i <= step - 1 {
			s = append(s, cur.Value.(string))
		} else if i == step && lLen % 2 == 1{//如果奇数节点，中间节点忽略
			continue
		} else {
			if s[step - uint(1) - j] != cur.Value.(string) {
				return false
			}
			j++
		}
	}
	return true
}

func isPalindrome2(l *LinkedList) bool {
	lLen := l.Length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	var isPa bool = true
	step := lLen / 2
	var pre *ListNode = nil
	cur := l.Head.Next

	for i := uint(0); i < step; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	var left, right *ListNode = pre, nil
	if lLen % 2 == 0 {
		right = cur
	} else {
		right = cur.Next
	}
	fmt.Println("----")
	fmt.Println(left)
	fmt.Println(right)
	fmt.Println("----")

	for nil != left && nil != right {
		if left.Value.(string) != right.Value.(string) {
			isPa = false
			break
		}
		left = left.Next
		right = right.Next
	}

	//还原
	tmp := cur
	cur = pre
	pre = tmp
	
	for i:= uint(0); i < step; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return isPa
}






