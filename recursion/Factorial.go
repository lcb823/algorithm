package recursion

import (

)

type Fact struct {
	data map[int]uint
}

func NewFact(n int) *Fact {
	return &Fact{
		data:	make(map[int]uint, n),
	}
}

func (this *Fact) Factorial(n int) uint {
	if n <= 0 {
		return uint(0)
	}
	if n == 1 {
		return uint(1)
	}
	if this.data[n] != uint(0) {
		return this.data[n]
	}

	result := uint(n)*this.Factorial(n - 1)
	this.data[n] = result
	return result
}