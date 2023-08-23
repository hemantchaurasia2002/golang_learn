package main

import "fmt"

func main() {
	//Define map
	// emails := make(map[string]string)

	// //assign kv
	// emails["phoenix"] = "phoenix@gmail.com"
	// emails["skye"] = "skye@gmail.com"
	// emails["brim"] = "brime@gmail.com"

	emails := map[string]string{"phoenix": "phoenix@gmail.com", "skye": "skye@gmail.com", "brim": "brim@gmail.com"}

	fmt.Println(emails)
	//delete
	delete(emails, "brim")
	fmt.Println(emails)
}
