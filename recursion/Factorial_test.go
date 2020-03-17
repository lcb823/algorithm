package recursion

import (
	"testing"
	"fmt"
)

func TestFactorial(t *testing.T) {
	n := 10
	fact := NewFact(n)
	fmt.Println("factorial begin")
	result := fact.Factorial(n)
	fmt.Println(result)
	fmt.Println("factorial end")
}