package maps

import (
	"fmt"
)

func main() {
	websites := []string{"google.com", "aws.com"}
	webMap := map[string]string{
		"google": "https://google.com",
		"reddit": "https://reddit.com",
	}

	fmt.Println(webMap["google"], websites[1])

	webMap["instragram"] = "https://instagram.com"
	fmt.Println(webMap)

	delete(webMap, "reddit")

	fmt.Println(webMap)
}
