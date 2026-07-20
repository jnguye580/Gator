package main

import (
	"errors"

	"github.com/jnguye580/GATOR-PROJECT/internal/config"
	"github.com/jnguye580/GATOR-PROJECT/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("Handler doesnt exist, try again")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
