package finitesets

import "fmt"

/*
 * Finite sets aren't a built-in like they are in Python. Sets are an incredibly useful
 * data type that allows you to add things to a list, _but not if it already exists in
 * the list_. They are un-ordered unlike an array, and are functionally easy to use once
 * you have it built. Here's how to make them.
 */

// This object is not relating to structs directly, I am only doing this for reference later.
type Coord struct {
	X, Y int
}

func main() {
	// The easiest (and smallest in memory) way to create a set is to use a map object
	// and set the values equal to an empty struct. You can set it to a boolean object
	// if you want, but each boolean object takes up one byte. Why use one byte when
	// you could use no bytes?

	// This sets up a composite literal struct object. Essentially the first set of
	// brackets state that it is an actionable struct object, and the second set of
	// brackets state that the struct object consists of literally nothing. Note that
	// we're declaring this variable so no need to infer a datatype.
	var exists = struct{}{}

	// Now we'll create a set object by making a map object of whatever type of data
	// we want to add to the set (in this case I'll choose an integer, but you can
	// literally use any valid Go object, even another struct object).
	mySet := make(map[int]struct{})

	// Now we can add to the set by setting any value equal to the `exists` variable
	// we created above.
	mySet[50] = exists
	mySet[100] = exists
	mySet[9001] = exists
	mySet[100] = exists

	// The above won't error out even though we already created an object with the value
	// of 100. We can confirm it exists with a quick 'if' statement.
	if _, ok := mySet[100]; ok {
		fmt.Println("100 exists in the set!")
	}

	// As I said before, we're not limited to common datatypes for items in the set. We can
	// even use structs!
	myCoordSet := make(map[Coord]struct{})

	myCoordSet[Coord{50, 100}] = exists
	myCoordSet[Coord{45, 20}] = exists
	myCoordSet[Coord{50, 100}] = exists

	if _, ok := myCoordSet[Coord{50, 100}]; ok {
		fmt.Println("The coordinates of X: 50, Y: 100 exist!")
	}

	// You can even loop through the values using a for loop.
	for k := range mySet {
		fmt.Printf("Value: %d\n", k)
	}
}
