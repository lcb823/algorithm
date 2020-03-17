package recursion

import (
	"fmt"
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
func (this *RangeArray) Range(start uint) {
	if start == this.capacity - 1  {
		fmt.Println(this.data)
		return
	}

	for i := start; i < this.capacity; i++ {
		// 如果不同下标的元素相等，忽略
		if this.data[start] == this.data[i] && start != i {
			continue
		}
		this.data[start], this.data[i] = this.data[i], this.data[start]
		this.Range(start + uint(1))
		this.data[start], this.data[i] = this.data[i], this.data[start]	
	}
}

//由小到大排序，依次寻找相邻且更大的数
func (this *RangeArray) Range1(){
	
}