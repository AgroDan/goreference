package main

/*
 * This module discusses how data is serialized with Go. If you are creating
 * the data object to be serialized, then it should be pretty straightforward
 * in serialization. Generally you don't need to be too specific, just create
 * the object and Marshall it using the json.Marshall() function. However, if
 * another endpoint creates the object and you want to create the Go struct
 * that the data will be deserialized into, you can create _field tags_ to
 * assist your code into how it should expect serialized data to appear. Let's
 * take a look.
 */

import (
	"encoding/json"
	"fmt"
	"log"
)

// With the below object, I am currently defining it as such. But if there is
// another source that will be providing me with a JSON object, I can define
// very specific field tag metadata to ensure that the object will translate
// well. For instance, if I were to take the below struct type and serialize
// it with the below example, it will look like this:
//
// {"name":"Contra","year":1987,"genre":"Run and Gun"}
//
// But if I didn't have the field tags, it would look like this:
// {"Name":"Contra","YearCreated":1987,"Genre":"Run and Gun"}
//
// Note that the names of each attribute are different! This is useful if you
// are expecting data to be named something else.
//
// Similarly, you can use the same method to marshall and unmarshall XML data
// as well, and even gRPC data!

type Game struct {
	Name        string `json:"name"`
	YearCreated int    `json:"year"`
	Genre       string `json:"genre"`
}

func main() {
	myGame := Game{
		Name:        "Contra",
		YearCreated: 1987,
		Genre:       "Run and Gun",
	}

	// Let's create the JSON object here:
	jsonOut, err := json.Marshal(myGame)
	if err != nil {
		log.Fatal("Could not marshall")
	}
	fmt.Printf("JSON'd: %s\n", string(jsonOut))

	// Now let's take that JSON data and build ANOTHER object from it.
	var returnedGame Game

	// Note that the json.Unmarshall() function will modify the object
	// in place, so make sure you send the pointer to it.
	if err := json.Unmarshal(jsonOut, &returnedGame); err != nil {
		log.Fatal("Could not unmarshall")
	}

	fmt.Printf("In memory: %v+\n", returnedGame)

}
