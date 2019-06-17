package graphs

import "testing"

func TestBinarySearchTreeFromSlice(t *testing.T) {
	var slice = []int{9,8,2,4,3,5,3,7,8}
	var tree = BinarySearchTreeFromSlice(slice)
	testBinarySearchTreeProperty(tree.Root, t)
}

func TestIsBinaryTreeBinarySearchTree(t *testing.T) {
	tree := BinaryTreeNodeFromSlice([]int{1,4,2,8,5,9})
	if IsBinaryTreeBinarySearchTree(tree) {
		t.Errorf("Tree %v is not a binary search tree!", tree)
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	var slice = []int{9,8,2,4,3,5,3,7,8}
	var tree = BinarySearchTreeFromSlice(slice)
	if !tree.Search(2) {
		t.Errorf("Unable to find value 2 in tree %v", tree)
	}
}

func testBinarySearchTreeProperty(tree *BinaryTreeNode, t *testing.T) {
	if tree != nil {
		if &tree.Value != nil {
			if tree.Left != nil && &tree.Left.Value != nil && tree.Left.Value > tree.Value {
				t.Errorf("Left Root %v greater than root %v", tree.Left.Value, tree.Value)
			}
			if tree.Right != nil && &tree.Right.Value != nil&& tree.Right.Value < tree.Value {
				t.Errorf("Right Root %v lower than root %v", tree.Right.Value, tree.Value)
			}

			testBinarySearchTreeProperty(tree.Left, t)
			testBinarySearchTreeProperty(tree.Right, t)
		}
	}
}
