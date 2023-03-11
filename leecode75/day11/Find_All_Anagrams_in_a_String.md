# LeetCode 438. Find All Anagrams in a String
## 题目描述
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
## 思路 1
1. 使用map记录下字符串p中每个字符出现的次数
2. 使用滑动窗口算法, 遍历字符串s, 窗口大小固定为p的大小
3. 窗口每次只向右走一步, 每一轮将窗口右端点处的字符在map中对应的次数-1, 将左端点处的字符在map中的次数+1
4. 如果窗口中所有字符出现的次数与map中数据一致, 就说明匹配到了一个Anagrams
## 数据结构 1
```go
var (
	    // 记录p中每个字符出现的次数, 因为只有小写字母, 所以map的大小为26
		fre    = make(map[byte]int, 26)
		// 记录窗口中每个字符出现的次数
		window = make(map[byte]int, 26)
		// 记录匹配到的Anagrams的起始位置
		ans    []int
)
```
## 算法 1
```go
func findAnagrams(s string, p string) []int {
    // 如果s的长度小于p的长度, 说明s中不可能存在p的Anagrams
	if len(s) < len(p) {
		return []int{}
	}
	for i := range p {
        // 初始化fre
		fre[p[i]]++
		// 初始化window
		window[s[i]]++
	}
    // 如果初始化后, 窗口中的字符出现的次数与p中的字符出现的次数一致, 说明s的前len(p)个字符就是p的Anagrams
	if equal(fre, window) {
		ans = append(ans, 0)
	}
	// 遍历s, 每次窗口向右滑动一步
	for i := len(p); i < len(s); i++ {
		// 窗口向右滑动, 右端点处的字符在map中对应的次数-1
		window[s[i-len(p)]]--
		// 窗口向右滑动, 左端点处的字符在map中对应的次数+1
		window[s[i]]++

		// 如果窗口满足字谜的要求
		if equal(fre, window) {
			// 把窗口的起始位置加入ans
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}

// 判断两个map是否相等
func equal(a, b map[byte]int) bool {
	// 遍历a, a中的字符在b中出现的次数不一致, 说明两个map不相等
	for i, va := range a {
		if va != b[i] {
			return false
		}
	}
	return true
}
```
## 思路 2
1. 使用fre这个map记录下字符串p中每个字符出现的次数, count记录窗口中满足字谜条件的字符个数
2. 使用滑动窗口算法, 遍历字符串s, 设置两个指针left和right, 开始指向0
3. right指针向右移动, 每次移动一步, 如果right指向的字符在p中出现过, count++, 表示窗口中满足条件的字符个数+1
4. right指向的字符在fre中的次数-1, 表示该字符已经被使用过
5. 如果right-left+1 == p的长度, 说明找到了一个窗口大小为p的长度的窗口
6. 如果count == p的长度, 说明窗口中的字符串满足字谜的条件, 将left指向的index加入答案 
7. 如果left指向的字符在p中出现过, count--, 表示窗口中满足条件的字符个数-1
8. left指向的字符在fre中的次数+1, 恢复初始状态, 表示该字符可以再次使用 
9. left指针向右移动一步, left++

## 数据结构 2
```go
var (
	    // 字符串s的长度
		slen = len(s)
		// 字符串p的长度
		plen = len(p)
	    // 记录p中每个字符出现的次数, 因为只有小写字母, 所以map的大小为26 
		freq = make(map[byte]int, 26)
		// 记录匹配到的Anagrams的起始位置
		ans  = make([]int, 0, slen)
)
```
## 算法 2
```go
func findAnagrams(s string, p string) []int {
	if slen < plen {
		return []int{}
	}

	for i := range p {
		freq[p[i]]++
	}

	for left, right, count := 0, 0, 0; right < slen; right++ {
		c := s[right]
		// 判断right指向的字符在p中有没有出现
		if v := freq[c]; v > 0 {
			count++ // 窗口中满足条件的字符个数+1
		}
		freq[c]-- // 出现过则次数-1

		if right-left+1 == plen { // 如果找到了窗口大小为p长度的窗口
			if count == plen { // 如果窗口中的字符串满足字谜的条件
				ans = append(ans, left) // 将left指向的index加入答案
			}
			// 如果left指向的字符在p中出现过
			if v := freq[s[left]]; v >= 0 {
				// 表示窗口中满足条件的字符个数-1
				count--
			}
			// 窗口左边界前进1步
			freq[s[left]]++
			left++
		}
	}

	return ans
}
```