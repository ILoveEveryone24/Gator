# Gator – A Local RSS Aggregator Written in Go

Gator is a lightweight RSS aggregator that runs on your local machine, allowing you to collect, follow, and browse RSS feeds efficiently.

## Prerequisites
To use Gator, ensure you have the following installed on your system:
- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

## Installation

**Warning**: The following steps will **update the PostgreSQL user password** to `postgres` and **create a new database** named `gator`. Please ensure you are aware of these changes before proceeding. The alternative would be to change the `install.sh` script to the correct credentials or alter the credentials back after running the script.

### To get started:

1. **Pull the repository**:
   ```bash
   git clone https://github.com/ILoveEveryone24/Gator.git
   ```

2. **Run the `install.sh` script**:
   Navigate to the repository's directory and execute the following:
   ```bash
   ./install.sh
   ```

The `install.sh` script will:
- Update the system package lists (`apt-get update`).
- Install PostgreSQL and PostgreSQL-contrib packages.
- Set the PostgreSQL `postgres` user's password to `postgres`.
- Create a new database named `gator`.

### Important Notes:
- The script will automatically **start PostgreSQL** and make the necessary changes to the database.
- Ensure that this password change does not conflict with any existing setup.

If you have any concerns or need to adjust the credentials, be sure to edit the `install.sh` script before running it.

Afterwards, you can install Gator on your machine and use it anywhere with the following Go command:

  ```sh
  go install github.com/ILoveEveryone24/Gator@latest
  ```

Before using the application, you must initialize the database with the following command:  

  ```sh
  Gator init <db_url>  
  ```

For example:  

  ```sh
  Gator init postgres://postgres:postgres@localhost:5432/gator?sslmode=disable
  ```
This will create ".gatorconfig.json" file at your home directory, which will contain the current user and the database url you provided.

## Usage

### User Management
- **Register a new user:**
  ```sh
  Gator register <username>
  ```
- **Log in as a registered user:**
  ```sh
  Gator login <username>
  ```
- **List all registered users:**
  ```sh
  Gator users
  ```
- **Reset everything:**
  ```sh
  Gator reset
  ```

### Feed Management
- **Add a new feed to the database:**
  ```sh
  Gator addfeed "<title>" "<url>"
  ```
- **List all added feeds:**
  ```sh
  Gator feeds
  ```
- **Follow a feed (must be added first):**
  ```sh
  Gator follow "<url>"
  ```
- **View all followed feeds:**
  ```sh
  Gator following
  ```
- **Unfollow a feed:**
  ```sh
  Gator unfollow "<url>"
  ```

### Fetching & Browsing Posts
- **Aggregate posts from followed feeds:**
  ```sh
  Gator agg <duration between requests>
  ```
  ⚠️ This starts an **infinite loop** that continuously fetches posts at the specified interval. To stop, use `Ctrl + C`. Once you've collected enough posts, use the `browse` command to view them.

- **Browse aggregated posts:**
  ```sh
  Gator browse <optional: number of posts>
  ```
  If no number is specified, the default is **2** posts.

---

## Notes
- A feed **must** be added before it can be followed.
- Aggregation (`agg`) will continue running indefinitely until manually stopped (`Ctrl + C`).

