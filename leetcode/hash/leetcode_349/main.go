package leetcode_349

func intersection(nums1 []int, nums2 []int) []int {
	freq := make(map[int]struct{})
	ans := make(map[int]struct{})
	var ret []int

	for _, v := range nums1 {
		freq[v] = struct{}{}
	}

	for _, v := range nums2 {
		if _, ok := freq[v]; ok {
			ans[v] = struct{}{}
		}
	}

	for k := range ans {
		ret = append(ret, k)
	}

	return ret
}
