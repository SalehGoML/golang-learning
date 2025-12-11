package main

import "fmt"

func main() {
	var users = map[string]int{}

	users["saleh"] = 2

	fmt.Println(users)

	for key, value := range users {
		fmt.Println(key, value)
	}
}
