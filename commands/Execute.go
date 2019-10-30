package commands

import (
	"errors"

	"github.com/jamesmoriarty/gobot/robots"
)

func Execute(tokens []string, robot *robots.Robot) (*robots.Robot, error) {
	switch tokens[0] {
	case "Place":
		return CommandPlace(tokens)
	case "Move":
		return CommandMove(tokens, robot)
	case "Left":
		return CommandLeft(tokens, robot)
	case "Right":
		return nil, errors.New("command not implemented")
	default:
		return nil, errors.New("invalid command")
	}
}
