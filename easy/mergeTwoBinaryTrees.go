/*
You are given two binary trees root1 and root2.
Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. 
You need to merge the two trees into a new binary tree. 
The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. 
Otherwise, the NOT null node will be used as the node of the new tree.
Return the merged tree.
Note: The merging process must start from the root nodes of both trees.
*/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil return nil 

	if root1 == nil return root2

	if root2 == nil return root1

	resRoot := &TreeNode {
			Val: root1.Val + root2.Val,
			Left: mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
	}

	return resRoot

}
