package main

import (
    "context"
    _ "github.com/lib/pq"
    "github.com/google/uuid"
    "fmt"
    "time"

    "github.com/ILoveEveryone24/Gator/internal/database"
)

func handlerRegister(s *state, cmd command) error{
    if len(cmd.args) < 3{
        return fmt.Errorf("The register handler received no arguments, expected arguments: name")
    }
    user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
        ID: uuid.New(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Name: cmd.args[2],
    })
    if err != nil{
        return fmt.Errorf("Failed to create user: %v", err)
    }
    
    s.config.SetUser(user.Name)

    fmt.Println("User was successfully created:", user.Name)
    return nil
}
