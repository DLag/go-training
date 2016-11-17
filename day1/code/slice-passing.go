package main

import "fmt"

func passByLink(s *[]int) {
	(*s)[0] = 9
	*s = append(*s, 100)
}

func passByReference(s []int) {
	s[0] = 9
	//s = append(s, 100)
}

func main() {
	s := make([]int, 2, 3)
	s[0] = 1
	s[1] = 2
	fmt.Println(s)
	passByReference(s)
	fmt.Println(s)
	passByLink(&s)
	fmt.Println(s)
}
