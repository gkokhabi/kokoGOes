/*
Given an integer array nums where the elements are sorted in ascending order, 
convert it to a height-balanced binary search tree.
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
			return nil
	}
	return &TreeNode{
			Val: nums[len(nums)/2],
			Left: sortedArrayToBST(nums[:len(nums)/2]),
			Right: sortedArrayToBST(nums[len(nums)/2 + 1:]),
	}
}

