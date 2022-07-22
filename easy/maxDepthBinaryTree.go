/*
Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

func mx(x,y int) int {
	if x < y return y

	return x
}

func maxDepth(root *TreeNode) int {
	if(root == nil) {
			return 0
	}
	
	return 1+mx(maxDepth(root.Left),maxDepth(root.Right))
	
}
