package main

import (
	"errors"
	"strconv"
)

func CommandPlace(tokens []string) (*Robot, error) {
	x, err := strconv.ParseInt(tokens[1], 0, 64)

	if err != nil {
		return nil, errors.New("invalid x")
	}

	y, err := strconv.ParseInt(tokens[2], 0, 64)

	if err != nil {
		return nil, errors.New("invalid y")
	}

	direction, err := StringToDirection(tokens[3])

	if err != nil {
		return nil, errors.New("invalid direction")
	}

	return &Robot{x, y, *direction}, nil
}
