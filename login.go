package main

import "fmt"

func handlerLogin(s *state, cmd command) error{
    if len(cmd.args) < 3{
        return fmt.Errorf("The login handler received no arguments, expected arguments: username")
    }
    
    err := s.config.SetUser(cmd.args[2])
    if err != nil{
        return fmt.Errorf("Error setting a user: %v", err)
    }

    fmt.Printf("User %v has been set!\n", cmd.args[2])
    return nil
}
