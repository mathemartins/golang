package main

import (
	"math"
)

//num := 5
//pi := 3.1456
//name := "Alice"
//isRaining := true
//var b byte = 65
//var r rune = '🌟'
//number_list := [5]int{1, 2, 3, 4, 5}
//fruits := []string{"Pawpaw", "Apple", "Mango"}
//ages := map[string]int{
//"Alice": 25,
//"Grace": 18,
//}
//type Person struct {
//	Name string
//	Age  int
//}
//numPointer := &num
//type Shape interface {
//	Area() float64
//}
//type Circle struct {
//	Radius float64
//}
//func (c Circle) Area() float64 {
//	return 3.14159 * c.Radius * c.Radius
//}

// basic print hello world
//func hello() {
//	fmt.Println("Hello World")
//}
//
//// for loop in golang
//func counter() int {
//	for i := 0; i < 10; i++ {
//		fmt.Println(i)
//	}
//	return 1
//
//}
//
//// control statements and loop in golang
//func selectionStatement() {
//	weather := "rainy"
//
//	if weather == "rainy" {
//		fmt.Println("Take umbrella")
//	} else {
//		fmt.Println("Go like that")
//	}
//}

// Add two numbers
//func addNumbers(number1 int, number2 int) int {
//	total := number1 + number2
//	return total
//}

// Password validator, contain Caps 1 atleast, Lower 1 atleast, 1 digit atleast, len atleast 8
// atleast 1 symbol
//func validatePass(password string) bool {
//	if len(password) < 8 {
//		return false
//	}
//	hasUppercase := false
//	hasLowercase := false
//	hasDigit := false
//
//	for _, char := range password {
//		if unicode.IsUpper(char) {
//			hasUppercase = true
//		} else if unicode.IsLower(char) {
//			hasLowercase = true
//		} else if unicode.IsDigit(char) {
//			hasDigit = true
//		}
//	}
//
//	return hasUppercase && hasLowercase && hasDigit
//}

//"""
//Pseudo Code:
//1. collect a list of numbers
//2. assign highest odd number as None for a start
//3. iterate through the list and check for odd condition while checking if highest number has been assigned a value
//4. Assign highest number a last value and loop again until loop is complete
//5. return highest number
//"""

func highestOddNumber(numberList []int) int {
	// get the number list
	highestOdd := math.MinInt
	// loop through the array
	for i := 0; i < len(numberList); i++ {
		if numberList[i]%2 != 0 {
			if numberList[i] > highestOdd {
				highestOdd = numberList[i]
			}
		}
	}
	return highestOdd
	// check for odd using modulo
	// assign to a variable
	// compare to the last number
	// return the highest one
}

// combine two list
func alternateCombineList(firstArray []interface{}, secondArray []interface{}) []interface{} {
	var combinedArray []interface{}
	for i := 0; i < len(secondArray); i++ {
		combinedArray = append(firstArray, secondArray[i])
	}
	return combinedArray
}

func main() {
	//numbers := []string{"2", "4", "7", "9", "12", "15", "18"}
	//fruits := []string{"Apple", "Mango", "Pawpaw", "Orange", "Pineapple"}
	//fmt.Println(alternateCombineList(numbers, fruits []interface{}))
	//numbers := []int{2, 4, 7, 9, 12, 15, 18}
	//highest := highestOddNumber(numbers)
	//fmt.Println("Highest odd number:", highest)
	//fmt.Println(highestOddNumber([]int{1, 2, 3, 4, 5}))
	//fmt.Println(validatePass("Mathemartins1@"))

	//// call the function
	//sumValue := addNumbers(2, 3)
	//// print the function
	//fmt.Println("Action: ", sumValue)
	//
	//// loop
	//outputLoop := counter()
	//fmt.Println("Output: ", outputLoop)
}
