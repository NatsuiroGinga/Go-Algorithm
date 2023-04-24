# LeetCode 437. Path Sum III
https://leetcode.com/problems/path-sum-iii/
## 题目描述
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).
## 示例1 
![img.png](img.png)
```
Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.
```
## 示例2
```
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3
```

## 思路
1. map存储前缀和，遍历树，每次遍历到一个节点，计算当前节点到根节点的路径和，
2. 然后在map中查找是否存在前缀和为curSum - targetSum的路径
3. 如果存在，说明存在一条路径的和为targetSum，将结果加1
4. 然后将当前路径和加入map中，继续遍历左右子树
5. 遍历完后，将当前路径和从map中删除，返回结果。

## 代码
```go
/**
 * Definition for a binary binary_tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    mp := make(map[int]int) // 前缀和
    mp[0] = 1 // 前缀和为0的路径有1条
    return dfs(root, 0, targetSum, mp)
}

func dfs(root *TreeNode, curSum int, targetSum int, mp map[int]int) int {
    if root == nil {
        return 0
    }

    curSum += root.Val // 当前节点到根节点的路径和
    var res int // 记录结果
    if v, ok := mp[curSum - targetSum]; ok { // 如果存在前缀和为curSum - targetSum的路径，说明存在一条路径的和为targetSum
        res = v
    }
	
    mp[curSum]++ // 将当前路径和加入map中
    res += dfs(root.Left, curSum, targetSum, mp) + dfs(root.Right, curSum, targetSum, mp) // 遍历左右子树
    mp[curSum]-- // 遍历完后，将当前路径和从map中删除

    return res
}
```