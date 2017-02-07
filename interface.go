package main

import "fmt"

type (
	person struct {
		num  int
		name string
	}

	place struct {
		num   int
		place string
	}
)

// Interface that requires structs to have a pointer receiver of 'WriteInt()'
type test interface {
	WriteInt()
}

// Write the int for the struct
func (p person) WriteInt() {
	fmt.Printf("Person int: %d\n", p.num)
}

// Write the int for the struct
func (p place) WriteInt() {
	fmt.Printf("Place int: %d\n", p.num)
}

// This uses the 'fmt' interface for 'String', so we get a string formated to
// our desire if we have a 'String' pointer receiver for our type.
func (p person) String() string {
	return fmt.Sprintf("The number is %d, and the name is %s", p.num, p.name)
}

func main() {

	// The interface requirement means that any struct added must have a pointer
	// receiver that has the WriteInt() function. This is the only requirement.
	var holding []test

	// Adding 'person', which has a WriteInt() pointer receiver.
	holding = append(holding, person{num: 4, name: "Bob"})

	// Adding 'place', which has a WriteInt() pointer receiver.
	holding = append(holding, place{num: 6, place: "here"})

	for _, v := range holding {
		v.WriteInt()
	}

	// We will get a nicely formatted string, because we satisfy the interface
	// requirement.
	fmt.Println(holding[0])

	// We get a printout of the object, because there is pointer receiver to
	// satisfy the interface
	fmt.Println(holding[1])
}
