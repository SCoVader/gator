package main

import (
	"context"
	"fmt"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), name)
	if err != nil {
		return err
	}

	fmt.Println("Successfully created user ", user.Name)
	args := []string{user.Name}
	err = handlerLogin(s, command{
		Name: "login",
		Args: args,
	})
	if err != nil {
		return fmt.Errorf("register failed: %v", err)
	}

	return nil
}
