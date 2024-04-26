package leetcode 

func containsNearbyDuplicate(nums []int, k int) bool {
    var l, r int = 0, 0

    var kClone int = k

    if len(nums) == 1 || k == 0 {
        return false
    }

    for r < len(nums) {
        if r != len(nums)-1 && k != 0 {
            r++
            k--
        } else {
            l++
            k = kClone
            r = l + 1
            k--

            if l == len(nums)-1 {
                break
            }
        }

        if nums[l] == nums[r] && l != r {
            return true
        }
    }

    return false
}

func abs(x int) int{
    if x < 0 {
        return x * -1
    }

    return x
}
