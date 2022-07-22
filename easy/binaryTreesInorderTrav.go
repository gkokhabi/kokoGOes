/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

func inorderTraversal(root *TreeNode) []int {
	if root == nil return []int{}

	l := inorderTraversal(root.Left)
	r := inorderTraversal(root.Right)
	l = append(l, root.Val)
	l = append(l, r...)
	return l
}
