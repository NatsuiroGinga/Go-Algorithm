# LeetCode 454. 4Sum II

## 题目描述

Given four lists A, B, C, D of integer values, compute how many tuples `(i, j, k, l)` there are such that `A[i] + B[j] + C[k] + D[l]` is zero.

To make problem a bit easier, all A, B, C, D have same length of N where `0 ≤ N ≤ 500`. All integers are in the range of `-2^28` to `2^28 - 1` and the result is guaranteed to be at most `2^31 - 1`.

**Example:**

```

Input:

A = [ 1, 2]

B = [-2,-1]

C = [-1, 2]

D = [ 0, 2]

Output:

2

Explanation:

The two tuples are:

1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0

2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

```

## 解题思路

1. 使用哈希表记录A和B中元素的和，以及出现的次数
2. 遍历C和D中元素的和，如果其相反数在哈希表中出现，则将其出现次数加到结果中
3. 遍历完C和D中元素的和，返回结果
4. 时间复杂度为O(n^2)，空间复杂度为O(n^2)

## AC代码

```python
class Solution(object):
    def fourSumCount(self, A, B, C, D):
        """
        :type A: List[int]
        :type B: List[int]
        :type C: List[int]
        :type D: List[int]
        :rtype: int
        """
        res = 0
        sumAB = {}
        for a in A:
            for b in B:
                if a + b in sumAB:
                    sumAB[a + b] += 1
                else:
                    sumAB[a + b] = 1
        for c in C:
            for d in D:
                if -(c + d) in sumAB:
                    res += sumAB[-(c + d)]
        return res
```

