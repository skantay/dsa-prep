func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	res := make([]string, 0)
	curr := make([]rune, 0)
	letterCombinationsHelper(digits, 0, curr, &res)

	return res
}

func letterCombinationsHelper(digits string, index int, curr []rune, res *[]string) {
	if index == len(digits) {
		*res = append(*res, string(curr))
		return
	}

	for _, letter := range mapping(rune(digits[index])) {
		curr = append(curr, letter)
		letterCombinationsHelper(digits, index+1, curr, res)
		curr = curr[:len(curr)-1]
	}
}

func mapping(digit rune) []rune {
	mapping := map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}

	return mapping[digit]
}
