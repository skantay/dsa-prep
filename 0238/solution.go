package leetcode

func productExceptSelf(nums []int) []int {
    l := make([]int, len(nums))
    r := make([]int, len(nums))

    l[0] = 1

    for i := 1; i < len(l); i++ {
        l[i] = l[i-1] * nums[i-1]
    }

    r[len(r)-1] = 1

    for i := len(r) - 2; i >= 0; i-- {
        r[i] = r[i + 1] * nums[i+1]
    }

    res := make([]int, len(nums))

    for i := range l {
        res[i] = l[i] * r[i]
    }

    return res
}
