package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func createUsersTable(db *sql.DB)error{
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id UUID PRIMARY KEY,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			name TEXT NOT NULL UNIQUE
		);
	`
	_, err := db.Exec(query)
	if err != nil{
		return err
	}
	return nil
}

func createFeedsTable(db *sql.DB)error{
	query := `
		CREATE TABLE IF NOT EXISTS feeds(
			id UUID PRIMARY KEY,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			name TEXT NOT NULL,
			url TEXT NOT NULL UNIQUE,
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			last_fetched_at TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	if err != nil{
		return err
	}
	return nil
}

func createFeedFollowsTable(db *sql.DB)error{
	query := `
		CREATE TABLE feed_follows(
			id UUID PRIMARY KEY,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
			UNIQUE (user_id, feed_id)
		);

	`
	_, err := db.Exec(query)
	if err != nil{
		return err
	}
	return nil
}

func createPostsTable(db *sql.DB)error{
	query := `
		CREATE TABLE posts(
			id UUID PRIMARY KEY,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			title TEXT NOT NULL,
			url TEXT UNIQUE NOT NULL,
			description TEXT,
			published_at TIMESTAMP NOT NULL,
			feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
		);
	`
	_, err := db.Exec(query)
	if err != nil{
		return err
	}
	return nil
}

func runMigrations(dbURL string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = createUsersTable(db)
	if err != nil{
		return fmt.Errorf("Failed to create Users table: %v", err)
	}
	err = createFeedsTable(db)
	if err != nil{
		return fmt.Errorf("Failed to create Feeds table: %v", err)
	}
	err = createFeedFollowsTable(db)
	if err != nil{
		return fmt.Errorf("Failed to create FeedFollows table: %v", err)
	}
	err = createPostsTable(db)
	if err != nil{
		return fmt.Errorf("Failed to create Posts table: %v", err)
	}

	fmt.Println("Database migrations applied successfully.")
	return nil
}

