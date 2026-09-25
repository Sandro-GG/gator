package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Sandro-GG/gator/internal/config"
	"github.com/Sandro-GG/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	// load config
	content, err := config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// initialize state
	currentState := NewState()
	currentState.cfg = &content

	// setup db connection
	db, err := sql.Open("postgres", content.DbUrl)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer db.Close()

	// wrap db with sqlc queries
	dbQueries := database.New(db)
	currentState.db = dbQueries

	// setup command registry map
	commands := NewCommands()
	commands.commandList = make(map[string]func(*state, command) error)

	// register all commands
	commands.register("login", handlerLogin)
	commands.register("register", handleRegister)
	commands.register("reset", handleReset)

	// parse cli arguments
	if len(os.Args) < 2 {
		log.Fatalf("please provide at least one argument")
	}

	cmd := NewCommand()
	cmd.name = os.Args[1]
	cmd.args = os.Args[2:]

	// run the command as the LAST step
	err = commands.run(&currentState, cmd)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
