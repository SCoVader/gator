package main

import (
	"context"
	"fmt"

	"github.com/SCoVader/gator/internal/database"
)

// Logged in function
func handlerGetFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, record := range following {
		fmt.Println(record.FeedName)
	}

	return nil
}
