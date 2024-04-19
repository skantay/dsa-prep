package main

func main() {
	tests := []struct { // change types based on your exercise
		want	int // expected output
		values	any // values which function will expect
	}{
		{// describe test cases
			want: 1,
			values: 1,
		},
	}


	var count int

	for i, tt := range tests {
		got := testingFunction(tt.values) // test your function

		if got != tt.want {
			count++
			fmt.Printf("test case: #%s\ngot: %d\nwant: %d\n-------\n", i+1,got,tt.want)
		}
	}

	if count != 0 {
		fmt.Println("don't give up, try again")
	} else {
		fmt.Println("good job")
	}
}


// edit this function
func testingFunction(values any) int {
	return 0
}
