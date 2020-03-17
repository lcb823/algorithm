package BinaryTree

import (
	"fmt"
)

type BinarySearchTree struct {
	Root *TreeNode
	CompareFun func(v, nodeV interface{}) int
}

func NewBinarySearchTree(v interface{}, compareFun func(v, nodeV interface{}) int) *BinarySearchTree {
	rootNode := NewTreeNode(v)
	return &BinarySearchTree{
		Root: rootNode,
		CompareFun: compareFun,
	}
}

func (this *BinarySearchTree) isLeaf(node *TreeNode) bool {
	if node.left == nil && node.right == nil {
		return true
	}
	return false
}

//查找某个值,返回全部相同的节点
func (this *BinarySearchTree) Find(v interface{}) []interface{} {
	root := this.Root
	if nil == root {
		return nil
	}
	node := root
	var result []interface{}
	for node != nil {
		if this.CompareFun(v, node.value) == 0 {
			result = append(result, node.value)
			node = node.right
		} else if this.CompareFun(v, node.value) < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}

	return result
}

//插入节点
func (this *BinarySearchTree) Insert(v interface{}) {
	node := this.Root
	for nil != node {
		compareResult := this.CompareFun(v, node.value)
		if compareResult >= 0 {
			if nil == node.right {
				node.right = NewTreeNode(v)
				break
			}
			node = node.right
		} else {
			if nil == node.left {
				node.left = NewTreeNode(v)
				break
			}
			node = node.left
		}
	}
}

//删除某个节点
func (this *BinarySearchTree) DeleteNode(nodeToDel *TreeNode) bool{
	root := this.Root
	if nil == root {
		return false
	}
	// if root == nodeToDel {
	// 	this.Root = nil
	// 	return true
	// }
	node := root
	var preNode *TreeNode
	var isRight bool
	for nil != node {
		// 判断当前节点是否等于 nodeToDel
		if node == nodeToDel {
			//删除
			if this.isLeaf(node) {
				if preNode != nil {	
					if isRight {
						preNode.right = nil
					} else {
						preNode.left = nil
					}
				} else {
						this.Root = nil
				}
			} else if node.left != nil && node.right != nil {
				//如果当前节点存在左右节点,找到右子树的最小节点smallestNode，将smallestNode的value替换给nodeToDel,然后删除
				rightNode := node.right
				var smallestNode *TreeNode
				for rightNode != nil {
					smallestNode = rightNode
					if rightNode.left != nil {
						rightNode = rightNode.left
					} else {
						rightNode = rightNode.right
					}
				}
				//更新待删除节点的value
				node.value = smallestNode.value
				//删除右子树最小节点
				if isRight {
					preNode.right = nil
				} else {
					preNode.left = nil
				}
			} else {
				//找到待删除节点的下一个节点
				var nextNode *TreeNode
				if node.left != nil {
					nextNode = node.left
				} else{
					nextNode = node.right
				}
				//将待删除节点的父节点的指针设为nextNode
				if nil != preNode {
					if isRight {
						preNode.right = nextNode
					} else {
						preNode.left = nextNode
					}
				} else {
					this.Root = nextNode
				}
			}
			return true
		}
		compareV := this.CompareFun(nodeToDel.value, node.value)
		if compareV >= 0 {
			preNode = node
			node = node.right
			isRight = true
		} else {
			preNode = node
			node = node.left
			isRight = false
		}
	}
	return false
}

//删除符合条件的所有节点
func (this *BinarySearchTree) DeleteNodeByV(v interface{}) {
	root := this.Root
	if nil == root {
		return
	}
	node := root
	var nodesToDel []*TreeNode
	for node != nil {
		if this.CompareFun(v, node.value) == 0 {
			nodesToDel = append(nodesToDel, node)
			node = node.right
		} else if this.CompareFun(v, node.value) < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	
	for _, item := range nodesToDel {
		this.DeleteNode(item)
	}
	fmt.Println(nodesToDel)
}