package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch the feed: %w", err)
	}

	fmt.Printf("Channel Title: %s\nChannel Description: %s\n", feed.Channel.Title, feed.Channel.Description)

	fmt.Printf("%+v\n", feed)

	return nil
}
