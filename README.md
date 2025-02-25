# Gator – A Local RSS Aggregator Written in Go

Gator is a lightweight RSS aggregator that runs on your local machine, allowing you to collect, follow, and browse RSS feeds efficiently.

## Prerequisites
To use Gator, ensure you have the following installed on your system:
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

Before using the application, you must initialize the database with the following command:  

  ```sh
  <program> init <db_url>  
  ```

For example:  

  ```sh
  <program> init postgres://postgres:postgres@localhost:5432/gator?sslmode=disable
  ```
This will create ".gatorconfig.json" file at your home directory, which will contain the current user and the database url you provided.

## Usage

### User Management
- **Register a new user:**
  ```sh
  <program> register <username>
  ```
- **Log in as a registered user:**
  ```sh
  <program> login <username>
  ```
- **List all registered users:**
  ```sh
  <program> users
  ```
- **Reset everything:**
  ```sh
  <program> reset
  ```

### Feed Management
- **Add a new feed to the database:**
  ```sh
  <program> addfeed "<title>" "<url>"
  ```
- **List all added feeds:**
  ```sh
  <program> feeds
  ```
- **Follow a feed (must be added first):**
  ```sh
  <program> follow "<url>"
  ```
- **View all followed feeds:**
  ```sh
  <program> following
  ```
- **Unfollow a feed:**
  ```sh
  <program> unfollow "<url>"
  ```

### Fetching & Browsing Posts
- **Aggregate posts from followed feeds:**
  ```sh
  <program> agg <duration between requests>
  ```
  ⚠️ This starts an **infinite loop** that continuously fetches posts at the specified interval. To stop, use `Ctrl + C`. Once you've collected enough posts, use the `browse` command to view them.

- **Browse aggregated posts:**
  ```sh
  <program> browse <optional: number of posts>
  ```
  If no number is specified, the default is **2** posts.

---

## Notes
- A feed **must** be added before it can be followed.
- Aggregation (`agg`) will continue running indefinitely until manually stopped (`Ctrl + C`).

