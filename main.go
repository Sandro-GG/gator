package main

import (
	"log"
	"os"

	"github.com/Sandro-GG/gator/internal/config"
)

func main() {
	content, err := config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	currentState := NewState()
	currentState.cfg = &content

	commands := NewCommands()
	commands.commandList = make(map[string]func(*state, command) error)

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("please provide at least one argument")
	}

	cmd := NewCommand()
	cmd.name = os.Args[1]
	cmd.args = os.Args[2:]

	err = commands.run(&currentState, cmd)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
