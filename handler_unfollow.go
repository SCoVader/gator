package main

import (
	"context"
	"fmt"

	"github.com/SCoVader/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	// Check argument amount
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	url := cmd.Args[0]
	// Get feed ID
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	return s.db.Unfollow(context.Background(), database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
}
