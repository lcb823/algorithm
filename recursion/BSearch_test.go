package recursion

import (
	"testing"
	"fmt"
)

func TestBSearch(t *testing.T) {
	slice := make([]interface{}, 100)
	for i := 0; i < 100; i ++ {
		slice[i] = i + 10
	}

	index := BSearch(slice, 43)
	fmt.Println(index)
}