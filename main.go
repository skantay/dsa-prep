package main

import "fmt"

type say struct {
	m string
}

type p struct {
	*say
}

func (s *say) print() {
	fmt.Println(s)
}

func main() {
	ad := p{}
	ad.print()
}
