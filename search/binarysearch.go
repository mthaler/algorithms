package search

func BinarySearch(a []interface{}, comp func(interface{}) int) int {
	return binarySearch(a, 0, len(a) - 1, comp)
}

func binarySearch(a []interface{}, l, r int, comp func(interface{}) int) int {
	if (r >= l) {
		m := l + (r-l)/2
		c := comp(a[m])
		if c == 0 {
			return m
		} else if c > 0 {
			return binarySearch(a, l, m - 1, comp)
		} else {
			return binarySearch(a, m + 1, r, comp)
		}
	} else {
		return -1
	}
}
