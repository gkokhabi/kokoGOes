/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil return true

	if left == nil || right == nil return false

	return left.Val == right.Val && helper(left.Right, right.Left) && helper(left.Left, right.Right)
	
}
