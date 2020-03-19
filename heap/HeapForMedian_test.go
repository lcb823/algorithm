package heap

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestMedian(t *testing.T) {
	heap := NewHeapForMedian(100, compareFunc)
	for i := 0; i < 100; i ++ {
		fmt.Println(fmt.Sprintf("i=%+v", i))
		heap.AddItem(i)
		median := heap.FindMedian()
		fmt.Println(fmt.Sprintf("median=%+v", median))
	}


	//随机
	heap = NewHeapForMedian(10, compareFunc)
	var format string
	for i := 0; i < 10; i ++ {
		item := rand.Intn(100)
		fmt.Println(fmt.Sprintf("随机数:%+v", item))
		heap.AddItem(item)
		median := heap.FindMedian()
		fmt.Println(fmt.Sprintf("median=%+v", median))
		format += fmt.Sprintf(" %+v", item)
	}
	fmt.Println(format)
	heap.HeapSmall.Print()
	heap.HeapBig.Print()
	fmt.Println("------")
}