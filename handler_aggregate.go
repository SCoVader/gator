package main

import (
	"fmt"
	"time"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <duration>", cmd.Name)
	}

	time_between_requests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println("Collecting feeds every ", time_between_requests.String())
	ticker := time.NewTicker(time_between_requests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
	//return nil
}
