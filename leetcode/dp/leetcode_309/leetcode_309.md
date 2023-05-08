# LeetCode 309. 最佳买卖股票时机含冷冻期

### 题目描述

给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，​你可以尽可能地完成更多的交易（多次买卖一支股票）:

- 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
- 卖出股票后，你无法在第二天买入股票 (即冷冻期为1天)

## 示例:

```
输入: [1,2,3,0,2]
输出: 3 
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
```

### 解题思路

1. 状态定义：dp[i][j]表示第i天的状态为j时的最大利润，j=0表示持有股票，j=1表示不持有股票且处于冷冻期，j=2表示不持有股票且不处于冷冻期
2. 状态转移方程：
   - dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
   ```text
    dp[i-1][0] 表示第i-1天就持有股票
    dp[i-1][2] - prices[i] 表示第i-1天不持有股票且不处于冷冻期，第i天买入股票
   ```
   - dp[i][1] = dp[i-1][0] + prices[i]
    ```text
     dp[i-1][0] + prices[i] 表示第i-1天持有股票，第i天卖出股票
    ```
   - dp[i][2] = max(dp[i-1][1], dp[i-1][2])
    ```text
     dp[i-1][1] 表示第i-1天不持有股票且处于冷冻期，第i天不操作
     dp[i-1][2] 表示第i-1天不持有股票且不处于冷冻期，第i天不操作
    ```
3. 初始状态：dp[0][0] = -prices[0]，dp[0][1] = 0，dp[0][2] = 0
    ```text
    dp[0][0] = -prices[0] 表示第0天持有股票，只能是第0天买入股票
    dp[0][1] = 0 表示第0天不持有股票且处于冷冻期，只能是第0天卖出股票
    dp[0][2] = 0 表示第0天不持有股票且不处于冷冻期，只能是第0天不操作
    ```
4. 返回值: max(dp[n-1][1], dp[n-1][2])
    ```text
    因为持有股票的状态不可能是最大利润，
    所以只需要比较不持有股票且处于冷冻期和不持有股票且不处于冷冻期的最大利润即可
    
    dp[n-1][1] 表示第n-1天不持有股票且处于冷冻期，第n天不操作
    dp[n-1][2] 表示第n-1天不持有股票且不处于冷冻期，第n天不操作
    ```
   
### 代码

**Java**

```java

public class Solution {
    public int maxProfit(int[] prices) {
        if (prices == null || prices.length == 0) {
            return 0;
        }
        int n = prices.length;
        int[][] dp = new int[n][3];
        dp[0][0] = -prices[0];
        dp[0][1] = 0;
        dp[0][2] = 0;
        for (int i = 1; i < n; i++) {
            dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][2] - prices[i]);
            dp[i][1] = dp[i - 1][0] + prices[i];
            dp[i][2] = Math.max(dp[i - 1][1], dp[i - 1][2]);
        }
        return Math.max(dp[n - 1][1], dp[n - 1][2]);
    }
}

```

### 复杂度分析

- 时间复杂度：O(n)，n为数组长度
- 空间复杂度：O(n)，n为数组长度

**Python3**

```python

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if not prices:
            return 0
        n = len(prices)
        dp = [[0] * 3 for _ in range(n)]
        dp[0][0] = -prices[0]
        dp[0][1] = 0
        dp[0][2] = 0
        for i in range(1, n):
            dp[i][0] = max(dp[i - 1][0], dp[i - 1][2] - prices[i])
            dp[i][1] = dp[i - 1][0] + prices[i]
            dp[i][2] = max(dp[i - 1][1], dp[i - 1][2])
        return max(dp[n - 1][1], dp[n - 1][2])

```
