package main

import (
	"context"

	"github.com/SCoVader/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		userData, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, userData)
		if err != nil {
			return err
		}
		return nil
	}
}
