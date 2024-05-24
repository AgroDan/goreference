package structs

import "strings"

type Animal struct {
	name                string
	age, height, weight int
	carrying            []string
}

// The above is a pretty contrived object that has some attributes we can work with.
// Now I can create functions that are owned by this object!

// First it's pretty useful to create a constructor object
func NewAnimal(name string, age, height, weight int) Animal {
	// We will use this constructor to ensure that whoever
	// calls this struct object creates the object properly
	// and doesn't leave anything blank. Also this would be
	// a great place to make a map object if this requires one
	return Animal{
		name:     name,
		age:      age,
		height:   height,
		weight:   weight,
		carrying: []string{},
	}
}

// None of the variables within the Animal struct begin with a capital
// letter, which means that even if we import this module, we won't be
// able to reference those variables directly. An easy fix is to just
// change the first character of each variable to a capital letter,
// but another is to create functions that return the values of those
// variables. This is probably more preferable since you can always
// go back later and change the function to output a slightly modified
// version (such as a printer-friendly one or something) without allowing
// the caller to reference the variable directly. Example here:

func (a Animal) GetName() string {
	return strings.ToUpper(a.name)
}

// For the above function, I am passing by value for the object itself,
// meaning that the object will return (in essence) a _copy_ of the object.
// This is fine for most situations, but if you ever need to modify a value
// of the object itself you need to pass _by reference_. This is denoted
// by adding a * character next to the name of the struct object, in this case
// *Animal like so:

func (a *Animal) SetAge(ageValue int) {
	a.age = ageValue
}

// If you didn't pass by reference, this would have no effect since it would
// modify the value of a copy of the object and not the object itself.
