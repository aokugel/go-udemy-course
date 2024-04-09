package lists

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	prices := []float64{10.12, 211.11}

	fmt.Println(prices[1:])

	prices = append(prices, 5.65, 12.99)

	discountPrices := []float64{1.12, 3.40, 4.22}

	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}

// func main() {
// 	var hobbies = [3]string{"Running", "traveling", "Lifting"}
// 	fmt.Println(hobbies, hobbies[0], hobbies[1:3])

// 	//hobby_slice1 := hobbies[:2]

// 	hobby_slice2 := hobbies[0:2]

// 	hobby_slice2 = hobby_slice2[1:3]

// 	fmt.Println(hobby_slice2)

// 	goals := []string{"be profficient in go", "be able to right backend services and operators in go. "}

// 	goals[1] = "Get a fat sweet gig"

// 	goals = append(goals, "Make paper")

// 	var products = []Product{
// 		{
// 			title: "allen wrench",
// 			id:    1,
// 			price: 12.24,
// 		},
// 		{
// 			title: "Lemon drop",
// 			id:    2,
// 			price: .24,
// 		},
// 	}
// 	products = append(products, Product{
// 		title: "Bobby boshay",
// 		id:    3,
// 		price: 45.24,
// 	})

// 	// products[0] = Product{
// 	// 	title: "allen wrench",
// 	// 	id:    1,
// 	// 	price: 12.24,
// 	// }

// 	fmt.Println(products)

// }
