package main

import "fmt"

func add() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func main() {
	a := add()
	for i := 0; i < 10; i++ {
		fmt.Printf("1 +2 + ..... +%d = %d\n", i, a(i))
	}
}
