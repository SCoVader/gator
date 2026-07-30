package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	rows, err := s.db.GetFeedsShort(context.Background())

	if err != nil {
		return err
	}
	for _, row := range rows {
		fmt.Println(row.Name, " ", row.Url, " ", row.AddedBy)
	}
	return nil
}
