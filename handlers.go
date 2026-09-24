package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username provided")
	}

	_, err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	fmt.Printf("The username %s has been set", cmd.args[0])

	return nil
}
