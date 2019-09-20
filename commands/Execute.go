package commands

import (
	"errors"

	"../robots"
)

func Execute(tokens []string, robot *robots.Robot) (*robots.Robot, error) {
	switch tokens[0] {
	case "Place":
		return CommandPlace(tokens)
	case "Move":
		return CommandMove(tokens, robot)
	case "Left":
		return nil, errors.New("command not implemented")
	case "Right":
		return nil, errors.New("command not implemented")
	default:
		return nil, errors.New("invalid command")
	}
}
