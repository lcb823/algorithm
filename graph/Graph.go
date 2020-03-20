package graph

import (
	"../linkedlist"
	"../queue"
	"fmt"
)


//无向图
type Graph struct {
	adj []*singlylinkedlist.LinkedList//此处用自己实现的单链表,也可以用 container/list
	v int //节点个数
}

func NewGraph(v int) *Graph {
	graph := &Graph{
		adj: make([]*singlylinkedlist.LinkedList, v),
		v: v,
	}

	for i := 0; i < v; i ++ {
		graph.adj[i] = singlylinkedlist.NewLinkedList()
	}

	return graph
}


//添加边
func (this *Graph) AddEdge(s, t int) {
	this.adj[s].InsertToTail(t)
	this.adj[t].InsertToTail(s)
}

//BFS,广度优先搜索节点s 到节点t的路径
func (this *Graph) BFS(s, t int) []int{
	// 维护变量 visited:存放节点是否访问过 prev:存放节点的前驱节点 queue:队列，存放访问节点
	visited := make([]int, this.v)
	prev := make([]int, this.v)
	for _k, _ := range prev {
		 prev[_k] = -1
	}
	queue := queue.NewCircleQueue(uint(this.v))

	queue.Enqueue(s)
	visited[s] = 1

	isFound := false
	for !isFound {
		if queue.IsEmpty() {
			break
		}
		itemIndex := queue.Dequeue().(int)
		list := this.adj[itemIndex]
		next := list.Head.Next

		for ; next != nil; next = next.Next {
			if visited[next.Value.(int)] == 0 {//未访问
				queue.Enqueue(next.Value)
				visited[next.Value.(int)] = 1
				prev[next.Value.(int)] = itemIndex
				//判断是否要找的t
				if next.Value.(int) == t {
					isFound = true
					break
				}
			}
		}
	}

	//打印路径
	if !isFound {
		return nil
	}

	fmt.Println(prev)
	return this.findPath(prev, s, t)
}

//输出从s到t的路径（倒序）
func (this *Graph) findPath(prev []int, s, t int) []int{
	// fmt.Println(fmt.Sprintf("s=%+v", s))
	// fmt.Println(fmt.Sprintf("t=%+v", t))
	var path []int
	path = append(path, t)
	if s == t || prev[t] == -1 {
		return path
	} else {
		tmp := this.findPath(prev, s ,prev[t])
		path = append(path, tmp...)
		return path
	}
	
	return path
}

//DFS,深度优先搜索节点s 到节点t的路径
func (this *Graph) DFS(s, t int) []int {
	visited := make([]int, this.v)
	prev := make([]int, this.v)
	for _k, _ := range prev {
		 prev[_k] = -1
	}

	isFound := false
	this.dfsRecurse(prev, visited, s, t, isFound)

	return this.findPath(prev, s, t)
}

func (this *Graph) dfsRecurse(prev, visited []int, s, t int, isFound bool) {
	if isFound {
		return
	}

	visited[s] = 1
	if s == t {
		isFound = true
		return
	}

	list := this.adj[s]
	next := list.Head.Next

	for ; next != nil; next = next.Next {
		if visited[next.Value.(int)] == 0{
			prev[next.Value.(int)] = s
			this.dfsRecurse(prev, visited, next.Value.(int), t, false)
		}
	}

}