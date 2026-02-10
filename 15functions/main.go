package main

import "fmt"

func main() {
	fmt.Println("This is main")
	greet()
	result := adder(2, 3)
	fmt.Println(result)
	res, resmess := proAdder(1, 2, 3, 4, 5)

	fmt.Println(res)
	fmt.Println(resmess)

	fmt.Println(proAdder(1, 2, 3, 4, 5))
}

func greet() {
	fmt.Println("Greet function")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hello world"
}
