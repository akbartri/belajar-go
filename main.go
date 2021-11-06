package main

import "fmt"

func main() {

	var Age1 int16 = 2 //init var with types (capital in first letter means public, if not its private)
	var age = 2        //init var without types
	age2 := 2          //another way to init var

	fmt.Printf("age1 : %v\n", Age1)
	fmt.Printf("age : %v\n", age)
	fmt.Printf("age2 : %v\n", age2)
}
