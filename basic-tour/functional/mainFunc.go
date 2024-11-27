package main

import (
	"Learning/functional/Fibonacci"
	"fmt"
)

func main() {
	f := Fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
