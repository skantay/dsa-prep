package leetcode

func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for i := range strs {
        letters := []rune(strs[i])

        slices.SortFunc(letters, func(a, b rune) int {
            if a < b {
                return -1
            } else if a == b {
                return 0
            }

            return 1
        })

        slice := m[string(letters)]
        
        slice = append(slice, strs[i])

        m[string(letters)] = slice
    }

    anagrams := make([][]string, len(m))

    i := 0
    for _, v := range m {
        for _, v2 := range v {
            anagrams[i] = append(anagrams[i], v2)
        }
        i++
    }

    return anagrams
}

