package day11

func findAnagrams(s string, p string) []int {
	var (
		slen = len(s)
		plen = len(p)
		freq = make(map[byte]int, 26)
		ans  = make([]int, 0, slen)
	)
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
		freq[c]-- // 出现过则频率-1

		if right-left+1 == plen { // 如果找到了窗口大小为p长度的窗口
			if count == plen { // 如果窗口中的字符串满足字谜的条件
				ans = append(ans, left) // 将left指向的index加入答案
			}

			// 如果left指向的字符在p中出现过
			if v := freq[s[left]]; v >= 0 {
				// 表示窗口size-1
				count--
			}

			// 窗口左边界前进1步
			freq[s[left]]++
			left++
		}
	}

	return ans
}
