# LeetCode 583. 两个字符串的删除操作

## 题目描述

给定两个单词 *word1* 和 *word2*，找到使得 *word1* 和 *word2* 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

**示例：**

```

输入: "sea", "eat"

输出: 2

解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"

```

**提示：**

1. 给定单词的长度不超过500。
2. 给定单词中的字符只含有小写字母。

## 解题思路

1. 状态定义：dp[i][j]表示word1的前i个字符和word2的前j个字符相同所需的最小步数
2. 状态转移方程：
   - dp[i][j] = dp[i-1][j-1]，当word1[i-1] == word2[j-1]
   - dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1，当word1[i-1] != word2[j-1]
   - dp[i][0] = i，dp[0][j] = j
3. 初始状态：dp[0][0] = 0
4. 返回值：dp[m][n]

## 代码

**java**

```java

class Solution {
    public int minDistance(String word1, String word2) {
        int m = word1.length();
        int n = word2.length();
        int[][] dp = new int[m+1][n+1];
        for(int i = 0; i <= m; i++){
            dp[i][0] = i;
        }
        for(int j = 0; j <= n; j++){
            dp[0][j] = j;
        }
        for(int i = 1; i <= m; i++){
            for(int j = 1; j <= n; j++){
                if(word1.charAt(i-1) == word2.charAt(j-1)){
                    dp[i][j] = dp[i-1][j-1];
                }else{
                    dp[i][j] = Math.min(dp[i-1][j], dp[i][j-1]) + 1;
                }
            }
        }
        return dp[m][n];
    }
}

```

**复杂度分析**

- 时间复杂度：$O(mn)$
- 空间复杂度：$O(mn)$

**python**

```python

class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        m = len(word1)
        n = len(word2)
        dp = [[0 for _ in range(n+1)] for _ in range(m+1)]
        for i in range(m+1):
            dp[i][0] = i
        for j in range(n+1):
            dp[0][j] = j
        for i in range(1, m+1):
            for j in range(1, n+1):
                if word1[i-1] == word2[j-1]:
                    dp[i][j] = dp[i-1][j-1]
                else:
                    dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
        return dp[m][n]

```
