package graphs

import (
	"fmt"
	"github.com/sircelsius/go-datastructures-algorithms/sort"
	"sync"
)

type BinarySearchTree struct {
	Root *BinaryTreeNode
	lock sync.RWMutex
}

func BinarySearchTreeFromSlice(slice []int) BinarySearchTree {
	sort.LomutoQuickSort(slice)
	var bst BinarySearchTree
	bst.Root = BinaryTreeNodeFromSlice(slice)
	return bst
}

func (tree *BinarySearchTree) Search(target int) bool {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return search(tree.Root, target)
}

func search(tree *BinaryTreeNode, target int) bool {
	if &tree.Value == nil {
		return false
	}

	if target == tree.Value {
		return true
	}

	if target < tree.Value {
		if tree.Left == nil {
			return false
		}
		return search(tree.Left, target)
	}

	if target > tree.Value {
		if tree.Right == nil {
			return false
		}
		return search(tree.Right, target)
	}
	return false
}

func IsBinaryTreeBinarySearchTree(tree *BinaryTreeNode) bool {
	if tree == nil || &tree.Value == nil {
		return true
	}

	if tree.Right == nil || tree.Left == nil {
		return true
	}

	if tree.Left.Value > tree.Value {
		return false
	}
	if tree.Right.Value < tree.Value {
		return false
	}
	if !IsBinaryTreeBinarySearchTree(tree.Left) || !IsBinaryTreeBinarySearchTree(tree.Right) {
		return false
	}

	return true
}

func (tree *BinarySearchTree) String() string {
	return fmt.Sprintf("%v", tree.Root)
}