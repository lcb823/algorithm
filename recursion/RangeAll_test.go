package recursion

import "testing"

func TestRangeALL(t *testing.T) {
	// slice1 := NewRangeArray(4)
	// for i:=0; i<4; i++{
	// 	slice1.data[i] = i+1
	// }
	// slice1.Range(0)

	slice2 := NewRangeArray(3)
	slice2.data[0] = "a"
	slice2.data[1] = "a"
	slice2.data[2] = "c"
	slice2.Range(0)

}