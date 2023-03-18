package day13

import (
	"bytes"
	"unicode"
)

func decodeString(s string) string {
	var (
		countStack []int        // 用于存储数字
		resStack   []string     // 用于存储字符串
		res        bytes.Buffer // 用于存储结果
		idx        int          // 用于遍历字符串
	)

	for idx < len(s) {
		if unicode.IsDigit(rune(s[idx])) { // 如果是数字
			// 获取数字
			count := 0
			// 累加数字
			for unicode.IsDigit(rune(s[idx])) {
				count = count*10 + int(s[idx]-'0')
				idx++
			}
			countStack = append(countStack, count) // 入栈
		} else if s[idx] == '[' {
			// 当前结果入栈
			resStack = append(resStack, res.String())
			// 清空结果
			res.Reset()
			idx++
		} else if s[idx] == ']' {
			// 获取栈顶的数字
			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			// 获取栈顶的字符串
			tmp := resStack[len(resStack)-1]
			resStack = resStack[:len(resStack)-1]
			// 准备重复拼接字符串
			buffer := bytes.NewBufferString(tmp)
			for i := 0; i < count; i++ {
				// 将当前结果拼接到buffer中, 重复count次
				buffer.WriteString(res.String())
			}
			// 清空结果
			res.Reset()
			// 重置结果
			res.WriteString(buffer.String())
			idx++
		} else {
			// 字母直接拼接
			res.WriteByte(s[idx])
			idx++
		}
	}

	return res.String()
}
