package mazemap

// This package will be the character that we move around.

// This handy enum will be useful for later
const (
	North = iota
	South
	East
	West
)

// technically, all directions are mapped to literally 0, 1, 2, and 3.
// But we can refer to them as North South East and West.

type Player struct {
	location Coord
}

func NewPlayer(c Coord) Player {
	return Player{
		location: c,
	}
}

// Useful to have a move function
func (p *Player) Move(c Coord) {
	p.location = c
}

// This will return 4 coordinates in all directions from the player.
// Considering this isn't even using the map context here, best bet
// is to take the results of these and confirm if the
func (p Player) PotentialMoves() [4]Coord {
	// The handy thing is I can even set North/South/East/West as the
	// index of the return value!
	var retval [4]Coord
	retval[North] = Coord{
		X: p.location.X,
		Y: p.location.Y - 1,
	}

	retval[South] = Coord{
		X: p.location.X,
		Y: p.location.Y + 1,
	}

	retval[East] = Coord{
		X: p.location.X + 1,
		Y: p.location.Y,
	}

	retval[West] = Coord{
		X: p.location.X - 1,
		Y: p.location.Y,
	}

	return retval
}
