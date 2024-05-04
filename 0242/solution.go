package leetcode

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int, len(s))

	for i := range s {
		m[rune(s[i])]++
	}

	var cnt int

	for _, i := range t {
		if v, ok := m[i]; ok {
			if v != 0 {
				cnt++
				m[i]--
			}
		}
	}

	return cnt == len(s)
}

func isAnagram2(s string, t string) bool {
	v1 := []rune(s)
	v2 := []rune(t)

	f := func(a, b rune) int {
		if a < b {
			return -1
		} else if a == b {
			return 0
		}

		return 1
	}

	slices.SortFunc(v1, f)
	slices.SortFunc(v2, f)

	return string(v1) == string(v2)
}
