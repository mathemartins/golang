package basic

// Struct is the combination of class and dict in both python and Javascript
// is a collection type - collection of key value pair
// struct keys can hold any type

type Wheel struct {
	Radius int
	Material string
}

type car struct {
	Make string
	Model string
	Height int
	Width int
	FrontWheel Wheel
	BackWheel Wheel
}

// instantiating the struct
myCar := car{}
myCar.FrontWheel.Radius = 5

// Anonymous structs are structs that don't have a name
// and are to be instantiated immediately they are created
// using a second pair of curly brackets.
// only reason to use an anonymous struct is if you choose not
// to call it more than once

myCar := struct {
	Make string
	Model string
} {
	Make: "tesla"
	Model: "model y"
}


// Embedded structs are structs that are completely shoved
// Inside another struct and all it's field are now top level fields
// of the structs it's inside

type vehicle struct {
	make string
	model string
}

type truck struct {
	vehicle
	bedSize int
}

// all the fields of vehicle are easily accessible on one-level access
// in the truck class
// instantiating the truck but still accessed one-level dot access

lanesTruck := truck {
	bedSize: 10,
	car: vehicle {
		make: "toyota",
		model: "camry"
	}
}

// struct methods, traditionally go is not object-oriented but it uses
// other special ways to achieve the same result as object-oriented programming
// one of the ways it does this is having struct methods
// struct methods are basically functions that receives special argument before the function name
// to represent the struct they are connected to

type rect struct {
	width int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

// instantiating the rect struct
r := rect {
	width: 5,
	height: 10,
}

// since r is an instance of the struct, therefore it can access the struct method
fmt.Println(r.area())
