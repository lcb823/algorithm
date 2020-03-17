package BinaryTree

import (
	// "fmt"
)

type TreeNode struct {
	value interface{}
	left *TreeNode
	right *TreeNode
}

func NewTreeNode(v interface{}) *TreeNode {
	return &TreeNode{v, nil, nil}
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree(root *TreeNode) *BinaryTree {
	return &BinaryTree{root}
}

//前序遍历
func (this *BinaryTree) PreOrderTraversal() []interface{}{
	root := this.Root
	if nil == root {
		return nil
	}
	if nil == root.left && nil == root.right {
		return []interface{}{root.value}
	}

	var stack []*TreeNode
	var result []interface{}

	stack = append(stack, root)
	for len(stack) != 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, item.value)
		if item.right != nil {
			stack = append(stack, item.right)
		}
		if item.left != nil {
			stack = append(stack, item.left)
		}
	}

	return result
}

//中序遍历
func (this *BinaryTree) InOrderTraversal() []interface{}{
	root := this.Root
	if nil == root {
		return nil
	}
	if nil == root.left && nil == root.right {
		return []interface{}{root.value}
	}

	leftTree := NewBinaryTree(root.left)
	rightTree := NewBinaryTree(root.right)

	result := leftTree.InOrderTraversal()
	result = append(result, root.value)
	result = append(result, rightTree.InOrderTraversal()...)
	return result
}

//后序遍历
func (this *BinaryTree) PostOrderTraversal() []interface{} {
	root := this.Root
	if nil == root {
		return nil
	}
	if nil == root.left && nil == root.right {
		return []interface{}{root.value}
	}

	leftTree := NewBinaryTree(root.left)
	rightTree := NewBinaryTree(root.right)

	result := leftTree.PostOrderTraversal()
	result = append(result, rightTree.PostOrderTraversal()...)
	result = append(result, root.value)
	return result
}

//按层遍历
func (this *BinaryTree) LayerTraversal() []interface{}{
	root := this.Root
	if nil == root {
		return nil
	}
	if nil == root.left && nil == root.right {
		return []interface{}{root.value}
	}

	nodes := this.layerTraversal([]*TreeNode{root})
	results := nodes
	for len(nodes) >0 {
		nodes = this.layerTraversal(nodes)
		results = append(results, nodes...)
	}
	
	res := []interface{}{root.value}
	for _, v := range results {
		res = append(res, v.value)
	}
	return res
}

func (this *BinaryTree) layerTraversal(parents []*TreeNode) []*TreeNode{
	if len(parents) == 0 {
		return nil
	}
	result := []*TreeNode{}
	for _, v := range parents {
		if nil != v.left {
			result = append(result, v.left)
		}
		if nil != v.right {
			result = append(result, v.right)
		}
	}
	return result
}