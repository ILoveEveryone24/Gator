package main

import (
    "fmt"
    "context"
	"time"
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/ILoveEveryone24/Gator/internal/database"
)

func handlerAgg(s *state, cmd command)error{
	if len(cmd.args) < 3{
		return fmt.Errorf("The aggregation handler received no arguments, expected arguments: time_between_requests")
	}
	t, err := time.ParseDuration(cmd.args[2])
	if err != nil{
		return fmt.Errorf("Failed to parse duration: %v", err)
	}
	tenSec := time.Second * 10
	
	if t < tenSec{
		t = tenSec
		fmt.Printf("Too short duration, set to default (%v)\n", tenSec)
	} else{
		fmt.Printf("Duration set to %v\n", t)
	}
	ticker := time.NewTicker(t)
	for ; ; <-ticker.C{
		scrapeFeeds(s)
	}
    return nil
}

func scrapeFeeds(s *state){
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil{
		fmt.Printf("Failed to fetch url for feed: %v\n", err)
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil{
		fmt.Printf("Failed to mark feed as fetched: %v\n", err)
		return
	}
    f, err := fetchFeed(context.Background(), feed.Url)
    if err != nil{
        fmt.Printf("Failed to fetch feed: %v\n", err)
		return
    }

	fmt.Printf("\nFetching from... %v\n\n", f.Channel.Title)
	for _, item := range f.Channel.Item{
		//Creating posts
		pubDate, err := parseDate(item.PubDate)
		if err != nil{
			fmt.Println("Failed to parse publication time")
			continue
		}
		isValid := true
		if item.Description == ""{
			isValid = false
		}
		description := sql.NullString{
			String: item.Description,
			Valid: isValid,
		}

		_, err = s.db.CreatePost(
			context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: description,
			PublishedAt: pubDate,
			FeedID: feed.ID,
		})
		if err != nil{
			if pqErr, ok := err.(*pq.Error); ok{
				if pqErr.Code != "23505"{
					fmt.Printf("\n\nERROR: %v\n\n\n", err)
				}
			}
		}
	}
	fmt.Printf("\nFinished fetching\n")
}

func parseDate(dateStr string) (time.Time, error){
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, layout := range layouts{
		t, err := time.Parse(layout, dateStr)
		if err == nil{
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("Couldn't find correct layout")
}
