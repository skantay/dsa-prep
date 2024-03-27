package leetcode

func sortColors(nums []int) {
	var (
		c0 int
		c1 int
		c2 int
	)

	for _, n := range nums {
		switch n {
		case 0:
			c0++
		case 1:
			c1++
		case 2:
			c2++
		}
	}

	for i := 0; i < c0+c1+c2; i++ {
		switch {
		case i < c0:
			nums[i] = 0
		case i >= c0+c1:
			nums[i] = 2
		default:
			nums[i] = 1
		}
	}
}
