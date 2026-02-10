package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["v"]) // store the key value pairs and return the values

	for i, val := range qparams {
		fmt.Printf("%v : %v\n", i, val)
	}

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "Ico.dev",
		Path:    "/tutess",
		RawPath: "user=hitesh",
	}
	anotherURL := partsofurl.String()
	fmt.Println(anotherURL)

}
