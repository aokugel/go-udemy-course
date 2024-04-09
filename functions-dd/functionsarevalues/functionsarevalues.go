package functionsarevalues

import "fmt"

type transformFn func(int) int

func main() {
	fmt.Println("Hello World")
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 45, 300}

	fmt.Println(transformNumbers(&numbers, triple), transformNumbers(&numbers, double))

	tffunc1 := getTranformFunction(&numbers)
	tffunc2 := getTranformFunction(&moreNumbers)

	fmt.Println(transformNumbers(&numbers, tffunc2), transformNumbers(&moreNumbers, tffunc1))

}

func transformNumbers(nums *[]int, transformation transformFn) (dubs []int) {
	for _, num := range *nums {
		dubs = append(dubs, transformation(num))
	}
	return dubs
}

func getTranformFunction(nums *[]int) (tf transformFn) {
	if (*nums)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
