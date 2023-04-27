package leetcode_1

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for i, num := range nums {
		if v, ok := record[target-num]; ok {
			return []int{v, i}
		}
		record[num] = i
	}

	return nil
}
