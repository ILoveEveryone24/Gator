package main

import (
    _ "github.com/lib/pq"
    "fmt"
    "os"
    "database/sql"

    "github.com/iloveeveryone/gator/internal/config"
    "github.com/iloveeveryone/gator/internal/database"
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

    db, err := sql.Open("postgres", c.Db_url)
    if err != nil{
        fmt.Println("Failed to load database")
        os.Exit(1)
    }
    defer db.Close()
    
    dbQueries := database.New(db)

    s := state{
        db: dbQueries,
        config: &c,
    }

    cmds := commands{
        commands: make(map[string]func(*state, command)error),
    }
    cmds.register("login", handlerLogin)
    cmds.register("register", handlerRegister)
    cmds.register("reset", handlerReset)
    cmds.register("users", handlerUsers)
    cmds.register("agg", handlerAgg)
    cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
    cmds.register("feeds", handlerFeeds)
    cmds.register("follow", middlewareLoggedIn(handlerFollow))
    cmds.register("following", middlewareLoggedIn(handlerFollowing))
    cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

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
