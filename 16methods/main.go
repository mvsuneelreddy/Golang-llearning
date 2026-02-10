package main

import "fmt"

func main() {
	reddy := User{"Reddy", "reddy@gmail.com", 21, true}
	fmt.Println(reddy)
	fmt.Printf("reddy details are: %+v\n", reddy)
	fmt.Printf("Reddy name is: %v\n", reddy.Name)
	reddy.GetStatus()
	reddy.NewMail()
	fmt.Printf("Reddy name is: %v\n", reddy.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println(u.Status)
}

func (u User) NewMail() {
	u.Email = "newmail@go.com"   // This will not change the actual object as it is just a copy of User we are passing
	fmt.Println(u.Email)
}
