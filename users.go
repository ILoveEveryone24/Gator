package main

import (
    "fmt"
    "context"
)

func handlerUsers(s *state, cmd command)error{
    users, err := s.db.GetUsers(context.Background())
    if err!= nil{
        return fmt.Errorf("Failed to retrieve users: %v", err)
    }
    for _, name := range users{
        if name == s.config.Current_user_name{
            fmt.Printf("* %s (current)\n", name)
            continue
        }
            fmt.Printf("* %s\n", name)
    }
    return nil
}
