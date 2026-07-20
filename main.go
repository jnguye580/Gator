package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jnguye580/GATOR-PROJECT/internal/config"
	"github.com/jnguye580/GATOR-PROJECT/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	st := state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerList)

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
