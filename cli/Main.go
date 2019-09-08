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

const HELP = `
COMMANDS:

  - PLACE X<-N..N> Y<-N..N> Direction<North|East|South|West>
  - MOVE
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
