package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var robot *Robot

	reader := bufio.NewReader(os.Stdin)
	tokenizer := regexp.MustCompile(`(,| )+`)

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		tokens := tokenizer.Split(strings.TrimRight(text, "\r\n"), -1)
		newRobot, err := Execute(tokens, robot)

		if newRobot != nil {
			robot = newRobot

			fmt.Printf("%v\n", robot)
		} else {
			fmt.Printf("%v\n", err)
		}
	}
}
