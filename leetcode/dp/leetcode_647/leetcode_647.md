# LeetCode 647. 回文子串

## 题目描述

给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。

**示例 1:**

```

输入: "abc"

输出: 3

解释: 三个回文子串: "a", "b", "c".

```

**示例 2:**

```

输入: "aaa"

输出: 6

说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".

```

**注意:**

1. 输入的字符串长度不会超过1000。

## 解题思路

1. 状态定义：dp[i][j]表示字符串s从i到j的子串是否为回文子串
2. 状态转移方程：dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
3. 初始状态：dp[i][i] = true
4. 设置result变量记录回文子串的个数，每次dp[i][j]为true时，result++
5. 返回值：result

## 代码

**java**

```java

class Solution {
    public int countSubstrings(String s) {
        int n = s.length();
        boolean[][] dp = new boolean[n][n];
        int result = 0;
        for(int i = n-1; i >= 0; i--){
            for(int j = i; j < n; j++){
                dp[i][j] = s.charAt(i) == s.charAt(j) && (j - i < 2 || dp[i+1][j-1]);
                if(dp[i][j]){
                    result++;
                }
            }
        }
        return result;
    }
}

```

**python**

```python

class Solution:
    def countSubstrings(self, s: str) -> int:
        n = len(s)
        dp = [[False] * n for _ in range(n)]
        result = 0
        for i in range(n-1, -1, -1):
            for j in range(i, n):
                dp[i][j] = s[i] == s[j] and (j - i < 2 or dp[i+1][j-1])
                if dp[i][j]:
                    result += 1
        return result

```