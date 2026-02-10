package main

import "fmt"

func main() {
	var myPtr *int
	fmt.Println("Value is ", myPtr)

	mynum := 12
	numPtr := &mynum
	fmt.Println("Value is", numPtr)
	fmt.Println("Value is", *numPtr)

	*numPtr = *numPtr + 3
	fmt.Println("New value is", mynum)
}
