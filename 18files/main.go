package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is the content to be written into the file "
	file, err := os.Create("./mylog.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println(length)
	defer file.Close()
	readfile("./mylog.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println(string(databyte))
}

// To save repetition
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
