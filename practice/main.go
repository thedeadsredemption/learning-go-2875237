package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()

	fmt.Println("I did this at: ", n)

	t := time.Date(2003, time.March, 22, 23, 0, 0, 0, time.UTC)

	fmt.Println("I was born at: ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Sat Mar 22 23:00:00 2003")
	fmt.Printf("The type of parsed time is %T\n", parsedTime)

}
