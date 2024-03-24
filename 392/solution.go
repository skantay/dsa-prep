package leetcode

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	var i int

	for i2 := 0; i2 < len(t); i2++ {
		if s[i] == t[i2] {
			i++

			if i == len(s) {
				return true
			}
		}
	}

	return false
}
