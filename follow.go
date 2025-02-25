package main

import (
	"fmt"
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/ILoveEveryone24/Gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User)error{
	if len(cmd.args) < 3{
		return fmt.Errorf("The follow handler received no arguments, expected arguments: url")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[2])
	if err != nil{
		return fmt.Errorf("Failed to fetch the feed: %v", err)
	}

	feedFollows, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil{
		return fmt.Errorf("Failed to create a feed follow: %v", err)
	}

	fmt.Printf("%s successfully followed: %s\n", feedFollows.UserName, feedFollows.FeedName)
	return nil
}
