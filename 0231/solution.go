package leetcode

func isPowerOfTwo(n int) bool {
    if n == 1{
        return true
    } else if n % 2 == 1 || n == 0 {
        return false
    }
    
    return isPowerOfTwo(n/2)
}
