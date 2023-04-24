# 中序遍历二叉搜索树(BST)及其应用(Go语言)

## 数据结构

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

## 1. 标准的中序遍历BST

Leecode94: [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)

```go
func inorderTraversal(root *TreeNode){
        var (
		stack []*TreeNode // 用于存储节点
		ans   []int       // 用于存储结果
	)

	if root == nil {
		return nil
	}

	for root != nil || len(stack) > 0 {
		for root != nil { // 将左子树全部入栈
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]   // 取出栈顶元素
		stack = stack[:len(stack)-1] // 删除栈顶元素

		ans = append(ans, root.Val) // 将栈顶元素的值添加到结果中
		root = root.Right           // 更新当前节点 为当前节点的右子树
	}

	return ans
}
```



## 2. 找到 BST 中第 K 个最小的元素

Leecode230: [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)

```go
func kthSmallest(root *TreeNode, k int) int {
    var stack []*TreeNode
	if root == nil {
		return 0
	}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if k == 1 {
			return root.Val
		}
		k--

		root = root.Right
	}

	return 0
}
```

## 3. BST验证问题

Leecode98 : [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)

```go
func isValidBST(root *TreeNode) bool {
    var (
		stack []*TreeNode
		pre   *TreeNode
	)

	if root == nil {
		return false
	}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		root = root.Right
	}

	return true
}
```
