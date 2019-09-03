package main

import "errors"

func Execute(tokens []string, robot *Robot) (*Robot, error) {
	switch tokens[0] {
	case "place":
		return CommandPlace(tokens)
	default:
		return nil, errors.New("invalid command")
	}
}
