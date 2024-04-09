package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "billybob"

	userNames[1] = "william mcdinkle"

	userNames = append(userNames, "Anthony")

	userNames = append(userNames, "Retardo Mcclean")

	fmt.Println(userNames)

	courses := make(floatMap, 2)

	//courses := map[string]float64{}

	courses["go"] = 4.7
	courses["react"] = 4.8
	courses["angular"] = 4.9

	courses.output()

	for i, course := range courses {
		fmt.Println(i, course)
	}

}
