package main

/**
 * Definition for a binary tree node.
 */
type CoinTreeNode struct {
    Val int
    Left *CoinTreeNode
    Right *CoinTreeNode
}


func distributeCoins(root *CoinTreeNode) int {
	var result = 0;
	_dfs(root, &result)
	return result
}

func _dfs(root *CoinTreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := _dfs(root.Left, result)
	right := _dfs(root.Right, result)
	coins := root.Val + left + right - 1

	if coins < 0 {
		*result += -coins
	} else {
		*result += coins
	}
	return coins
}
