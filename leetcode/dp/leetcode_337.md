# Leetcode 337. House Robber III

## Description

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

## 翻译

偷窃者发现了一个新的地方来偷窃。这个地方只有一个入口，称为“根”。除了根之外，每个房子都有一个且只有一个父房子。经过一番观察，聪明的小偷意识到“这个地方的所有房子都形成了一个二叉树”。如果两个直接链接的房子在同一晚上被打破，它将自动联系警察。

确定小偷今晚可以偷窃的最大金额，而不会警告警察。

**Example 1:**

    Input: [3,2,3,null,3,null,1]

         3
        / \
       2   3
        \   \ 
         3   1

    Output: 7 
    Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

**Example 2:**

    Input: [3,4,5,1,3,null,1]

         3
        / \
       4   5
      / \   \ 
     1   3   1

    Output: 9
    Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

## Solution

1. 树形DP，对于每个节点，有两种状态，偷和不偷，偷的话，那么就不能偷孩子节点，不偷的话，那么就可以偷孩子节点，然后取两种状态的最大值。
2. 设置一个长为2的`dp`数组，第一个元素表示偷当前节点能获得的最大收益，第二个元素表示不偷当前节点能获得的最大收益，然后**后序遍历**求解。
3. 由于递归求解，所以每次递归都会在栈中保存一个关于当前节点`dp`数组, 存储了当前节点的偷和不偷两种情况能获得的最大收益。
4. 最后, 返回根节点的两种情况的最大值。

**注意**

采用后序遍历的原因是，后序遍历是先遍历左右子树，然后再遍历根节点，这样就可以保证在遍历根节点的时候，左右子树的`dp`数组已经求解完毕，这样就可以直接使用左右子树的`dp`数组来求解当前节点的`dp`数组。

**Java**
```java
class Solution {
    public int rob(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int[] res = robHelper(root);
        // 返回偷和不偷两种情况的最大值
        return Math.max(res[0], res[1]);
    }

    private int[] robHelper(TreeNode root) {
        if (root == null) { 
            return new int[2];
        }
        //  递归求解左右子树
        int[] left = robHelper(root.left);
        int[] right = robHelper(root.right);
        int[] res = new int[2];
        // 偷当前节点，那么就不能偷孩子节点
        res[0] = root.val + left[1] + right[1];
        // 不偷当前节点，那么就可以偷孩子节点, 取两种情况的最大值
        res[1] = Math.max(left[0], left[1]) + Math.max(right[0], right[1]);
        return res;
    }
}
```
**cpp**
```cpp
class Solution {
public:
    int rob(TreeNode* root) {
        if (root == nullptr) {
            return 0;
        }
        vector<int> res = robHelper(root);
        return max(res[0], res[1]);
    }

    vector<int> robHelper(TreeNode* root) {
        if (root == nullptr) {
            return vector<int>(2, 0);
        }
        vector<int> left = robHelper(root->left);
        vector<int> right = robHelper(root->right);
        vector<int> res(2, 0);
        res[0] = root->val + left[1] + right[1];
        res[1] = max(left[0], left[1]) + max(right[0], right[1]);
        return res;
    }
};
```