package main

import "fmt"

func main() {
	var username string = "Reddy"
	fmt.Println("Usernmae is:", username)
	fmt.Printf("Type of Variable is %T \n", username)

	var isLoggedin bool = true
	fmt.Println("Looged in :", isLoggedin)
	fmt.Printf("Type of Variable is %T \n", isLoggedin)

	var defval string
	fmt.Println("Deafault value is:", defval)

	//Otehr ways of Assigning

	var name = "Guest"
	fmt.Println("name:", name)

	anotehrname := "Noname for now"
	fmt.Println("Anotehrname is:", anotehrname)

	num := 100
	fmt.Println(num)

}
