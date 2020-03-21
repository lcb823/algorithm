package recursion

import (
	// "fmt"
)	

type RangeArray struct {
	data	[]interface{}
	capacity uint
}

func NewRangeArray(capacity uint) *RangeArray {
	return &RangeArray{
		data:	make([]interface{}, capacity),
		capacity:	capacity,
	}
}

//递归法,参考  https://zhuanlan.zhihu.com/p/24889336
func (this *RangeArray) Range(start uint) [][]interface{}{
	var result [][]interface{}
	if start == this.capacity - 1  {
		tmp := make([]interface{}, this.capacity)
		copy(tmp, this.data)
		result = append(result, tmp)
		// fmt.Println(this.data)
		return result
	}

	for i := start; i < this.capacity; i++ {
		// 如果不同下标的元素相等，忽略
		if this.data[start] == this.data[i] && start != i {
			continue
		}
		this.data[start], this.data[i] = this.data[i], this.data[start]
		tmp := this.Range(start + uint(1))
		result = append(result, tmp...)
		this.data[start], this.data[i] = this.data[i], this.data[start]	
	}
	return result
}