package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("1")
		fmt.Print(recover())
	}()
	defer func() {

		fmt.Println("2")
		defer fmt.Print(recover())
		defer panic(1)
		recover()
	}()
	defer recover()
	panic(2)
}
