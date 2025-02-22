package main

import (
	"fmt"
	"context"
	
	"github.com/iloveeveryone/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User)error{
	if len(cmd.args) < 3{
		return fmt.Errorf("The unfollow handler received no arguments, expected arguments: URL")
	}
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url: cmd.args[2],
	})
	if err != nil{
		return fmt.Errorf("Failed to unfollow url: %v", err)
	}
	return nil
}
