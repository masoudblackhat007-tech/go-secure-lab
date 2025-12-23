package main

import (
	"fmt"
	"strconv"
)

func main() {
	var message string = "my first go program"
	age := 1

	ageText := strconv.Itoa(age)
	fmt.Println(message + " " + ageText)
}
