package main

import "strconv"
import "strings"


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
    idx := 0
    var val strings.Builder
	for {
		if idx >= len(S) || S[idx] == '-' {
            break
		}
		val.WriteByte(S[idx])
		idx += 1
	}
	stack := []*TreeNode {
        newNode(val.String()),
	}
    depth := 0
	for {
		if idx >= len(S) {
			break
		}
		if S[idx] == '-' {
			idx += 1
			depth += 1
			continue
		}
        val.Reset()
		for {
			if idx >= len(S) || S[idx] == '-' {
                break
			}
            val.WriteByte(S[idx])
			idx +=1
		}
		node := newNode(val.String())
		if len(stack) > depth {
			stack = stack[:depth]
		}
		if stack[depth-1].Left == nil {
			stack[depth-1].Left = node
		} else {
			stack[depth-1].Right = node
		}
		stack = append(stack, node)
        depth = 0
	}
	return stack[0]
}


func newNode(val string) *TreeNode {
	i, err := strconv.Atoi(val)
	if err != nil {
		i = -1
	}
	return &TreeNode{i, nil, nil}
}
