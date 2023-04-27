package day6

func lengthOfLongestSubstring(s string) int {
	var (
		window = make(map[byte]struct{}, len(s)) // 用于存储滑动窗口中的字符和对应的索引
		res    = 0                               // 滑动窗口中不重复字符的最大长度
	)

	for left, right := 0, 0; right < len(s); right++ { // 右指针不断向右移动
		c := s[right]                                   // 取出右指针对应的字符
		for _, ok := window[c]; ok; _, ok = window[c] { // 如果右指针对应的字符已经在窗口中存在
			delete(window, s[left]) // 则删除窗口中最左边的字符
			left++                  // 左指针向右移动
		}
		if res < len(window) {
			res = len(window)
		}
	}

	return res
}
