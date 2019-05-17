package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func distributeCoins(root *TreeNode) int {
	var result = 0;
	_dfs(root, &result)
	return result
}

func _dfs(root *TreeNode, result *int) int {
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
