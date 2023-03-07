package day2

func twoSum(nums []int, target int) []int {
	// record num and it's index
	mp := make(map[int]int, len(nums))
	// traverse the nums
	for i, num := range nums {
		tmp := target - num
		// check if map contains (target - num)
		if v, ok := mp[tmp]; ok {
			return []int{i, v}
		}
		mp[num] = i
	}

	return []int{}
}
