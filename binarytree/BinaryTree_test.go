package BinaryTree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	pre, in, post, layer []interface{}
}{

	{
		[]interface{}{1, 2, 3},
		[]interface{}{1, 3, 2},
		[]interface{}{3, 2, 1},
		[]interface{}{1, 2 ,3},
	},

	{
		[]interface{}{1, 2, 4, 5, 3, 6, 7},
		[]interface{}{4, 2, 5, 1, 6, 3, 7},
		[]interface{}{4, 5, 2, 6, 7, 3, 1},
		[]interface{}{1, 2, 3, 4, 5, 6, 7},
	},
	// 可以有多个 testCase
}

func PreIn2Tree(pre, in []interface{}) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		value: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.value, in)

	res.left = PreIn2Tree(pre[1:idx.(int)+1], in[:idx.(int)])
	res.right = PreIn2Tree(pre[idx.(int)+1:], in[idx.(int)+1:])

	return res
}

func indexOf(val interface{}, nums []interface{}) interface{} {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}

func Test_preOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)

		ast.Equal(tc.pre, NewBinaryTree(root).PreOrderTraversal(), "输入:%v", tc)
	}
}

func Test_inOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.in, NewBinaryTree(root).InOrderTraversal(), "输入:%v", tc)
	}
}

func Test_postOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.post, NewBinaryTree(root).PostOrderTraversal(), "输入:%v", tc)
	}
}

func Test_layerOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.layer, NewBinaryTree(root).LayerTraversal(), "输入:%v", tc)
	}
}