package stringsearch

import (
	"testing"
)

func TestIndexOfByBF(t *testing.T) {
	a := "abcdefgi"
	b := "fg"

	index := IndexOfByBF(a, b)
	t.Log(a)
	t.Log(b)
	t.Log(index)
	if index != 5 {
		t.Fatal("index error")
	}
}

func TestIndexOfByRK(t *testing.T) {
	a := "abcdefgi"
	b := "fg"

	index := IndexOfByRK(a, b)
	t.Log(a)
	t.Log(b)
	t.Log(index)
	if index != 5 {
		t.Fatal("index error")
	}
}

func TestFindAllByBM(t *testing.T) {
	a := "abcdefgi"
	b := "fg"

	index, _ := FindAllByBM(a, b)
	if index[0] != 5 || len(index) != 1 {
		t.Log(a)
		t.Log(b)
		t.Log(index)
		t.Fatal("index error")
	}

	a = "GTTATAGCTGATCGCGGCGTAGCGGCGAAGTAGCGGCGTTD"
	b = "GTAGCGGCG"

	index, _ = FindAllByBM(a, b)
	if len(index) != 2 || index[0] != 18 || index[1] != 29 {
		t.Log(a)
		t.Log(b)
		t.Log(index)
		t.Fatal("index error")
	}
}