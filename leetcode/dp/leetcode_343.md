# LeetCode 343. Integer Break

## 题目
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.

Return the maximum product you can get.

## 翻译
给定一个整数n，将其分解为k个正整数的和，其中k>=2，并最大化这些整数的乘积。
返回你可以得到的最大乘积。

## 样例
### Example 1:
```
Input: 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
```
### Example 2:
```
Input: 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
```

## 思路
### 思路1
动态规划，dp[i]表示i的最大乘积，dp[i] = max(dp[i], max(j, dp[j]) * max(i-j, dp[i-j]))，其中j为1到i-1的数。

### 思路2
数学方法，当n>=5时，尽可能多的分解为3，当n=4时，分解为2*2。

## 代码
### 代码1
```java
class Solution {
    public int integerBreak(int n) {
        int[] dp = new int[n + 1];
        dp[1] = 1;
        for (int i = 2; i <= n; i++) {
            for (int j = 1; j < i; j++) {
                dp[i] = Math.max(dp[i], Math.max(j, dp[j]) * Math.max(i - j, dp[i - j]));
            }
        }
        return dp[n];
    }
}
```

### 代码2
```java
class Solution {
    public int integerBreak(int n) {
        if (n == 2) {
            return 1;
        }
        if (n == 3) {
            return 2;
        }
        int res = 1;
        while (n > 4) {
            res *= 3;
            n -= 3;
        }
        return res * n;
    }
}
```