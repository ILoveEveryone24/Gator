package main

import (
    _ "github.com/lib/pq"
    "fmt"
    "os"
    "database/sql"

    "github.com/ILoveEveryone24/Gator/internal/config"
    "github.com/ILoveEveryone24/Gator/internal/database"
)

func main(){
    args := os.Args
    if len(args) < 2{
        fmt.Println("Too few arguments")
        os.Exit(1)
    }

	if args[1] == "init"{
		if len(args) < 3{
			fmt.Println("The init handler received no arguments, expected arguments: database_url\nFormat: postgres://username:password@localhost:5432/dbname?sslmode=disable\nExample: postgres://postgres:postgres@localhost:5432/gator?sslmode=disable")
			os.Exit(1)
		}

		db_url := args[2]

		c := config.Config{
			Db_url: "",
			Current_user_name: "",
		}
		c.SetDbUrl(db_url)

		err := c.SetUser("Default_user")
		if err != nil{
			fmt.Printf("Failed to create a default user: %v\n", err)
			os.Exit(1)
		}
		
		/*err = runMigrations(db_url)
		if err != nil{
			fmt.Printf("Failed to run migrations: %v\n", err)
			os.Exit(1)
		}*/

		fmt.Println("Successfully initialized!")
		os.Exit(0)
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
    cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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
