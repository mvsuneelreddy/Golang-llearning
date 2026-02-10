package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code ", response.StatusCode)
	fmt.Println(response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
