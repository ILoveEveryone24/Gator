package main

import (
	"fmt"
	"context"
)

func handlerFollowing(s *state, cmd command)error{
	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil{
		return fmt.Errorf("Failed to fetch the user: %v", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil{
		return fmt.Errorf("Failed to fetch following feeds: %v", err)
	}
	for _, feed := range feedFollows{
		fmt.Printf("Feed name: %s\n", feed.FeedName)
	}
	return nil
}
