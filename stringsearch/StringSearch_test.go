package stringsearch

import (
	"testing"
)

func TestIndexOfByBf(t *testing.T) {
	a := "abcdefgi"
	b := "fg"

	index := IndexOfByBf(a, b)
	t.Log(a)
	t.Log(b)
	t.Log(index)
	if index != 5 {
		t.Fatal("index error")
	}
}