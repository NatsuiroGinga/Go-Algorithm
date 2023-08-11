package two_pointers

import "unsafe"

func reverseWords(s string) string {
	b := unsafe.Slice(unsafe.StringData(s), len(s))
	// 1. 去除多余空格
	removeExtraSpaces(&b)
	// 2. 整个字符串进行反转
	reverse(b, 0, len(b)-1)
	// 3. 对每个单词进行反转
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}

	return unsafe.String(unsafe.SliceData(b), len(b))
}

func removeExtraSpaces(b *[]byte) {
	slow := 0
	for i := 0; i < len(*b); i++ {
		if (*b)[i] != ' ' {
			if slow != 0 { // 除了第一个单词, 其余单词前面都要有一个空格
				(*b)[slow] = ' '
				slow++
			}
			for i < len(*b) && (*b)[i] != ' ' {
				(*b)[slow] = (*b)[i]
				slow++
				i++
			}
		}
	}
	*b = append([]byte{}, (*b)[:slow]...)
}

func reverse(b []byte, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
