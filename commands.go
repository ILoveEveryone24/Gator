package main

import (
    "fmt"
    
    "github.com/iloveeveryone/gator/internal/config"
    "github.com/iloveeveryone/gator/internal/database"
)

type state struct{
    db *database.Queries
    config *config.Config
}

type command struct{
    name string
    args []string
}

type commands struct{
    commands map[string]func(*state, command)error
}

func (c *commands) register(name string, f func(*state, command)error){
    c.commands[name] = f
}

func (c *commands) run(s *state, cmd command)error{
    v, ok := c.commands[cmd.name]
    if !ok{
        return fmt.Errorf("Command doesn't exist", cmd.name)
    }
    return v(s, cmd)
}
