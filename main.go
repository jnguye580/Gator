package main

import (
	"log"
	"os"

	"github.com/jnguye580/GATOR-PROJECT/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	st := state{
		config: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Fatal error, program has no command")
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := command{
		name: cmdName,
		args: args,
	}

	if err := cmds.run(&st, cmd); err != nil {
		log.Fatalf("Fatal error has occured: %v", err)
	}
}
