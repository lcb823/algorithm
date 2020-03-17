package sort

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestQuickSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := 50
	rand := r.Perm(len)

	ary := make([]interface{}, 0,len)
	for i := 0; i < len; i++ {
		ary = append(ary, rand[i])
	}

	fmt.Println("quicksort begin------------")
	fmt.Println(ary)
	QuickSort(ary)
	fmt.Println(ary)
	fmt.Println("quicksort end------------")
}

func TestFindK(t *testing.T){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := 50
	rand := r.Perm(len)

	ary := make([]interface{}, 0,len)
	for i := 0; i < len; i++ {
		ary = append(ary, rand[i])
	}

	fmt.Println("findK begin------------")
	v := FindK(ary, 4)
	fmt.Println(v)
	fmt.Println("findK end------------")	
}