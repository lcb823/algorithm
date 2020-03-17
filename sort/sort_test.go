package sort

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestSort1(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := 50
	rand := r.Perm(len)

	ary := make([]interface{}, 0,len)
	for i := 0; i < len; i++ {
		ary = append(ary, rand[i])
	}

	fmt.Println(ary)

	// ary = Sort1(ary)
	// ary = Sort1(ary)
	// return

	ary1 := make([]interface{}, len)
	ary2 := make([]interface{}, len)
	ary3 := make([]interface{}, len)
	ary4 := make([]interface{}, len)
	ary5 := make([]interface{}, len)

	sortTimes := 1

	timeBegin := time.Now().Unix()
	for i := 0; i < sortTimes; i++{
		copy(ary1,ary)
		Sort1(ary1)
	}
	timeEnd := time.Now().Unix()
	fmt.Println(fmt.Sprintf("sort1 消耗时间 %v",timeEnd-timeBegin))

	timeBegin = time.Now().Unix()
	for i := 0; i < sortTimes; i++{
		copy(ary2,ary)
		ary = Sort2(ary2)	
	}
	timeEnd = time.Now().Unix()
	fmt.Println(fmt.Sprintf("sort2 消耗时间 %v",timeEnd-timeBegin))

	timeBegin = time.Now().Unix()
	for i := 0; i < sortTimes; i++{
		copy(ary3,ary)
		ary = Sort3(ary3)	
	}
	timeEnd = time.Now().Unix()
	fmt.Println(fmt.Sprintf("sort3 消耗时间 %v",timeEnd-timeBegin))


	timeBegin = time.Now().Unix()
	for i := 0; i < sortTimes; i++{
		copy(ary4,ary)
		ary = Sort4(ary4)	
	}
	timeEnd = time.Now().Unix()
	fmt.Println(fmt.Sprintf("sort4 消耗时间 %v",timeEnd-timeBegin))

	timeBegin = time.Now().Unix()
	for i := 0; i < sortTimes; i++{
		copy(ary5,ary)
		ary = Sort5(ary5)	
	}
	timeEnd = time.Now().Unix()
	fmt.Println(fmt.Sprintf("sort5 消耗时间 %v",timeEnd-timeBegin))
	// fmt.Println(ary)
}