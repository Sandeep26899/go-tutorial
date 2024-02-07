package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println("presentTime", presentTime)
	fmt.Println("Formated presentTime", presentTime.Format("01-02-2006 15:04:06 Monday"))

	createdDate := time.Date(2003, time.June, 2, 12, 0, 0, 0, time.Local)

	fmt.Println("createdDate", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
