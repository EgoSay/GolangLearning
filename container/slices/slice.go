package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}
func main() {
	var s []int
	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := make([]int, 2)
	s1 = append(s1, 2)
	s1 = append(s1, 1)
	printSlice(s1)

	s2 := make([]int, 4, 5)
	tmp := s1[:3]
	tmp2 := s1[2:]
	fmt.Println(tmp, tmp2)
	s2 = append(s1[:1], s1[2:]...)
	printSlice(s2)
}
