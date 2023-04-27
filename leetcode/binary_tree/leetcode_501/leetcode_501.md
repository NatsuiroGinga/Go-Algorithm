# LeetCode 501. Find Mode in Binary Search Tree

## 题目描述

给定具有重复项的二叉搜索树 (BST) 的根，返回其中的所有模式（即最常出现的元素）。

假设 BST 有如下定义：

- 结点左子树中所含结点的值小于等于当前结点的值
- 结点右子树中所含结点的值大于等于当前结点的值
- 左子树和右子树都是二叉搜索树

## 示例

```bash

例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2

返回[2].

提示：如果众数超过1个，不需考虑输出顺序

```

## 解题思路

### 方法一：中序遍历

二叉搜索树的中序遍历是一个非递减的序列，因此我们可以对二叉搜索树进行一次中序遍历，就可以得到有序的结果。

在中序遍历的过程中，我们需要记录当前的数以及当前的数的出现次数，对于二叉搜索树而言，相同的数一定出现在相邻的位置，因此我们可以在中序遍历的过程中更新当前数的出现次数，从而得到所有数出现的次数。

在中序遍历结束后，我们使用哈希表记录所有数出现的次数中的最大值，然后收集所有出现次数等于该最大值的数，即为所有出现次数最多的数。

### 方法二：Morris 中序遍历

在方法一中，我们对整棵树进行了一次中序遍历，然后对所有数出现的次数进行了一次遍历，因此总共进行了两次遍历。

我们也可以在中序遍历的过程中直接找出所有出现次数最多的数，这样可以只进行一次遍历。

我们可以用一种叫做 Morris 遍历的方法，来实现中序遍历。这种方法能将原本需要使用栈来解决的问题，转换为对空间复杂度的要求为 O(1) 的问题。

Morris 遍历的核心思想是利用树中大量的空指针，从而省去使用栈来进行中序遍历的过程。我们首先将当前节点current 初始化为根节点，然后对于当前节点current，如果其左子节点不为空，我们将其左子节点的最右侧节点的右子节点指向current，否则我们将其左子节点指向current。对于前者，我们在遍历完左子树之后，通过这个指向走回了current，因此我们可以将这个右子节点的右指针重新指向空。这样做的目的是为了将这个节点重新接回原来的树，并且不影响遍历的过程。而对于后者，我们在遍历完当前节点之后，通过这个指向走回了current，因此我们可以将左子节点指向current。这样做的目的是为了在遍历完当前节点之后，能够遍历其右子树。

在 Morris 遍历中，每个节点会被访问两次，因此总时间复杂度为 O(2n)=O(n)。

## 代码

### Java

```java
class Solution {
    public int[] findMode(TreeNode root) {
        List<Integer> list = new ArrayList<>();
        int[] res = new int[0];
        if (root == null) {
            return res;
        }
        TreeNode cur = root;
        TreeNode pre = null;
        int curCount = 1;
        int maxCount = 1;
        while (cur != null) {
            if (cur.left == null) {
                if (pre != null && pre.val == cur.val) {
                    curCount++;
                } else {
                    curCount = 1;
                }
                if (curCount == maxCount) {
                    list.add(cur.val);
                } else if (curCount > maxCount) {
                    maxCount = curCount;
                    list.clear();
                    list.add(cur.val);
                }
                pre = cur;
                cur = cur.right;
            } else {
                TreeNode node = cur.left;
                while (node.right != null && node.right != cur) {
                    node = node.right;
                }
                if (node.right == null) {
                    node.right = cur;
                    cur = cur.left;
                } else {
                    if (pre != null && pre.val == cur.val) {
                        curCount++;
                    } else {
                        curCount = 1;
                    }
                    if (curCount == maxCount) {
                        list.add(cur.val);
                    } else if (curCount > maxCount) {
                        maxCount = curCount;
                        list.clear();
                        list.add(cur.val);
                    }
                    pre = cur;
                    node.right = null;
                    cur = cur.right;
                }
            }
        }
        res = new int[list.size()];
        for (int i = 0; i < list.size(); i++) {
            res[i] = list.get(i);
        }
        return res;
    }
}
```