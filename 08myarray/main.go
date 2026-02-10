package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Fruit list is", len(fruitList))

	var vegList =[3]string{"Potato","Aloo","beans"}
	fmt.Println(vegList)
}
