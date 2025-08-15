package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func solve(preorder []int, inorder []int, l, h int) *TreeNode {
// 	nodeIndex := l
// 	leftIndex := -1

// 	return nil
// }

// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	return solve(preorder, inorder, 0, len(preorder)-1)
// }
