package main

import (
	"errors"
	"strconv"
)

func CommandPlace(tokens []string) (*Robot, error) {
	x, _ := strconv.ParseInt(tokens[1], 0, 64)
	y, _ := strconv.ParseInt(tokens[2], 0, 64)
	direction, err := StringToDirection(tokens[3])
	if err != nil {
		return nil, errors.New("invalid direction")
	}
	return &Robot{x, y, *direction}, nil
}
