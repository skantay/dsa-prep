package leetcode


type p struct {
    v int
    k int
}
func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)

    for i := range nums {
        m[nums[i]]++
    }

    var sl []p

    for k, v := range m {
        p := p{
            v: k,
            k: v,
        }

        sl = append(sl, p)
    }

    slices.SortFunc(sl, func(a, b p) int {
        if a.k < b.k {
            return 1
        } else if a.k == b.k {
            return 0
        }

        return -1
    })

    var res []int

    for i := range sl[:k] {
        res = append(res, sl[i].v)
    }

    return res
}
