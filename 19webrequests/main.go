package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	defer response.Body.Close() // We should close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
