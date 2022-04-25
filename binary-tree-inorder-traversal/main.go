/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	inorder(root, &res)
	// preorder(root, &res)
	// postorder(root, &res)

	return res
}

func preorder(node *TreeNode, res *[]int) {
	if node != nil {
		*res = append(*res, node.Val)
		inorder(node.Left, res)
		inorder(node.Right, res)
	}
}

func inorder(node *TreeNode, res *[]int) {
	if node != nil {
		inorder(node.Left, res)
		*res = append(*res, node.Val)
		inorder(node.Right, res)
	}
}

func postorder(node *TreeNode, res *[]int) {
	if node != nil {
		inorder(node.Left, res)
		inorder(node.Right, res)
		*res = append(*res, node.Val)
	}
}