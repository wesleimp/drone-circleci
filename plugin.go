package main

import (
	"errors"
	"log"

	"github.com/jszwedko/go-circleci"
)

type (
	//Plugin values
	Plugin struct {
		Token  string
		User   string
		Repo   string
		Branch string
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	log.Println(p.User, p.Repo, p.Token, p.Branch)
	if len(p.Token) == 0 || len(p.User) == 0 || len(p.Repo) == 0 {
		return errors.New("Missing circle-ci configs")
	}

	c := &circleci.Client{Token: p.Token}

	_, err := c.Build(p.User, p.Repo, p.Branch)
	if err != nil {
		return err
	}

	return nil
}
