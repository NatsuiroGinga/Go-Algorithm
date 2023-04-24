# LeetCode 131.Palindrome Partitioning

## 题目
Given a string s, partition s such that every
substring
of the partition is a
palindrome
. Return all possible palindrome partitioning of s.

## 题意
给定一个字符串，将其分割成若干个回文串，返回所有可能的分割方案。

## 示例

### 示例1
```text
Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
```
### 示例2
```text
Input: "a"
Output:
[
  ["a"]
]
```

## 分析
1. 回溯法，递归求解。
2. 递归函数的参数为字符串`s`和当前分割方案的起始位置`cur`。
3. 递归函数的终止条件为当前分割方案的起始位置等于字符串`s`的长度。
4. 递归函数的逻辑为：
    1. 从当前分割方案的起始位置开始，遍历字符串`s`。
    2. 判断当前分割方案的一部分（范围是[`cur`, `i`]）是否为回文串，如果是，则将其加入到当前分割方案中。
    3. 递归调用函数，将当前分割方案的起始位置设置为当前分割方案的一部分的下一个位置(`i+1`)。
    4. 如果当前分割方案不为空，则将其最后一个元素删除。
5. 递归函数的返回值为void。

## 代码
```go
package main

var (
	ret  [][]string
	path []string
)

func partition(s string) [][]string {
	ret = make([][]string, 0)
	path = make([]string, 0)
	dfs(0, s)
	return ret
}

func isPalindrome(s string) bool {
	for i, sz := 0, len(s); i < sz/2; i++ {
		if s[i] != s[sz-i-1] {
			return false
		}
	}
	return true
}

func dfs(cur int, s string) {
	if cur == len(s) {
		subset := make([]string, len(path))
		copy(subset, path)
		ret = append(ret, subset)
		return
	}
	for i := cur; i < len(s); i++ {
		substr := s[cur : i+1]
		if isPalindrome(substr) {
			path = append(path, substr)
		} else {
			continue
		}

		dfs(i+1, s)

		if len(path) > 0 {
			path = path[:len(path)-1]
		}
	}
}

```

