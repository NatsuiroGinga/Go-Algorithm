# LeetCode 424. Longest Repeating Character Replacement
https://leetcode.com/problems/longest-repeating-character-replacement/
## 题目描述
You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.
## 思路
1. 滑动窗口
2. 用一个map记录窗口中每个字符出现的次数, 取名fre
3. 用一个变量记录窗口中出现次数最多的字符的次数, 取名maxCount
4. 如果窗口的长度减去出现次数最多的字符的次数大于k，那么就需要移动左指针，直到窗口的长度减去出现次数最多的字符的次数小于等于k
   * 例如：s = "AABABBA", k = 1, 当窗口为[0, 2]时，map中记录的字符出现的次数为{'A': 2, 'B': 1}, maxCount = 2, 此时窗口的长度减去出现次数最多的字符的次数为2(right) - 0(left) + 1 - 2(maxCount) = 1 == 1(k), 此时窗口的长度减去出现次数最多的字符的次数小于等于k, 不需要移动左指针
5. 每次移动左指针时，需要将左指针指向的字符在map中的次数减1, left++
6. 每次移动右指针时，需要将右指针指向的字符在map中的次数加1, right++
7. 每次移动左指针或者右指针时，都需要更新maxCount, maxCount = max(maxCount, fre[s[right]])

## 数据结构
```go
var (
		size     = len(s) // s字符串的长度
		fre      = make(map[byte]int, 26) // map记录窗口中每个字符出现的次数, 因为题目中说了只有大写字母，所以map的长度为26
		maxLen   = 0 // 最长的子串的长度
		maxCount = 0 //  窗口中出现次数最多的字符的次数
 )
```

## 算法
```go
func characterReplacement(s string, k int) int {
	for left, right := 0, 0; right < size; right++ {
		fre[s[right]]++               // 右指针指向的字符在map中的次数加1
		if maxCount < fre[s[right]] { // 更新maxCount
			maxCount = fre[s[right]]
		}

		if right-left+1-maxCount > k { // 如果窗口的长度减去出现次数最多的字符的次数大于k，那么就需要移动左指针
			fre[s[left]]-- // 左指针指向的字符在map中的次数减1
			left++         // 窗口左指针右移
		}

		if maxLen < right-left+1 { // 更新maxLen
			maxLen = right - left + 1
		}
	}

	return maxLen
}
```