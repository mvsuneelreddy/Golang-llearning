package main

import (
	"fmt"
	"math/rand"
)

func main() {

	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("The number is 1")
	case 2:
		fmt.Println("The number is 1")
	case 3:
		fmt.Println("The number is 1")
		fallthrough
	case 4:
		fmt.Println("The number is 1")
		fallthrough							// fallthrough makes the next line also to print
	case 5:
		fmt.Println("The number is 1")
	case 6:
		fmt.Println("The number is 1")

	}
}
