# LeetCode 222.Count Complete Tree Nodes

## 题目

Given the root of a complete binary tree, return the number of the nodes in the tree.

According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Design an algorithm that runs in less than O(n) time complexity.

## 翻译

给定一个完全二叉树的根节点，返回树中节点的数量。

根据维基百科，除了最后一层，每一层都是完全填充的完全二叉树，并且最后一层的所有节点都尽可能靠左。最后一层的节点数可以在1到2h之间，包括h。

设计一个时间复杂度小于O(n)的算法。

## 示例

```bash
输入：root = [1,2,3,4,5,6]
输出：6

输入：root = []
输出：0

输入：root = [1]
输出：1
```

## 解析

这道题目是一道二叉树的题目，题目要求是计算完全二叉树的节点数量，这里的完全二叉树是指除了最后一层，其他层都是满的，最后一层的节点都靠左的二叉树。

这道题目的解法有很多，这里介绍两种解法，一种是递归解法，一种是迭代解法。

### 递归解法

递归解法的思路是，先计算左子树的节点数量，再计算右子树的节点数量，最后加上根节点，就是整个树的节点数量。

递归解法的时间复杂度是O(n)，空间复杂度是O(h)，h是树的高度。

### 迭代解法

迭代解法的思路是，先计算树的高度，然后计算最后一层的节点数量，最后加上其他层的节点数量，就是整个树的节点数量。

迭代解法的时间复杂度是O(n)，空间复杂度是O(1)。

## 代码

### 递归解法

```python

class Solution:
    def countNodes(self, root: TreeNode) -> int:
        if not root:
            return 0
        return self.countNodes(root.left) + self.countNodes(root.right) + 1

```

### 迭代解法

```python

class Solution:
    def countNodes(self, root: TreeNode) -> int:
        if not root:
            return 0
        left_height = self.get_height(root.left)
        right_height = self.get_height(root.right)
        if left_height == right_height:
            return self.countNodes(root.right) + (1 << left_height)
        else:
            return self.countNodes(root.left) + (1 << right_height)

    def get_height(self, root: TreeNode) -> int:
        height = 0
        while root:
            height += 1
            root = root.left
        return height

```