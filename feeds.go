package main

import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command)error{
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil{
		return fmt.Errorf("Failed to fetch feeds: %v", err)
	}
	for i := range feeds{
		fmt.Printf("Name: %s\n", feeds[i].Name)
		fmt.Printf("URL: %s\n", feeds[i].Url)
		user, err := s.db.GetUserByID(context.Background(), feeds[i].UserID)
		if err != nil{
			fmt.Printf("Failed to fetch user: %v\n", err)
			continue
		}
		fmt.Printf("Username: %s\n", user.Name)
	}
	return nil
}
