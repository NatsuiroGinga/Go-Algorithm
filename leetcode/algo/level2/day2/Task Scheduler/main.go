package main

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
