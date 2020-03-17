package lru

import (
	"testing"
	"fmt"
)

func TestLru(t *testing.T) {
	lru := NewLRUCache(100)

	fmt.Println("初始化begin-----")
	fmt.Println(lru)
	fmt.Println("初始化end-----")

	for i := 0;i < 100; i++ {
		lru.Put(fmt.Sprintf("key%+v", i),i)
		// fmt.Println(lru)
	}

	for i := 0;i < 100; i++ {
		fmt.Println(lru.Get(fmt.Sprintf("key%+v", i)))
	}	
}