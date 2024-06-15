package leetcode

func numsSameConsecDiff(n int, k int) []int {
    ans := make([]int, 0)
    for cur := 1; cur <= 9; cur++ {
        try(&ans, cur, n - 1, k)    
    }
    return ans
}

func try(ans *[]int, curNum, n, k int) {
    if n == 0 {
        (*ans) = append((*ans), curNum)
        return
    }
    d := curNum % 10

    if d - k >= 0 { 
        try(ans, curNum * 10 + (d - k), n - 1, k)
    }
    if k > 0 && d + k <= 9 {
        try(ans, curNum * 10 + (d + k), n - 1, k)
    }    
}
