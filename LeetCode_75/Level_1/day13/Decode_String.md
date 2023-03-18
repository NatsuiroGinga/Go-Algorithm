# LeetCode 394. Decode String
https://leetcode.com/problems/decode-string/
## 思路
1. 用两个栈`countStack`和`resStack`来存储数字和字符串
2. 遍历字符串，如果是数字，就计算数字，然后入栈`countStack`
3. 如果是'[', 就把当前的字符串入栈`resStack`，然后清空当前字符串
4. 如果是']', 就把`resStack`栈顶的字符串取出，然后重复`countStack`栈顶的次数，然后赋值给当前字符串
5. 如果是字母，就直接加到当前字符串后面
6. 最后返回当前字符串
## 代码
```go
func decodeString(s string) string {
    countStack := []int{}
    resStack := []string{}
    res := ""
    count := 0
    for _, c := range s {
        if c >= '0' && c <= '9' {
            count = count * 10 + int(c - '0')
        } else if c == '[' {
            countStack = append(countStack, count)
            resStack = append(resStack, res)
            count = 0
            res = ""
        } else if c == ']' {
            tmp := resStack[len(resStack) - 1]
            resStack = resStack[:len(resStack) - 1]
            for i := 0; i < countStack[len(countStack) - 1]; i++ {
                tmp += res
            }
            res = tmp
            countStack = countStack[:len(countStack) - 1]
        } else {
            res += string(c)
        }
    }
    return res
}
```

## 优化
拼接字符串时可以使用`strings.Builder`或者`bytes.Buffer`，这样可以减少内存分配的次数，提高性能