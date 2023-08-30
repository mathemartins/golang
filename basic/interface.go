package basic

// interfaces is the collection of struct methods
// interfaces allows you to categorize various structs as one global name
// when a function collects interface as its argument,
// the interface is can allow any of the struct instance be passed as argument

// for a struct type to implement an interface it just needs to have all the methods of that interface

type employee interface {
	// method1
	// method2
}

// Copier interface with methods that takes an argument

type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
