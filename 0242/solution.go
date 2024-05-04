package leetcode

func isAnagram(s string, t string) bool {
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
