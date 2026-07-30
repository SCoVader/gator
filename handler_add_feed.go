package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/SCoVader/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name-of-feed> <feed-url>", cmd.Name)
	}

	url, err := url.Parse(cmd.Args[1])
	if err != nil {
		return err
	}

	curUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    url.String(),
		UserID: curUser.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed.ID.String())
	fmt.Println(feed.CreatedAt.GoString())
	fmt.Println(feed.UpdatedAt.GoString())
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)

	return nil
}
