package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Sandro-GG/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username provided")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatalf("user %q doesn't exist", cmd.args[0])
	}

	_, err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	fmt.Printf("The username %s has been set", cmd.args[0])

	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no name provided")
	}

	username := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		log.Fatalf("user %q already exists", username)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	_, err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	s.cfg.CurrentUsername = user.Name
	fmt.Printf("New user has been created\nID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName:%s", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
}
