package main

import "fmt"

func main() {
	reddy := User{"Reddy", "reddy@gmail.com", 21, true}
	fmt.Println(reddy)
	fmt.Printf("reddy details are: %+v\n", reddy)
	fmt.Printf("Reddy name is: %v\n", reddy.Name)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
