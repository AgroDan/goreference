package mazemap

import "fmt"

/*
 *
 * The mazemap package will involve converting the data received from the readmaze
 * return value ([]string) into an object that we can work with. I will create a
 * struct object called a Coord{} which will be useful for referencing a specific
 * coordinate within the map. But keep in mind that due to the way that the data is
 * read (which is to say, left to right, top to bottom), the Y coordinates are read
 * from top down.
 */

// Don't need to associate a function here, this is just useful.
type Coord struct {
	X, Y int
}

type Maze struct {
	// I'll have a datatype to store the contents of the maze itself.
	contents [][]rune
	start    Coord
	end      Coord
}

// standard Constructor
func NewMaze(m []string) Maze {
	retval := Maze{}
	var mapVal [][]rune
	for Y, rVal := range m {
		// iterate over rows
		var rows []rune
		for X, cVal := range rVal {
			rows = append(rows, cVal)

			// check for start and end
			if cVal == 'S' {
				retval.start = Coord{
					X: X,
					Y: Y,
				}
			}
			if cVal == 'F' {
				retval.end = Coord{
					X: X,
					Y: Y,
				}
			}
		}
		mapVal = append(mapVal, rows)
	}
	retval.contents = mapVal
	return retval
}

// I typically will have a print function which dumps the contents of the map
// just to make sure I have it right
func (m Maze) PrintMap() {
	for Y := 0; Y < len(m.contents); Y++ {
		for X := 0; X < len(m.contents[Y]); X++ {
			fmt.Printf("%c", m.contents[Y][X])
		}
		fmt.Printf("\n")
	}
}

// Now I'll use a function to determine if a space is a valid walking space.
// This will be particularly useful as I look to create a moveable character.
// I can use a function that will check north, south, east and west to see
// if any of those are considered valid.
func (m Maze) ValidSpace(c Coord) bool {
	// This, given a coordinate, will determine
	// if the space is capable of walking onto.
	// first, let's do some bounds checking
	if c.X < 0 || c.X >= len(m.contents[0]) {
		// off the map
		return false
	}

	if c.Y < 0 || c.Y >= len(m.contents) {
		return false
	}

	if m.contents[c.Y][c.X] == '#' {
		return false
	}

	// all other positions are true.
	return true
}

// Now some standard return functions
func (m Maze) Starting() Coord {
	return m.start
}

func (m Maze) End() Coord {
	return m.end
}
