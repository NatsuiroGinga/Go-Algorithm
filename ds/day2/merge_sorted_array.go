package day2

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		i = m - 1 // pointer to nums1
		j = n - 1 // pointer to nums2
		k = m + n - 1
	)

	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
			k--
		} else {
			nums1[k] = nums1[i]
			i--
			k--
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
