package leetcode_106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var record map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	record = map[int]int{}
	for i, v := range inorder {
		record[v] = i
	}

	return recur(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func recur(inorder []int, inLow, inHigh int, postorder []int, postLow, postHigh int) *TreeNode {
	if inLow > inHigh || postLow > postHigh {
		return nil
	}
	root := &TreeNode{Val: postorder[postHigh]}
	index := record[root.Val]

	root.Left = recur(inorder, inLow, index-1, postorder, postLow, postLow+index-inLow-1)
	root.Right = recur(inorder, index+1, inHigh, postorder, postLow+index-inLow, postHigh-1)

	return root
}
