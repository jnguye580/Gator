package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/jnguye580/GATOR-PROJECT/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Empty argument slice")
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		log.Fatalf("Fatal error has occured: %v", err)
	}

	if err := s.config.SetUser(username); err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Empty argument slice")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		log.Fatalf("Fatal error has occured: %v", err)
	}

	if err := s.config.SetUser(user.Name); err != nil {
		log.Fatalf("Fatal error has occured: %v", err)
	}
	log.Printf("%+v\n", user)

	return nil
}
