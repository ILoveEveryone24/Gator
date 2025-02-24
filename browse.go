package main

import (
	"fmt"
	"strconv"
	"context"

	"github.com/iloveeveryone/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User)error{
	limit := 2
	var err error
	if len(cmd.args) > 2{
		limit, err = strconv.Atoi(cmd.args[2])
		if err != nil{
			return fmt.Errorf("Incorrect input, expected integer: %v", err)
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil{
		return fmt.Errorf("Failed to fetch posts: %v", err)
	}

	for _, post := range posts{
		fmt.Printf("Title: %v\n", post.Title)
		fmt.Printf("Description: %v\n\n", post.Description.String)
		fmt.Printf("Link: %v\n\n\n", post.Url)
	}

	return nil
}
