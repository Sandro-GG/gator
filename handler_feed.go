package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/Sandro-GG/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("please provide only name and url")
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("unable to get current user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currUser.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create a feed follow: %w", err)
	}

	fmt.Printf("Feed successfully created!\n")
	fmt.Printf("* Name:        %s\n", feed.Name)
	fmt.Printf("* URL:         %s\n", feed.Url)
	fmt.Printf("* Created By:  %s\n", currUser.Name)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("too many arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %s\nURL: %s\nCreated By: %s\n\n", feed.FeedName, feed.Url, feed.UserName)
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("please only provide the url parameter")
	}

	feedUrl := cmd.args[0]

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("unable to get current user: %w", err)
	}

	currFeed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("unable to get current feed: %w", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currUser.ID,
		FeedID:    currFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create a feed follow: %w", err)
	}

	fmt.Printf("User - %s: Feed - %s\n", follow.UserName, follow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("too many arguments")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("unable to get user's following feeds: %w", err)
	}

	fmt.Printf("User %s is following the feeds:\n", s.cfg.CurrentUsername)
	for _, follow := range follows {
		fmt.Printf(" - %s\n", follow.FeedName)
	}

	return nil
}
