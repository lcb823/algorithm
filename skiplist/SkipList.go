package skiplist


import (

)

type SkipListNode struct {
	v interface{}	//节点值
	score double//用于排序的分值
	level *SkipListLevel//每层的数据结构
	backward *SkipListNode//后一节点
}

type SkipListLevel struct {
	forward *SkipListNode//当前
	span uint//X.level[i].span 节点x在第i层到其下一个节点需跳过的节点数，相邻节点span=1
}

func NewSkipListLevel struct {
	forward:	NewSkipListNode()
}

func NewSkipListNode(v interface{}, level uint) *SkipListNode {
	return &SkipListNode{
		v: v,
		score: score,
		level:	nil,
		forwards: make([]SkipListNode, level)
	}
}

type SkipList struct {
	head	SkipListNode//头节点
	length	uint//跳表的长度
	level	uint//跳表的层数
}

func NewSkipList() *SkipList{
	head:	NewSkipListNode(nil, 1)
	level:	1
	length:	1
}

//插入新节点
func (this *SkipList) Insert(v interface{}) {

}

//删除节点
func (this *SkipList) Delete(v interface{}) {

}

func (this *SkipList) 