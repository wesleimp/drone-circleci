package main

import (
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
			},
		},
		Action: run,
	}
}

type Plugin struct {
	Token  string
	User   string
	Repo   string
	Branch string
}

func (p *Plugin) Exec() error {
	return nil
}

func run(c *cli.Context) error {
	p := &Plugin{
		Token:  c.String("token"),
		User:   c.String("User"),
		Repo:   c.String("Repo"),
		Branch: c.String("branch"),
	}

	return p.Exec()
}

// client := &circleci.Client{Token: "0b141a931d7563f426c70d07c731de9b8f8ee6b1"}

// builds, _ := client.ListRecentBuildsForProject("wesleimp", "Traveller", "master", "", -1, 0)

// for _, build := range builds {
// fmt.Printf("%v - %d: %s\n", build.Reponame, build.BuildNum, build.Status)
// }
