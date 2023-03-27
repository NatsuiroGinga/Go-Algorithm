# CodeForces 55B Smallest number
链接: https://codeforces.com/contest/55/problem/B
## 题意
现有4个整数(均小于等于1000)，并给出三个运算符(均为+或*)。要求每次取出不一定相邻的两个数，并依次使用给出的运算符对这两个数进行运算，并将结果当做一个新数如此操作，直到只剩下一个数为止。 编程求出最后剩下数的最小值。
## 思路
dfs + 枚举
1. 由于4个数的组合有4!种，所以可以枚举所有的组合，然后对每种组合进行dfs。
2. dfs的过程中，每次取出两个数`nums[i]`和`nums[j]`，将`nums[j]`标记为已访问，然后对`nums[i]`和`nums[j]`进行运算，将结果放在`nums[i]`的位置，然后对下一个运算符进行dfs。
3. 如果当前运算符为`+`，则将`nums[i]`和`nums[j]`相加，否则将`nums[i]`和`nums[j]`相乘。
4. 如果当前3个运算符已用完，则将4个数中未被标记的数与最小值进行比较，更新最小值。
5. 如果当前3个运算符未用完，则对下一个运算符进行dfs。
## 代码
```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	nums      [4]int64 // 4个数
	operators [3]string // 3个运算符
	ans       int64 // 最小值
	vis       [4]bool // 标记数组
)

func main() {
	ans = math.MaxInt64 // 初始化最小值为最大值
	gin := bufio.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		fmt.Scan(&nums[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Fscan(gin, &operators[i])
	}
	
	dfs(0) // dfs, 从第一个运算符开始

	fmt.Println(ans)
}

func min(a, b int64) int64 { // min函数
	if a < b {
		return a
	}
	return b
}

func dfs(cur int) {
	if cur == 3 { // 如果3个运算符已用完
		for i := 0; i < 4; i++ { // 遍历4个数
			if !vis[i] { // 如果该数未被标记
				ans = min(ans, nums[i]) // 更新最小值
			}
		}
		return
	}

	for i := 0; i < 4; i++ { // 枚举两个运算数中的第一个数
		if vis[i] { // 如果该数已被标记, 跳过
			continue
		}

		for j := i + 1; j < 4; j++ { // 枚举两个运算数中的第二个数
			if vis[j] { // 如果该数已被标记, 跳过
				continue 
			}

			tmp := nums[i] // 保存nums[i]的值
			vis[j] = true // 标记nums[j]为已访问

			// 进行运算
			if operators[cur] == "+" { 
				nums[i] += nums[j]
			} else {
				nums[i] *= nums[j]
			}
                         // 剪枝, 如果当前最小值大于等于当前最小值, 意味着后续的运算不会使最小值变小, 因此可以直接跳过
			if nums[i] < ans {
				dfs(cur + 1)
			}
			
                        // 回溯
			vis[j] = false
			nums[i] = tmp
		}
	}
}

```