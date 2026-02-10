package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user "
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")

	// Comma ok / comma err

	rating, _ := reader.ReadString(('\n'))     // Reads till next line -> we are defining \n to tell
	fmt.Println("Thanks for rating",rating)
	fmt.Printf("Type is %T \n", rating)

}
