package main

import (
	"fmt"
	"context"

	"github.com/iloveeveryone/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User)error{
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil{
		return fmt.Errorf("Failed to fetch following feeds: %v", err)
	}
	for _, feed := range feedFollows{
		fmt.Printf("Feed name: %s\n", feed.FeedName)
	}
	return nil
}
