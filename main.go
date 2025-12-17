//package main
//
//import "fmt"
//
//func main() {
//	var users = map[string]int{}
//
//	users["saleh"] = 2
//
//	fmt.Println(users)
//
//	for key, value := range users {
//		fmt.Println(key, value)
//	}
//
//	var u = map[int]int{}
//
//	fmt.Println(u)
//	fmt.Println(u[3])
//
//	// ok paradigm
//	value, ok := u[3]
//	if !ok {
//		fmt.Println("there is no value of key 3")
//	} else {
//		fmt.Println("value:", value)
//	}
//
//	type user struct {
//		ID    int
//		Name  string
//		Email string
//	}
//
//	var userList = make(map[uint]user)
//
//	userList[1] = user{
//		ID:    1,
//		Name:  "saleh",
//		Email: "sa@Sa",
//	}
//	fmt.Println("userList[1]", userList[1])
//
//	var userL = make(map[user]int)
//
//	userL[user{
//		ID:    1,
//		Name:  "saleh",
//		Email: "sa@Sa",
//	}] = 100
//
//	v1, ok := userL[user{
//		ID: 1,
//	}]
//	fmt.Println("test userL 1", v1, ok)
//
//	v2, ok := userL[user{
//		ID:   1,
//		Name: "saleh",
//	}]
//	fmt.Println("test userL 2", v2, ok)
//
//	v3, ok := userL[user{
//		ID:    1,
//		Name:  "saleh",
//		Email: "sa@Sa",
//	}]
//	fmt.Println("test userL 3", v3, ok)
//}

// package main
//
// import "fmt"
//
// func main() {
//
//	var u1 User
//	fmt.Println("u1", u1)
//	u1.Email = "saleh@.com"
//	fmt.Println(u1.Email)
//
//	u1.pr()
//
//	pr("saleh")
//
// }
//
//	type User struct {
//		ID       uint
//		Name     string
//		Email    string
//		IsActive bool
//	}
//
//	func (u User) pr(name string) {
//		fmt.Println(u.Name)
//	}
//
//	func pr(name string) {
//		fmt.Println(name)
//	}
package main

import (
	"errors"
	"fmt"
	"net/netip"
	"time"
)

func main() {
	var t1 = Teacher{
		ID:       2,
		IsActive: true,
	}
	t1.Deactivate()

	DeactivateUser(&t1, nil, nil)
	DeactivateUser2(&t1)

	s := Student{
		ID:       3,
		Name:     "",
		Email:    "",
		IsActive: false,
		JoinDate: time.Time{},
	}
	DeactivateUser2(&s)
	DeactivateUser(nil, &s, nil)
	// copy u2 = u1
}

type Teacher struct {
	ID       uint
	Name     string
	Email    string
	IsActive bool
	Salary   uint
	Grade    string
}

func (t *Teacher) Deactivate() error {

	// if !t.IsActive == false
	if !t.IsActive {
		return errors.New("Teacher is already deactivated")
	}

	if t.IsActive {
		t.IsActive = false
	}

	return nil
}

type Student struct {
	ID       uint
	Name     string
	Email    string
	IsActive bool
	JoinDate time.Time
}

func (s *Student) Deactivate() error {
	// if s.IsActive == false
	if !s.IsActive {
		return errors.New("the student is Deactivated already")
	}

	if s.IsActive {
		s.IsActive = false
	}
	return nil
}

type Staff struct {
	ID     uint
	Name   string
	Status string
}

func (s *Staff) Deactivate() {

	if s.Status != "active" {
		return errors.New("the student ")
	}

	s.Status = "deactive"

}

func DeactivateUser(t *Teacher, s *Student) error {
	if t != nil {
		// directly return t.Deactivate() error
		return t.Deactivate()

		// return new error
		//err := t.Deactivate()
		//if err != nil {
		//	errors.New()
		//}

		// wrap error
		error := t.Deactivate()
		if err != nil {
			return fmt.Errorf("can't deactivate user: %w", err)
		}
	}

	if s != nil {
		return s.Deactivate()
	}

	if Staff != nil {
		return Staff.Deactivate()
	}
	return nil
}

type User interface {
	Deactivate() error
}

func DeactivateUser2(u User) error {
	err := u.Deactivate()
	return RichError{
		Operation: "DeactivateUser2",
		Message:   err.Error(),
	}
	return rError(err.Error())
}

type RichError struct {
	Operation string
	Message   string
}

func (r RichError) Error() string {
	return r.Message
}

// concrete type
type rError string

func (r rError) Error() string {
	return string(r)
}
