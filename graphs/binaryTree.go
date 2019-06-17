package graphs

import (
	"fmt"
	"math"
	"sync"
)

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
	lock sync.RWMutex
}

func BinaryTreeFromSlice(slice []int) *BinaryTree {
	var bt BinaryTree
	bt.root = BinaryTreeNodeFromSlice(slice)
	return &bt
}

func BinaryTreeNodeFromSlice(slice []int) *BinaryTreeNode {
	if &slice == nil || len(slice) == 0 {
		return nil
	}

	if len(slice) == 1 {
		return &BinaryTreeNode{
			slice[0],
			nil,
			nil,
		}
	}

	var centerIndex = len(slice) / 2
	root := &BinaryTreeNode{
		slice[centerIndex],
		nil,
		nil,
	}
	root.Left = BinaryTreeNodeFromSlice(slice[0 : centerIndex])
	root.Right = BinaryTreeNodeFromSlice(slice[centerIndex:])
	return root
}

func (tree *BinaryTreeNode) hasChildren() bool {
	if &tree.Value == nil {
		return false
	}
	if tree.Left == nil && tree.Right == nil {
		return false
	}

	return true
}

func (tree *BinaryTreeNode) IsBalanced() bool {
	var leftDepth = GetDepth(tree.Left)
	var rightDepth = GetDepth(tree.Right)

	if math.Abs(float64(rightDepth-leftDepth)) > 1 {
		return false
	}

	if tree.Left != nil && !tree.Left.IsBalanced() {
		return false
	}

	if tree.Right != nil && !tree.Right.IsBalanced() {
		return false
	}

	return true
}

func (tree *BinaryTreeNode) Depth() int {
	if &tree == nil {
		return 0
	}

	var leftDepth = 0
	if tree.Left != nil {
		leftDepth = tree.Left.Depth()
	}

	var rightDepth = 0
	if tree.Right != nil {
		rightDepth = tree.Right.Depth()
	}
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func GetDepth(tree *BinaryTreeNode) int {
	if tree == nil {
		return 0
	}
	return tree.Depth()
}

func (tree *BinaryTree) String() string {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return tree.root.String()
}

func (node *BinaryTreeNode) String() string {
	if node == nil {
		return  "nil"
	}
	return fmt.Sprintf("[%v<-(%v)->%v]", node.Left, node.Value, node.Right)
}

