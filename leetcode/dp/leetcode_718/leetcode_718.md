# LeetCode 718. Maximum Length of Repeated Subarray

## Problem

Given two integer arrays `nums1` and `nums2`, return *the maximum length of a subarray that appears in **both** arrays*.

## 翻译

给定两个整数数组 `nums1` 和 `nums2`，返回两个数组中**公共子数组**的最大长度。

**Example 1:**

```

Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]

Output: 3

Explanation: The repeated subarray with maximum length is [3,2,1].

```

**Example 2:**

```

Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]

Output: 5

```

**Constraints:**

- `1 <= nums1.length, nums2.length <= 1000`
- `0 <= nums1[i], nums2[i] <= 100`

## 题目分析

1. 本题是一道典型的动态规划题目，我们可以使用动态规划的思想来解决这道题目。
2. 我们可以使用一个二维数组来存储两个数组中公共子数组的长度，`dp[i][j]`表示以`nums1[i]`和`nums2[j]`结尾的公共子数组的长度。
3. 如果`nums1[i] == nums2[j]`，那么`dp[i][j] = dp[i-1][j-1] + 1`，否则`dp[i][j] = 0`。
4. 我们可以使用一个变量来记录最大的公共子数组的长度，这样就可以得到最终的结果。
5. 时间复杂度为`O(n^2)`，空间复杂度为`O(n^2)`。

## 代码

**python**

```python

class Solution:

    def findLength(self, nums1: List[int], nums2: List[int]) -> int:
        m, n = len(nums1), len(nums2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        res = 0
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if nums1[i - 1] == nums2[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1] + 1
                    res = max(res, dp[i][j])
        return res

```

**cpp**

```cpp

class Solution {
public:
    int findLength(vector<int>& nums1, vector<int>& nums2) {
        int m = nums1.size(), n = nums2.size();
        vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));
        int res = 0;
        for (int i = 1; i <= m; ++i) {
            for (int j = 1; j <= n; ++j) {
                if (nums1[i - 1] == nums2[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                    res = max(res, dp[i][j]);
                }
            }
        }
        return res;
    }
};

```