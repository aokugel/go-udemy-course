package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(transformed)

	doubled := createTransformer(2)
	tripled := createTransformer(3)
	quads := createTransformer(4)

	transforemed2 := transformNumbers(&numbers, createTransformer(45))

	fmt.Println(transforemed2, transformNumbers(&numbers, doubled), transformNumbers(&numbers, tripled), transformNumbers(&numbers, quads))
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(degree int) func(int) int {
	return func(number int) int {
		return number * degree
	}
}
