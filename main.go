package main

import "fmt"

func main() {
	var users = map[string]int{}

	users["saleh"] = 2

	fmt.Println(users)

	for key, value := range users {
		fmt.Println(key, value)
	}

	var u = map[int]int{}

	fmt.Println(u)
	fmt.Println(u[3])

	// ok paradigm
	value, ok := u[3]
	if !ok {
		fmt.Println("there is no value of key 3")
	} else {
		fmt.Println("value:", value)
	}

	type user struct {
		ID    int
		Name  string
		Email string
	}

	var userList = make(map[uint]user)

	userList[1] = user{
		ID:    1,
		Name:  "saleh",
		Email: "sa@Sa",
	}
	fmt.Println("userList[1]", userList[1])

	var userL = make(map[user]int)

	userL[user{
		ID:    1,
		Name:  "saleh",
		Email: "sa@Sa",
	}] = 100

	v1, ok := userL[user{
		ID: 1,
	}]
	fmt.Println("test userL 1", v1, ok)

	v2, ok := userL[user{
		ID:   1,
		Name: "saleh",
	}]
	fmt.Println("test userL 2", v2, ok)

	v3, ok := userL[user{
		ID:    1,
		Name:  "saleh",
		Email: "sa@Sa",
	}]
	fmt.Println("test userL 3", v3, ok)
}
