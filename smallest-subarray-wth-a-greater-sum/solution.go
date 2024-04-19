package main

import "fmt"
import "math"

func main() {
	tests := []struct {
		S    int
		arr  []int
		want int
	}{
		{
			S:    10,
			arr:  []int{1, 1, 1},
			want: 0,
		},
		{
			S:    7,
			arr:  []int{2, 1, 5, 2, 8},
			want: 1,
		},
		{
			S:    8,
			arr:  []int{3, 4, 1, 1, 6},
			want: 3,
		},
		{
			S: 10,
			arr: []int{1,1,1,1,1,1,1,1,1,1,1,1,1},
			want: 10,
		},
	}
	var count int
	for index, tt := range tests {
		got := findMinSubArray(tt.S, tt.arr)

		if got != tt.want {
			count++
			fmt.Printf("test case: #%d\ngot: %d\nwanted:%d\n----------\n",index+1,got,tt.want)
		}
	}

	if count != 0 {
		fmt.Println("try again")
	} else {
		fmt.Println("good job")
	}
}

func findMinSubArray(S int, arr []int) int {
	var smallest int = math.MaxInt32

	var l, r int

	var sum int

	for r < len(arr) {
		sum += arr[r]

		if sum >= S {
			smallest = min(smallest, r-l+1)
			sum -= arr[l]
			l++
			sum -= arr[r]
		} else {
			r++
		}
	}

	if smallest == math.MaxInt32 {
		return 0
	}

	return smallest
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
