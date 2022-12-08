package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ineentho/aoc-2022/challenges"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Advent of Code 2022",
		Commands: []*cli.Command{
			{
				Name: "1",
				Subcommands: []*cli.Command{
					createChallengeCommand("1", challenges.Day01Part1),
					createChallengeCommand("2", challenges.Day01Part2),
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createChallengeCommand(name string, challenge func(input string) (string, error)) *cli.Command {
	return &cli.Command{
		Name: name,
		Action: func(ctx *cli.Context) error {
			bytes, err := ioutil.ReadAll(os.Stdin)

			if err != nil {
				return err
			}

			inputStr := strings.TrimSpace(string(bytes))

			result, err := challenge(inputStr)

			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}
}
