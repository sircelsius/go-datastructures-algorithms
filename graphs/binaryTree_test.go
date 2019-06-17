package graphs

import (
	"testing"
)

func TestBinaryTreeFromSlice(t *testing.T) {
	slice := []int{1,9,2,8,3,7,4,6,5}
	tree := BinaryTreeFromSlice(slice)
	if tree.root.Value != 3 {
		t.Errorf("Expected Value to have Value as center of array 3 but got %v", tree)
	}

	if tree.root.Depth() != 5 {
		t.Errorf("Expected Value to have depth 5 but got %v", tree)
	}

	if !tree.root.IsBalanced() {
		t.Errorf("Expected Value to be balanced")
	}
}

func TestBinaryTree_IsBalanced(t *testing.T) {
	var tree = BinaryTreeNode{
		1,
		nil,
		nil,
	}

	balancedOrFail(&tree, t, true)

	tree = BinaryTreeNode{
		1,
		&BinaryTreeNode{
			2,
			nil,
			nil,
		},
		nil,
	}

	balancedOrFail(&tree, t, true)

	tree = BinaryTreeNode{
		1,
		&BinaryTreeNode{
			2,
			nil,
			nil,
		},
		&BinaryTreeNode{
			3,
			nil,
			nil,
		},
	}

	balancedOrFail(&tree, t, true)

	tree = BinaryTreeNode{
		1,
		&BinaryTreeNode{
			2,
			&BinaryTreeNode{
				4,
				&BinaryTreeNode{
					5,
					nil,
					nil,
				},
				nil,
			},
			nil,
		},
		&BinaryTreeNode{
			3,
			nil,
			nil,
		},
	}

	balancedOrFail(&tree, t, false)
}

func balancedOrFail(tree *BinaryTreeNode, t *testing.T, shouldBeBalanced bool) {
	if tree.IsBalanced() && !shouldBeBalanced {
		t.Errorf("tree.IsBalanced=%v but got isBalanced=%v for tree %v", shouldBeBalanced, tree.IsBalanced(), tree)
	}

	if !tree.IsBalanced() && shouldBeBalanced {
		t.Errorf("tree.IsBalanced=%v but got isBalanced=%v for tree %v", shouldBeBalanced, tree.IsBalanced(), tree)
	}
}
