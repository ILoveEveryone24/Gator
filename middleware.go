package main

import (
	"fmt"
	"context"

	"github.com/iloveeveryone/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User)error) func(*state, command)error{
	return func(s *state, cmd command)error{
		user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
		if err != nil{
			return fmt.Errorf("middle failed: %v", err)
		}
		return handler(s, cmd, user)
	}	
}
