/*
You are given the root of a binary tree where each node has a value 0 or 1. 
Each root-to-leaf path represents a binary number starting with the most significant bit.

For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.
For all leaves in the tree, consider the numbers represented by the path from the root to that leaf. 
Return the sum of these numbers.
The test cases are generated so that the answer fits in a 32-bits integer.
*/

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, currSum int) int {
	currSum = (currSum << 1) | root.Val
	
	if root.Left == nil && root.Right == nil {
			return currSum
	}
	
	total := 0
	if root.Left != nil {
			total += dfs(root.Left, currSum)
	}
	
	if root.Right != nil {
			total += dfs(root.Right, currSum)
	}
	
	return total
}
