package leetcode

func longestOnes(nums []int, k int) int {
    l, r := 0,0 

    var lenght int

    for r < len(nums) {
        if nums[r] == 1 {
            r++
        } else if nums[r] == 0 && k != 0 {
            k--
            r++
        } else if nums[l] == 0 {
            k++
            l++
        } else if nums[l] == 1 {
            l++
        }

        lenght = max(lenght, r - l)
    }

    return lenght
}

