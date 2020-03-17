package sort

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestMergeSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := 50
	rand := r.Perm(len)

	ary := make([]interface{}, 0,len)
	for i := 0; i < len; i++ {
		ary = append(ary, rand[i])
	}


	fmt.Println("mergesort begin------------")
	fmt.Println(ary)
	MergeSort(ary)
	fmt.Println(ary)
	fmt.Println("mergesort end------------")
}