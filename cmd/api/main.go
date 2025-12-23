package main

import "fmt"

func main() {
	message := "my first go program"
	age := 1

	fmt.Printf("%s %d\n", message, age)

	if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}
}
