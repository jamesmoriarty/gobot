package commands

import (
	"errors"

	"github.com/jamesmoriarty/gobot/directions"
	"github.com/jamesmoriarty/gobot/robots"
)

func CommandMove(tokens []string, robot *robots.Robot) (*robots.Robot, error) {
	if robot == nil {
		return nil, errors.New("robot not placed")
	}

	x, y := getOffset(robot.Direction)

	return &robots.Robot{robot.X + x, robot.Y + y, robot.Direction}, nil
}

func getOffset(d directions.Direction) (int64, int64) {
	switch d {
	case directions.North:
		return 0, 1
	case directions.East:
		return 1, 0
	case directions.South:
		return 0, -1
	case directions.West:
		return -1, 0
	default:
		return 0, 0
	}
}
