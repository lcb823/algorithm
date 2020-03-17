package array

import (
	"testing"
	"errors"
)

func TestInsert(t *testing.T) {
	capacity := 1
	arr := NewArray(uint(capacity))
	for i := 0; i < 5; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}

	arr.Print()

	err := arr.Insert(uint(6), 99)
	if nil == err {
		t.Fatal(errors.New("insert error when index>length"))
	}
	arr.Print()

	arr.Insert(uint(6), 99999)
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i :=0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if nil != err {
			t.Fatal(err.Error())
		}
		arr.Print()
	}
}

func TestFind(t *testing.T) {
	arr := NewArray(uint(10))
	ary := []int{1,2,3,4,5}
	for k, v := range ary {
		arr.Insert(uint(k), v)
	}

	for i := 0;i <= 4; i++{
		val, _ := arr.Find(uint(i))
		if  val != ary[i] {
			t.Fatal(errors.New("find error"))
		}
	}
}