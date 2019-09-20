package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"../commands"
	"../robots"
)

const HELP = `Commands:

  - Place X<Number> Y<Number> Direction<North|East|South|West>
  - Move
  - Left
  - Right
`

func main() {
	var robot *robots.Robot

	reader := bufio.NewReader(os.Stdin)
	tokenizer := regexp.MustCompile(`(,| )+`)

	fmt.Println(HELP)

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		tokens := tokenizer.Split(strings.TrimRight(text, "\r\n"), -1)
		newRobot, err := commands.Execute(tokens, robot)

		if newRobot != nil {
			robot = newRobot

			fmt.Printf("%v %v %v\n", robot.X, robot.Y, robot.Direction.String())
		} else {
			fmt.Printf("%v\n", err)
		}
	}
}
