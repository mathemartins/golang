package logic

import "fmt"

// Password validator, contain Caps 1 atleast, Lower 1 atleast, 1 digit atleast, len atleast 8
// atleast 1 symbol
func validatePass(password string) bool {
	// collect the password,
	// loop through and check for 1 capital letter
	for i := 0; i < len(password); i++ {
		fmt.Println(i)
	}
	return true
}

func main() {
	validatePass("pass=123")
}
