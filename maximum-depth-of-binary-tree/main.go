/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	deep := treeDepth(root)

	return deep
}

func treeDepth(node *TreeNode) int {
	var deepL, deepR int

	if node == nil {
		return 0
	}

	deepL++
	deepR++

	if node.Left != nil {
		deepL += treeDepth(node.Left)

		fmt.Println("m", node.Val, "l", node.Left.Val, deepL)
	}

	if node.Right != nil {
		deepR += treeDepth(node.Right)

		fmt.Println("m", node.Val, "r", node.Right.Val, deepR)
	}

	if deepL > deepR {
		return deepL
	} else {
		return deepR
	}
}