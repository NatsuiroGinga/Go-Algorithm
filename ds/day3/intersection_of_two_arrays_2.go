package day3

func intersect(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int, len(nums1))
	ret := make([]int, 0, len(nums1)+len(nums2))

	// 1. build a map of nums1 with value as count
	for _, num := range nums1 {
		// 2. if num is in map, increment its value
		if _, ok := mp[num]; ok {
			mp[num]++
		} else { // 3. else, add it to map
			mp[num] = 1
		}
	}

	// 4. iterate over nums2
	for _, num := range nums2 {
		// 5. if num is in map, add it to result and decrement its value
		if v, ok := mp[num]; ok && v > 0 {
			ret = append(ret, num)
			mp[v]--
		}
	}

	return ret
}
