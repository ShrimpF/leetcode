package algo

// TreeNode has 2 nodes
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxDepth -- calculate maxium depth of Treenode
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(MaxDepth(root.Left), MaxDepth(root.Right))
}
