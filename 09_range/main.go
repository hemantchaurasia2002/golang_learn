package main

import "fmt"

func main() {
	ids := []int{36, 39, 42, 46, 50}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID - %d:\n", id)
	}

	//add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"phoenix": "phoenix@gmail.com", "skye": "skye@gmail.com", "brim": "brim@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
