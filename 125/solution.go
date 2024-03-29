package leetcode

func isPalindrome(s string) bool {
	var word string

	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			word += string(v)
		} else if v >= 'A' && v <= 'Z' {
			word += string(v + 32)
		}
	}

	i, j := 0, len(word)-1

	for i < j {
		if word[i] == word[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}
