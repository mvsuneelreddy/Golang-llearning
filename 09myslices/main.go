package main

import (
	"fmt"
	"sort"
)

func main() {
	var vegList = []string{"Potato", "Aloo", "beans"}
	fmt.Printf("Type is %T", vegList)

	vegList = append(vegList, "tomato", "brinjal")
	fmt.Println("\n", vegList)
	fmt.Println(vegList[1:])
	fmt.Println(vegList[1:3])

	highscores := make([]int, 4)
	highscores[0] = 90
	highscores[1] = 99
	highscores[2] = 98
	highscores[3] = 94
	// highscores[4] = 93

	highscores = append(highscores, 93, 90, 91)

	sort.Ints(highscores)

	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
