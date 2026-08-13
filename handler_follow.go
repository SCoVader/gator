package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/SCoVader/gator/internal/database"
)

// Logged in function
func handlerFollow(s *state, cmd command, user database.User) error {
	// Check argument amount
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feedUrl, err := url.Parse(cmd.Args[0])
	if err != nil {
		return err
	}

	// Get feed ID
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl.String())
	if err != nil {
		return err
	}

	followRecord, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(followRecord.FeedName)
	fmt.Println(followRecord.UserName)
	return nil
}
