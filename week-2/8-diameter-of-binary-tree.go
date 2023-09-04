/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMaxDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	maxDepth := findMaxDepth(node.Left, depth+1)

	if right := findMaxDepth(node.Right, depth+1); right > maxDepth {
		maxDepth = right
	}

	return maxDepth
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	var findMaxDepth func(node *TreeNode) int
	findMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := findMaxDepth(node.Left)
		right := findMaxDepth(node.Right)

		if left+right+1 > result {
			result = left + right + 1
		}

		if left > right {
			return left + 1
		}

		return right + 1
	}

	_ = findMaxDepth(root)
	return result - 1
}

// super slow way
// does double work
//func findMaxDepth(node *TreeNode, depth int) int {
//    if node == nil {
//        return depth
//    }
//
//    maxDepth := findMaxDepth(node.Left, depth+1)
//
//    if right := findMaxDepth(node.Right, depth+1); right > maxDepth {
//        maxDepth = right
//    }
//
//    return maxDepth
//}

//func diameterOfBinaryTree(root *TreeNode) int {
//    if root == nil {
//        return 0
//    }
//    onlyleft := diameterOfBinaryTree(root.Left)
//    onlyright := diameterOfBinaryTree(root.Right)
//    if onlyright > onlyleft {
//        onlyleft = onlyright
//    }
//    thruroot := findMaxDepth(root.Left, 0) + findMaxDepth(root.Right, 0)
//    if onlyleft > thruroot {
//        thruroot = onlyleft
//    }
//
//    return thruroot
//}

