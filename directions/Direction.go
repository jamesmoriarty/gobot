package directions

import "errors"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func StringToDirection(s string) (direction *Direction, err error) {
	directions := []Direction{North, East, South, West}

	for _, direction := range directions {
		if direction.String() == s {
			return &direction, nil
		}
	}

	return nil, errors.New("invalid direction")
}
