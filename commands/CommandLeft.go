package commands

import (
	"errors"

	"../robots"
)

func CommandLeft(tokens []string, robot *robots.Robot) (*robots.Robot, error) {
	if robot == nil {
		return nil, errors.New("robot not placed")
	}

	return &robots.Robot{robot.X, robot.Y, robot.Direction.Rotate(-1)}, nil
}
