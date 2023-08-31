/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

    var inv func(*TreeNode)
    inv = func(node *TreeNode) {
        if node == nil {
            return
        }

        node.Right, node.Left = node.Left, node.Right

        inv(node.Left)
        inv(node.Right)
    }

    inv(root)
    return root
}


// Or with no helper
func invertTree2(root *TreeNode) *TreeNode {
    if root == nil || (root.Right == nil && root.Left == nil) {
        return root
    }

    root.Right, root.Left = root.Left, root.Right
    invertTree2(root.Right)
    invertTree2(root.Left)

    return root
}

