package main

import (
	"fmt"
	"reflect"
)

type target struct {
	name string

	// This indicates that the data type, it DOES NOT initialize the map. If you
	// attempt to add something to this map without an allocation, you're going
	// to have a bad time.
	things map[string]string
}

// A common way of get an object with an initialized map
func newTarget() target {
	return target{things: make(map[string]string)}
}

func main() {

	// The long way of adding a map to a struct that contains a map
	init := make(map[string]string)
	first := target{things: init}
	first.things["foo"] = "bar"

	// The quick way of adding a map to a struct that contains a map
	second := newTarget()
	second.things["foo"] = "bar"

	// Deeply equal means that all elements of the struct must match.
	if reflect.DeepEqual(first, second) {
		fmt.Println("The two structs are equal.")
	} else {
		fmt.Println("The two structs are not equal.")
	}

}
