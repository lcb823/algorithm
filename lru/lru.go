package lru

import (
	"hash/fnv"
	"fmt"
)

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 100
)

type lruNode struct {
	prev *lruNode
	next *lruNode

	key string
	value interface {}

	hnext *lruNode
}

type LRUCache struct {
	node []lruNode
	head *lruNode
	tail *lruNode

	capacity uint
	used uint
}

func NewLRUCache(capacity uint) *LRUCache {
	return &LRUCache{
		node: make([]lruNode, capacity),
		head: nil,
		tail: nil,

		capacity: capacity,
		used: uint(0),
	}
}

//添加缓存
func (this *LRUCache) Put(key string,value interface{}) {
	tmp := this.searchNode(key)
	if tmp != nil {
		tmp.value = value
		this.moveToTail(tmp)
		return
	}
	this.addNode(key, value)

	if this.used > this.capacity {
		this.delNode()
	}
}

//查找缓存
func (this *LRUCache) Get(key string) interface{} {
	if this.tail == nil {//空
		return nil
	}
	tmp := this.searchNode(key)
	if tmp != nil {
		this.moveToTail(tmp)
		return tmp.value
	}
	return -1
}

//添加新节点
func (this *LRUCache) addNode(key string, value interface{}) {
	newNode := &lruNode{
		key: key,
		value: value,
	}

	tmp := &this.node[this.hash(key)]
	fmt.Println(tmp)
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	this.used++

	if this.tail == nil {
		this.head, this.tail = newNode, newNode
	}

	this.tail.next = newNode
	newNode.prev = this.tail
	this.tail = newNode
}

//查找节点
func (this *LRUCache) searchNode(key string) *lruNode{
	if this.tail == nil {
		return nil
	}
	node := this.node[this.hash(key)].hnext
	for node != nil {
		if node.key == key {
			return node
		}
		node = node.hnext
	}
	return nil
}

//删除头节点 包括 链表中的和散列表中的
func (this *LRUCache) delNode() {
	if this.head == nil {
		return
	}
	prev := &this.node[this.hash(this.head.key)]
	tmp := prev.hnext
	for tmp != nil && tmp.key != this.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if nil == tmp {
		return
	}
	prev.hnext = tmp.hnext
	this.head = this.head.next
	this.head.prev = nil
	this.used--
}

//移动节点到队列尾部
func (this *LRUCache) moveToTail(node *lruNode) {
	if this.tail == node {
		return
	}
	if this.head == node {
		this.head = node.next
		this.head.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.next = nil
	this.tail.next = node
	node.prev = this.tail

	this.tail = node
}


func (this *LRUCache) hash(key string) uint32 {
	h := fnv.New32a()
    h.Write([]byte(key))
    k := h.Sum32()


    var result uint32
    if hostbit {
		result = (k ^ (k >> 32)) & (LENGTH - 1)
	} else {
		result = (k ^ (k >> 16)) & (LENGTH - 1)
	}

	fmt.Println(fmt.Sprintf("hash-value for %s is %v", key, result))
	return result
}