package main

import (
	"fmt"
)

const aConst int = 45

func main() {
	var aString string = "This is Tobi lol"
	fmt.Println(aString)
	fmt.Printf("The variable is %T\n", aString)

	var anInteger int = 98
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString string = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable is %T\n", anotherString)

	var anotherInt int = 25
	fmt.Println(anotherInt)
	fmt.Printf("The variable is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variable is %T\n", aConst)

}
