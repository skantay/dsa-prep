package leetcode

func tribonacci(n int) int {
    if n == 0 {
        return 0 
    } else if n == 1 || n == 2 {
        return 1
    }

    nums := make([]int, 38)

    nums[1] = 1

    nums[2] = 1

    for i := 3; i < len(nums); i++ {
        nums[i] = nums[i-1] + nums[i-2] + nums[i-3] 
    }

    return nums[n-1] + nums[n-2] + nums[n-3]
}

