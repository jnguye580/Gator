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

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Fatal error has occured: %v", err)
	}
	fmt.Println("Success, Database has been reset")
	return nil
}

func handlerList(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Fatal error has occured: %v", err)
	}
	for _, u := range users {
		if u.Name == s.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("fetching feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("Empty argument slice")
	}

	userData, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Unable to get user record: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    userData.ID,
	})
	if err != nil {
		return fmt.Errorf("Unable to create feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userData.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Unable to create feed follow: %w", err)
	}

	fmt.Printf("%+v\n", feed)
	fmt.Printf("%+v\n", feedFollow)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to obtain feeds: %v", err)
	}
	for _, f := range feeds {
		fmt.Printf("%s\n", f.Name)
		fmt.Printf("%s\n", f.Url)
		fmt.Printf("%s\n", f.Username)
	}
	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("Invalid argument slice")
	}

	URL := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), URL)
	if err != nil {
		return fmt.Errorf("Unable to obtain feed: %v", err)
	}

	userData, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Unable to get user record: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userData.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Unable to create feed follow: %w", err)
	}

	fmt.Printf("%v\n", feedFollow.FeedName)
	fmt.Printf("%v\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("Invalid argument slice")
	}

	userData, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Unable to get user record: %w", err)
	}

	userFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), userData.ID)
	if err != nil {
		return fmt.Errorf("Unable to get user following data: %w", err)
	}

	for _, follows := range userFollowing {
		fmt.Printf("%s\n", follows.FeedName)
	}
	return nil
}
