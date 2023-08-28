package basic

// A function can have multiple return values
// and that's how golang handles its exceptions by throwing an error and a success value
// a function that has multiple return values, can specify its type by writing the return type
// in parenthesis (string string) for example.
// Go doesn't allow you to have unused variables so it requires you to remove it or make it unused using the
// underscore _ syntax

// Using named return values are great built-in mechanism
func getCord() (width, length, height int) {
	// automatically initialize the width, length and height to zero
	return // and returns them by default even if you don't put them here to be returned
}
