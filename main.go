package main

import (
    "fmt"
    "os"

    "github.com/iloveeveryone/gator/internal/config"
)

func main(){
    args := os.Args
    if len(args) < 2{
        fmt.Println("Too few arguments")
        os.Exit(1)
    }

    c, err := config.Read()
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }

    s := state{
        config: &c,
    }

    cmds := commands{
        commands: make(map[string]func(*state, command)error),
    }
    cmds.register("login", handlerLogin)

    cmd := command{
        name: args[1],
        args: args,
    }

    err = cmds.run(&s, cmd)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
}
