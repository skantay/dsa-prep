package leetcode

func numMatchingSubseq(s string, words []string) int {
	var count int

	for i := 0; i < len(words); i++ {
		if isSubsequence(words[i], s) {
			count++
		}
	}

	return count
}

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
