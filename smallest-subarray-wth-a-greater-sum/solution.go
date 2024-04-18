package main

import "fmt"

func main() {
	tests := []struct{
		S int
		arr []int
		want int
	}{
		{
			S: 7,
			arr: []int{2,1,5,2,3,2},
			want: 2,
		},
		{
			S: 7,
			arr: []int{2,1,5,2,8},
			want: 1,
		},
		{
			S: 8,
			arr: []int{3,4,1,1,6},
			want: 3,
		},
	}
	
	for _, tt := range tests {
		fmt.Printf("got: %d, want %d\n",findMinSubArray(tt.S, tt.arr), tt.want)
	}
}

func findMinSubArray(S int, arr []int) int {
	var smallest int = len(arr)

	var l, r int

	var sum int

	for r < len(arr) {
		sum += arr[r]
		if sum < S {
			r++
		} else if sum >= S {
			
			smallest = min(smallest, r - l + 1)
			sum -= arr[l]
			
			l++
			
			if l == r {
				sum -= arr[l]
			} else {
				sum -= arr[r]
			}
		}
	}

	return smallest
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
