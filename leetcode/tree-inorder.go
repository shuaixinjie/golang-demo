package leetcode

// 给定一个二叉树，返回其中序遍历，中序遍历的概念不赘述了
/**
	输入：root = [1,null,2,3]，传入的不是这样的数组，而是用数组表达树的形式
	输出：[1,3,2]
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal 注意写法，并不唯一
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}