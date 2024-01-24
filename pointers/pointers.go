package main

import "fmt"

func main() {
	age := 32 // reg var

	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years:", *agePointer)
	fmt.Println("Age after mutation bullshit: ", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
	//return *age
}
