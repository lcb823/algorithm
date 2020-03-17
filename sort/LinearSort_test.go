package sort

import (
	"fmt"
	"testing"
	"math/rand"
)

func TestBuckSort(t *testing.T) {
	ary := make([]interface{}, 1000)
	for i := 0; i < 1000; i ++ {
		ary[i] = rand.Intn(10)
	}
	
	fmt.Println("TestBuckSort begin------")
	fmt.Println(ary)
	BuckSort(ary)
	fmt.Println(ary)
	fmt.Println("TestBuckSort end------")
}

func TestCountingSort(t *testing.T) {
	ary := make([]interface{}, 1000)
	for i := 0; i < 1000; i ++ {
		ary[i] = rand.Intn(10)
	}

	fmt.Println("TestCountingSort begin------")
	fmt.Println(ary)
	CountingSort(ary)
	fmt.Println(ary)
	fmt.Println("TestCountingSort end------")
}

