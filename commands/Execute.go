package commands

import (
	"errors"

	"../robots"
)

func Execute(tokens []string, robot *robots.Robot) (*robots.Robot, error) {
	switch tokens[0] {
	case "place":
		return CommandPlace(tokens)
	case "move":
		return CommandMove(tokens, robot)
	default:
		return nil, errors.New("invalid command")
	}
}
