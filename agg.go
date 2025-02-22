package main

import (
    "fmt"
    "context"
)

func handlerAgg(s *state, cmd command)error{
    f, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
    if err != nil{
        return fmt.Errorf("Failed to fetch feed: %v", err)
    }

    fmt.Printf("%+v\n", f)
    return nil
}
