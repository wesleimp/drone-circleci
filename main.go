package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "drone-circleci",
		Usage:     "Drone plugin for trigger circle-ci builds",
		Copyright: "Copyright (c) Weslei Juan Moser Pereira",
		Authors: []*cli.Author{
			{
				Name:  "Weslei Juan Moser Pereira",
				Email: "wesleimsr@gmail.com",
			},
		},
		Action: run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Usage:   "Access token",
				EnvVars: []string{"TOKEN"},
			},
			&cli.StringFlag{
				Name:    "user",
				Usage:   "Username",
				EnvVars: []string{"USER"},
			},
			&cli.StringFlag{
				Name:    "repo",
				Usage:   "Repository name",
				EnvVars: []string{"REPO"},
			},
			&cli.StringFlag{
				Name:    "branch",
				Usage:   "Branch to build",
				EnvVars: []string{"BRANCH"},
				Value:   "master",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	p := Plugin{
		Token:  c.String("token"),
		User:   c.String("user"),
		Repo:   c.String("repo"),
		Branch: c.String("branch"),
	}

	return p.Exec()
}
