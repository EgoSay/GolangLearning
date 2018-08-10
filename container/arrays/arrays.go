package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	fmt.Println(arr1)
	var arr2 = make([]int, 3, 5)
	fmt.Println(len(arr2), cap(arr2))

}
