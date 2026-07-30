package main

import (
	"context"
	"fmt"
)

func handlerAggregate(s *state, cmd command) error {
	// Replace constant value with cmd.Args[0] when done testing
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(feed.Channel.Title)
	fmt.Println(feed.Channel.Description)
	fmt.Println(feed.Channel.Link)
	for _, entry := range feed.Channel.Item {
		fmt.Println(entry.Title)
		fmt.Println(entry.Description)
		fmt.Println(entry.Link)
		fmt.Println(entry.PubDate)
	}

	return nil
}
