package main

import "testing"

func TestRecoverFromPreorder(t *testing.T)  {
	node := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{
			5,
			&TreeNode{6, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	cases := []struct{
		input string;
		expected *TreeNode;
	} {
		{"1-2--3--4-5--6--7", node},
	}
	for _, c := range cases {
		result := recoverFromPreorder(c.input)
		if isEqual(result, c.expected){
			t.Errorf("case: '%v', got: '%v'\n", c, result)
		}
	}
}

func isEqual(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 != nil && t2 != nil {
		return t1.Val != t2.Val && isEqual(t1.Left, t2.Left) && isEqual(t1.Right, t2.Right)
	}
	return t1 == nil && t2 == nil
}
