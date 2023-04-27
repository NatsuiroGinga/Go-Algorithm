package main

import "sort"

/*
1. 先排序, 然后固定一个数, 用双指针找另外两个数
2. 因为是从小到大排序, 所以如果三个数的和小于0, 那么左指针右移, 如果大于0, 那么右指针左移
3. 如果等于0, 那么就是答案, 然后左指针右移, 右指针左移, 直到左右指针相等
4. 注意去重, 如果左指针的值和右一个值相等, 那么左指针右移, 如果右指针的值和左一个值相等, 那么右指针左移
5. 注意边界条件, 左指针不能超过右指针, 左指针不能超过数组长度-2, 因为左指针最多到倒数第二个数, 因为右指针是从左指针的下一个数开始的
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 先排序
	var res [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ { // 左指针最多到倒数第二个数
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}
		low, high := i+1, n-1

		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if sum < 0 {
				low++
			} else if sum > 0 {
				high--
			} else {
				res = append(res, []int{nums[i], nums[low], nums[high]})
				for ; low+1 < high && nums[low] == nums[low+1]; low++ {
				}
				for ; high-1 > low && nums[high] == nums[high-1]; high-- {
				}

				low++
				high--
			}
		}
	}

	return res
}

func main() {

}
