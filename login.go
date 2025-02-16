package main

import (
    "fmt"
    "context"
)

func handlerLogin(s *state, cmd command) error{
    if len(cmd.args) < 3{
        return fmt.Errorf("The login handler received no arguments, expected arguments: username")
    }

    user, err := s.db.GetUser(context.Background(), cmd.args[2])
    if err != nil{
        return fmt.Errorf("User does not exist: %v", err)
    }
    
    err = s.config.SetUser(user.Name)
    if err != nil{
        return fmt.Errorf("Error setting a user: %v", err)
    }

    fmt.Printf("User %v has been set!\n", user.Name)
    return nil
}
