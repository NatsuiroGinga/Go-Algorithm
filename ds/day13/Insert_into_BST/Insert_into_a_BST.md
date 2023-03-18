# LeetCode 701. Insert into a Binary Search Tree
## 题目描述
You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

## 思路
这道题目的思路很简单，就是找到要插入的位置，然后插入即可。这里有两种方法，一种是递归，一种是迭代。递归的方法比较简单，就是找到要插入的位置，然后插入即可。迭代的方法比较麻烦，需要找到要插入的位置，然后插入即可。

## 代码
### 递归
```java
class Solution {
    public TreeNode insertIntoBST(TreeNode root, int val) {
        if (root == null) {
            return new TreeNode(val);
        }
        if (root.val > val) {
            root.left = insertIntoBST(root.left, val);
        } else {
            root.right = insertIntoBST(root.right, val);
        }
        return root;
    }
}
```
### 迭代
```java
class Solution {
    public TreeNode insertIntoBST(TreeNode root, int val) {
        TreeNode node = root;
        while (node != null) {
            if (node.val > val) {
                if (node.left == null) {
                    node.left = new TreeNode(val);
                    break;
                } else {
                    node = node.left;
                }
            } else {
                if (node.right == null) {
                    node.right = new TreeNode(val);
                    break;
                } else {
                    node = node.right;
                }
            }
        }
        return root;
    }
}
```