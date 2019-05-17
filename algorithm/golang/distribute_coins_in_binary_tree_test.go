package main

import "testing"

func TestDistributeCoins(t *testing.T) {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{0, nil, nil},
		Right: &TreeNode{0, nil, nil},
	}
	result := distributeCoins(&root)
	if  result != 2 {
		t.Errorf("expected: %v, got: %v\n", 2, result)
	}
}
