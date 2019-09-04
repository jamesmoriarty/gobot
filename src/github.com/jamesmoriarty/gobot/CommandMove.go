package main

import "errors"

func CommandMove(tokens []string, robot *Robot) (*Robot, error) {
	switch robot.direction {
	case North:
		newRobot := Robot{robot.x, robot.y + 1, robot.direction}

		return &newRobot, nil
	case East:
		newRobot := Robot{robot.x + 1, robot.y, robot.direction}

		return &newRobot, nil
	case South:
		newRobot := Robot{robot.x, robot.y - 1, robot.direction}

		return &newRobot, nil
	case West:
		newRobot := Robot{robot.x - 1, robot.y, robot.direction}

		return &newRobot, nil
	default:
		return nil, errors.New("robot not placed")
	}

	return nil, nil
}
