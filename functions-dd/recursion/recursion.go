package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(num int) int {
	if num < 2 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
