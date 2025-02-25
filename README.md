# Gator – A Local RSS Aggregator Written in Go

Gator is a lightweight RSS aggregator that runs on your local machine, allowing you to collect, follow, and browse RSS feeds efficiently.

## Prerequisites
To use Gator, ensure you have the following installed on your system:
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

**IMPORTANT:**
Make sure to create a database called "gator" after configuring your PostgreSQL with the command:
  ```sh
  CREATE DATABASE gator;
  ```
(This will be automated in the near future)

You will also need to "git clone" the repository and run "Gator init ..." command from there the first time to run all the migrations. Alternatively, you can run the migrations from the "sql/schema" directory manually or with "goose".
(This will be automated in the near future)

## Installation

To install Gator, you can use the following Go command:

  ```sh
  go install github.com/ILoveEveryone24/Gator@latest
  ```

Before using the application, you must initialize the database with the following command:  

  ```sh
  gator init <db_url>  
  ```

For example:  

  ```sh
  gator init postgres://postgres:postgres@localhost:5432/gator?sslmode=disable
  ```
This will create ".gatorconfig.json" file at your home directory, which will contain the current user and the database url you provided.

## Usage

### User Management
- **Register a new user:**
  ```sh
  gator register <username>
  ```
- **Log in as a registered user:**
  ```sh
  gator login <username>
  ```
- **List all registered users:**
  ```sh
  gator users
  ```
- **Reset everything:**
  ```sh
  gator reset
  ```

### Feed Management
- **Add a new feed to the database:**
  ```sh
  gator addfeed "<title>" "<url>"
  ```
- **List all added feeds:**
  ```sh
  gator feeds
  ```
- **Follow a feed (must be added first):**
  ```sh
  gator follow "<url>"
  ```
- **View all followed feeds:**
  ```sh
  gator following
  ```
- **Unfollow a feed:**
  ```sh
  gator unfollow "<url>"
  ```

### Fetching & Browsing Posts
- **Aggregate posts from followed feeds:**
  ```sh
  gator agg <duration between requests>
  ```
  ⚠️ This starts an **infinite loop** that continuously fetches posts at the specified interval. To stop, use `Ctrl + C`. Once you've collected enough posts, use the `browse` command to view them.

- **Browse aggregated posts:**
  ```sh
  gator browse <optional: number of posts>
  ```
  If no number is specified, the default is **2** posts.

---

## Notes
- A feed **must** be added before it can be followed.
- Aggregation (`agg`) will continue running indefinitely until manually stopped (`Ctrl + C`).

