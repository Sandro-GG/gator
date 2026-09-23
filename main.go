package main

import (
	"fmt"
	"log"

	"github.com/Sandro-GG/gator/internal/config"
)

func main() {
	content, err := config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = content.SetUser("Sandro")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	content, err = config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("db_url: %s\ncurrent_user_name: %s", content.DbUrl, content.CurrentUsername)

}
