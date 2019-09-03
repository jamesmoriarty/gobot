package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Execute(tokens []string, robot *Robot) (*Robot, error) {
	switch tokens[0] {
	case "place":
		return CommandPlace(tokens)
	default:
		return nil, errors.New("invalid command")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tokenizer := regexp.MustCompile(`(,| )+`)

	var (
		robot *Robot
		err   error
	)

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')

		tokens := tokenizer.Split(strings.TrimRight(text, "\r\n"), -1)

		robot, err = Execute(tokens, robot)

		fmt.Printf("%v, %v\n", robot, err)
	}
}
