package main

import (
	"fmt"
	"maze/mazemap"
	"maze/readmaze"
)

func main() {
	fmt.Printf("Here is where the code will go.\n")

	thisMaze := readmaze.Ingest("./maze/maze.txt")

	m := mazemap.NewMaze(thisMaze)

	m.PrintMap()
}
