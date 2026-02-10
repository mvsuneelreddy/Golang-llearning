package main

import "fmt"

func main() {
	val := 25
	var ans string
	if val < 10 {
		ans = "less than 10"
	} else if val > 10 {
		ans = "greater than 10"
	} else {
		ans = "nothing"
	}

	fmt.Println(ans)

	// assigning and checking

	if num := 10; num >= 10 {
		fmt.Println("Num is greater or equal to 10")
	} else {
		fmt.Println("Num is less than 10")
	}

}
