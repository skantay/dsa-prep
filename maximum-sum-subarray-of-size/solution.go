package main

import "fmt"

func main() {
	tests := []struct{
		S int
		arr []int
		want int
	}{
		{
			S: 3,
			arr: []int{2,1,5,1,3,2},
			want: 9,
		},
		{
			S: 2,
			arr: []int{2,3,4,1,5},
			want: 7,
		},
	}
	
	for _, tt := range tests {
		fmt.Printf("got: %d, want %d\n",findMaxSumSubArray(tt.S, tt.arr), tt.want)
	}
}

func findMaxSumSubArray(k int, arr []int) int {
	var result int

	var l, r int

	var sum int

	for r < len(arr) {
		if k > 0 {
			sum += arr[r]
			r++
			k--
		} else {
			result = max(result, sum)
			sum -= arr[l]
			sum += arr[r]
			l++	
			r++
		}
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
