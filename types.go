package main

import (
	"fmt"

	"github.com/Sandro-GG/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func NewState() state {
	return state{}
}

type command struct {
	name string
	args []string
}

func NewCommand() command {
	return command{}
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func NewCommands() commands {
	return commands{}
}

func (c *commands) run(s *state, cmd command) error {
	funct, ok := c.commandList[cmd.name]
	if !ok {
		return fmt.Errorf("command %v doesn't exist", cmd.name)
	}

	err := funct(s, cmd)
	if err != nil {
		return fmt.Errorf("couldn't run %s, %w", cmd.name, err)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
}
