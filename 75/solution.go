package leetcode

func sortColors(nums []int)  {
    for i := 0; i < len(nums)-1; i++ {
		for i2 := i; i2 < len(nums)-1-i; i2++ {
			if nums[i2] >  nums[i2+1] {
				nums[i2], nums[i2+1] = nums[i2+1], nums[i2] 
			}
		}	
	}	
}