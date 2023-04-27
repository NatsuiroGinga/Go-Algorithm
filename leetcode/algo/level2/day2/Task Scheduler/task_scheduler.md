# LeetCode 621. Task Scheduler
## 题目描述
给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

然而, 每两个相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

返回 CPU 处理任务的最短时间。
## 示例
### 示例1
输入: tasks = ["A","A","A","B","B","B"], n = 2
输出: 8
执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
### 示例2
输入: tasks = ["A","A","A","B","B","B"], n = 0
输出: 6
执行顺序: A -> B -> A -> B -> A -> B.
### 示例3
输入: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
输出: 16
执行顺序: A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A -> (待命) -> (待命) -> A.
## 思路
### 首先
需要明确CPU处理任务的最短时间 = 总任务数 + 最少的冷却时间, 以示例1为例, 如下图所示.

![示例1.png](示例1.png)
总任务数为6, 冷却时间为2, 所以CPU处理任务的最短时间为8.

因此, 我们的目标就是求出最少的冷却时间, 最后再加上总任务数即可.
### 其次
使用贪心算法, 每次都优先安排数量最多的任务, 因为如果要使得冷却时间最短, 那么就要先安排数量最多的任务——数量最多的任务所需的冷却时间最多, 中间的冷却时间安排其他任务, 这样才能使得冷却时间最短.

比如, 任务A的数量为3, 任务B的数量为2, 任务C的数量为1, 冷却时间为2, 那么我们就先安排任务A, 然后安排任务B, 最后安排任务C, 这样就能使得冷却时间最短.
如图
![示例2.png](示例2.png)

每两个A任务之间都有两个空闲, 可以安排其他任务, 这个空闲的数量 = (任务A的数量 - 1) * 冷却时间
这里就是 (3 - 1) * 2 = 4

### 然后
在每两个A之间的空闲里安排其他任务, 如下图所示

![示例3.png](示例3.png)

可以看到, 最少的冷却时间 = (A出现的次数 - 1) * 冷却时间 - (除A外, 其他任务的数量)

综上, 我们可以得到一个公式:

假定:
* 出现次数最多的任务为A
* A的数量为max
* 被A分隔的部分数量为partCount
* 每个被A分隔的部分的长度为partLength(也就是空闲时间)
* 除A外, 其他任务的数量为availableTasks
* 所需最少冷却时间为idles
```
partCount = max - 1
partLength = n
availableTasks = len(tasks) - max
idles = partCount * partLength - availableTasks 
```

### 但是
如果有多个任务的数量都是max, 比如, 任务A的数量为3, 任务B的数量为3, 任务C的数量为1, 冷却时间`n`为2.
那么, 上面的公式就不适用了, 如下图所示

![示例5.png](示例5.png)

我们要把多个数量相同的任务放在一起, 让它们尽早的完成, 这样才能使得冷却时间最短.
这样一来, 可以看到被AB分隔的部分的长度就变成了1, 而不是n = 2了。
这时候可以把AB视为一个任务`X`, 如下图所示

![示例6.png](示例6.png)

可以用一个变量`maxCount`来记录出现次数最多的任务的数量,
那么每两个`X`任务之间的空闲数量为: n - (maxCount - 1), 这里的`maxCount - 1`是因为`X`任务的数量为`maxCount`

所以公式就变成了:
```
partCount = max - 1
partLength = n - (maxCount - 1)
availableTasks = len(tasks) - max * maxCount
idles = partCount * partLength - availableTasks 
```

### 但是
如果多个任务的出现次数之和大于`n`, 比如, 3A, 3B, 3C, 3D, 3E, 冷却时间`n`为2.
这时候, `maxCount = 5, partLength = n - (maxCount - 1) = -2`, 这是不合理的, 因为两个任务之间的空闲时间不能小于0.

![示例7](示例7.png)
**太长了, 只画前两个部分**

可见, 在这种情况下, 冷却时间为0

公式就变成了:
``` 
partCount = max - 1
partLength = n - (maxCount - 1)
availableTasks = len(tasks) - max * maxCount
idles = max(0, partCount * partLength - availableTasks) 
```

## 代码
```go
func leastInterval(tasks []byte, n int) int {
	fre := make([]int, 26) // 每个任务的频率
	var max, maxCount int  // 最大频率，最大频率的任务数量

	for _, t := range tasks {
		fre[t-'A']++          // 统计频率
		if fre[t-'A'] > max { // 更新最大频率
			max = fre[t-'A']
			maxCount = 1
		} else if fre[t-'A'] == max { // 更新最大频率的任务数量
			maxCount++
		}
	}

	partCount := max - 1                        // 间隔的数量
	partLen := n - (maxCount - 1)               // 间隔的长度
	emptySlots := partLen * partCount           // 空闲的位置
	availableTasks := len(tasks) - max*maxCount // 可用的任务数量
	idles := Max(0, emptySlots-availableTasks)  // 空闲的位置不够用，需要补充

	return idles + len(tasks) // 空闲的位置 + 任务数量
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```