package heap

import (
	"testing"
	"math/rand"
	"time"
)

func compareFunc(v, toV interface{}) int {
	if v.(int) > toV.(int) {
		return 1
	} else if v.(int) == toV.(int) {
		return 0
	}
	return -1
}

func TestHeap(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := 20
	rand := r.Perm(num)

	ary := make([]interface{},num)
	for i := 0; i < num; i++ {
		ary[i] = rand[i]
	}

	heap := NewHeap(uint(num), compareFunc, "asc")

	for i := 0; i < num; i ++ {
		heap.InsertNode(ary[i])
		heap.Print()
	}

	heap.Sort()

	// heap.Print()
}

func TestHeapSort(t *testing.T) {

}