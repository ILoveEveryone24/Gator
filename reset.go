package main

import (
    "fmt"
    "context"
)

func handlerReset(s *state, cmd command)error{
    err := s.db.DeleteUsers(context.Background())
    if err != nil{
        return fmt.Errorf("Failed to delete all users: %v", err)
    }
    fmt.Println("Successfully deleted all users!")
    return nil
}
