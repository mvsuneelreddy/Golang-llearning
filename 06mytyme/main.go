package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 1, 12, 1, 1, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 15:04:05 Monday"))

}
