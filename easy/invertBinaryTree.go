/*
Given the root of a binary tree, invert the tree, and return its root.
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil return root

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
