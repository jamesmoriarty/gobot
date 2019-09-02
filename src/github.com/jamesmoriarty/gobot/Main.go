package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func place(tokens []string) (*Robot, error) {
	x, _ := strconv.ParseInt(tokens[1], 0, 32)
	y, _ := strconv.ParseInt(tokens[2], 0, 32)
	direction, err := StringToDirection(tokens[3])
	if err != nil {
		return nil, errors.New("invalid direction")
	}
	return &Robot{x, y, *direction}, nil
}

func execute(tokens []string, robot *Robot) (*Robot, error) {
	switch tokens[0] {
	case "place":
		return place(tokens)
	default:
		return nil, errors.New("invalid command")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tokenizer := regexp.MustCompile(`(,| )+`)

	var robot *Robot
	var err error

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')

		tokens := tokenizer.Split(strings.TrimRight(text, "\r\n"), -1)

		robot, err = execute(tokens, robot)

		fmt.Printf("%q, %q\n", *robot, err)
	}
}
