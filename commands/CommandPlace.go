package commands

import (
	"errors"
	"strconv"

	"../directions"
	"../robots"
)

func CommandPlace(tokens []string) (*robots.Robot, error) {
	if len(tokens) != 4 {
		return nil, errors.New("invalid number of command arguments")
	}

	x, err := strconv.ParseInt(tokens[1], 0, 64)

	if err != nil {
		return nil, errors.New("invalid x")
	}

	y, err := strconv.ParseInt(tokens[2], 0, 64)

	if err != nil {
		return nil, errors.New("invalid y")
	}

	direction, err := directions.StringToDirection(tokens[3])

	if err != nil {
		return nil, errors.New("invalid direction")
	}

	return &robots.Robot{x, y, *direction}, nil
}
