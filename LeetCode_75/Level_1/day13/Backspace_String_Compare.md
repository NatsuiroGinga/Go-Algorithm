# LeetCode 844. Backspace String Compare
# 题目描述
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

 ## 思路 1
设置两个栈stack1和stack2, 一个用来存储字符串 s, 一个用来存储字符串 t, 然后遍历两个栈, 如果两个栈的元素相同, 则返回 true, 否则返回 false.

 ## 算法 1
```go
func backspaceCompare(s string, t string) bool {
	stack1 := sta(s)
	stack2 := sta(t)

    fmt.Println(stack1)
	fmt.Println(stack2)

	if len(stack1) != len(stack2) {
		return false
	}

	for i := 0; i < len(stack1); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}

	return true
}

func sta(str string) []rune {
	stack := make([]rune, 0)
	for _, v := range str {
		if v != '#' {
			stack = append(stack, v)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return stack
}
```

## 思路 2
1. 首先从后往前遍历字符串, 设置两个指针 `i` 和 `j`, 分别指向字符串 `s` 和 `t` 的末尾, 还有一个`back`变量记录 `#` 的个数
2. 如果遇到 `#`, 则 `back` 加 1, 如果遇到字符, 则判断 `back` 是否大于 0, 如果大于 0, 则 `back` 减 1, 表示下一个字符被删除了
3. 如果两个指针指向的字符相同, 则`i`和`j`继续向前移动, 否则返回 false
4. 如果两个指针都移动到了字符串的开头, 则返回 true

## 算法2
```go
func backspaceCompare(s string, t string) bool {
	i, j, back := len(s)-1, len(t)-1, 0

	for {
		back = 0
		for i >= 0 && (back > 0 || s[i] == '#') {
			if s[i] == '#' {
				back++
			} else {
				back--
			}
			i--
		}

		back = 0
		for j >= 0 && (back > 0 || t[j] == '#') {
			if t[j] == '#' {
				back++
			} else {
				back--
			}
			j--
		}

		if i >= 0 && j >= 0 && s[i] == t[j] {
			j--
			i--
		} else {
			break
		}
	}

	return i == -1 && j == -1
}
```