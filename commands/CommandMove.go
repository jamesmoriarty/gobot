package commands

import (
	"errors"

	"../directions"
	"../robots"
)

func CommandMove(tokens []string, robot *robots.Robot) (*robots.Robot, error) {
	switch robot.Direction {
	case directions.North:
		newRobot := robots.Robot{robot.X, robot.Y + 1, robot.Direction}

		return &newRobot, nil
	case directions.East:
		newRobot := robots.Robot{robot.X + 1, robot.Y, robot.Direction}

		return &newRobot, nil
	case directions.South:
		newRobot := robots.Robot{robot.X, robot.Y - 1, robot.Direction}

		return &newRobot, nil
	case directions.West:
		newRobot := robots.Robot{robot.X - 1, robot.Y, robot.Direction}

		return &newRobot, nil
	default:
		return nil, errors.New("robot not placed")
	}
}
