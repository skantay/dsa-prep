package leetcode

func lengthOfLongestSubstring(s string) int {
    var lenght int

    var l int

    letters := make(map[byte]int)

    for r := range s {
        if contains(letters, s[r]) && letters[s[r]] >= l {
            l = letters[s[r]] + 1
        } else {
            lenght = max(lenght, r - l + 1)
        }
        letters[s[r]] = r
    }

    return lenght
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func contains(letters map[byte]int, letter byte) bool {
    _, ok := letters[letter]

    return ok
}
