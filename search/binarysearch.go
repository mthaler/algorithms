package search

func BinarySearch(a []interface{}, l, r int, comp func(interface{}) int) int {
	if (r >= l) {
		m := l + (r-l)/2
		c := comp(a[m])
		if c == 0 {
			return m
		} else if c > 0 {
			return BinarySearch(a, l, m - 1, comp)
		} else {
			return BinarySearch(a, m + 1, r, comp)
		}
	} else {
		return -1
	}
}
