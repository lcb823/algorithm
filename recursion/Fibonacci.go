package recursion

import (
	// "fmt"
)

type Fibo struct {
	data map[int]int
}

func NewFibo(n int) *Fibo {
	return &Fibo{
		make(map[int]int, n),
	}
}

func (this *Fibo) Fibonacci(n int) int {
	if this.data[n] != 0 {
		return this.data[n]
	}

	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	result := this.Fibonacci(n - 1) + this.Fibonacci(n - 2)
	this.data[n] = result
	return result
}