package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("%v is %v\n", i, day)
	// }

	i := 1

	for i < 10 {

		if(i==2){
			goto hey
		}
		if(i ==8){
			break
		}

		fmt.Println(i)
		i++
	}

	hey :
	fmt.Println("This is a goto / jump statement")

}
