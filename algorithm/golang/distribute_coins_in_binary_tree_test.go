package main

import "testing"

func TestDistributeCoins(t *testing.T) {
	root := CoinTreeNode{
		Val: 3,
		Left: &CoinTreeNode{0, nil, nil},
		Right: &CoinTreeNode{0, nil, nil},
	}
	result := distributeCoins(&root)
	if  result != 2 {
		t.Errorf("expected: %v, got: %v\n", 2, result)
	}
}
