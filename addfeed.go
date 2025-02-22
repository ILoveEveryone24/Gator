package main

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"

	"github.com/iloveeveryone/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command)error{
	if len(cmd.args) < 4{
		return fmt.Errorf("The add feed handler received too few arguments, expected arguments: feed name, url")
	}

	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil{
		return fmt.Errorf("Failed to fetch the user: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[2],
		Url: cmd.args[3],
		UserID: user.ID,
	})
	if err != nil{
		return fmt.Errorf("Failed to create the feed: %v", err)
	}
	
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil{
		return fmt.Errorf("Failed to create a feed follow: %v", err)
	}

	fmt.Printf("Added feed: %+v\n", feed)

	return nil
}
