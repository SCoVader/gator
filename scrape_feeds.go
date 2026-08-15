package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	content, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, post := range content.Channel.Item {
		fmt.Println(post.Title)
	}
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	return nil
}
