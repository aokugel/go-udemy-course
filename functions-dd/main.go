package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 5, 6}
	fmt.Println(sumup(nums...))
	fmt.Println(nums)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
