package basic

import (
	"fmt"
)

// Syntaxing in golang
//length := getLength("email")
//if length < 1 {
//	fmt.Println("Email is invalid")
//}
//
//if length := getLength("email") < 1 {
//	fmt.Println("Email is invalid")
//}

func main() {
	var username string = "alphador@"
	name := "tomide"
	age := 23
	var password string = "password"

	// String interpolation
	// Use the %v for default value string substitution
	// Use the Sprintf() to format your string and return the value to an assigned variable
	msg := fmt.Sprintf("Hi my name is %s and I am %d years old", name, age)
	fmt.Println("Authorization Bearer: ", username+":"+password+":"+name+":"+msg)
}
