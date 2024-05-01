package leetcode

func maxVowels(s string, k int) int {
	var count, res, l int

	for i := 0; i < len(s); i++ {
		if is(rune(s[i])) {
			count++
		}

		if i >= k {

			if is(rune(s[i])) && is(rune(s[l])) {
				count--
			} else if !is(rune(s[i])) && is(rune(s[l])) {
				count--
			}

			l++
		}

		res = max(res, count)
	}

	return res
}

func is(s rune) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}
