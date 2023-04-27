package day1

func isBadVersion(version int) bool {
	return false
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 *
 */
func firstBadVersion(n int) int {
	var (
		low  = 0
		high = n
		ans  = -1
	)

	for low <= high {
		mid := (low + high) / 2
		if isBadVersion(mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}
