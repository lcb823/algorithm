package recursion

import (
	"testing"
	"fmt"
)

func TestFibonacci(t *testing.T) {
	n := 50
	fibo := NewFibo(n)
	result := fibo.Fibonacci(n)
	fmt.Println(result)
}