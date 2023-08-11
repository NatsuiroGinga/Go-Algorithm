# LeetCode 300. Longest Increasing Subsequence

## Problem

Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

A **subsequence** is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, `[3,6,2,7]` is a subsequence of the array `[0,3,1,6,2,2,7]`.

**Example 1:**

```

Input: nums = [10,9,2,5,3,7,101,18]

Output: 4

Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

```

**Example 2:**

```

Input: nums = [0,1,0,3,2,3]

Output: 4

```

**Example 3:**

```

Input: nums = [7,7,7,7,7,7,7]

Output: 1

```

**Constraints:**

- `1 <= nums.length <= 2500`
- `-104 <= nums[i] <= 104`

**Follow up:** Can you come up with an algorithm that runs in `O(n log(n))` time complexity?

## 翻译

给定一个整数数组 `nums`，返回最长的严格递增子序列的长度。

## 题目分析

1. 本题是一道典型的动态规划题目，我们可以使用动态规划的思想来解决这道题目。
2. 我们可以使用一个一维数组来存储以`nums[i]`结尾的最长严格递增子序列的长度，`dp[i]`表示以`nums[i]`结尾的最长严格递增子序列的长度。
3. 对于`nums[i]`，如果`nums[i] > nums[j]`，那么`dp[i] = max(dp[i], dp[j] + 1)`，其中`0 <= j < i`。
4. 我们可以使用一个变量来记录最长的严格递增子序列的长度，这样就可以得到最终的结果。
5. 时间复杂度为`O(n^2)`，空间复杂度为`O(n)`。

## 代码

**python**

```python

class Solution:

    def lengthOfLIS(self, nums: List[int]) -> int:
        n = len(nums)
        dp = [1] * n
        res = 1
        for i in range(1, n):
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j] + 1)
            res = max(res, dp[i])
        return res

```

**cpp**

```cpp

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        int n = nums.size();
        vector<int> dp(n, 1);
        int res = 1;
        for (int i = 1; i < n; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = max(dp[i], dp[j] + 1);
                }
            }
            res = max(res, dp[i]);
        }
        return res;
    }
};

```