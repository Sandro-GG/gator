# gator

A CLI blog aggregator written in Go.

## What it is

`gator` is a command-line tool for following RSS feeds. It stores users, feeds,
and posts in a PostgreSQL database so you can browse what you've subscribed to
without leaving the terminal.

## Stack

- **Go** for the CLI
- **PostgreSQL** for storage
- **[Goose](https://github.com/pressly/goose)** for schema migrations
- **[SQLC](https://sqlc.dev/)** to generate type-safe Go from raw SQL

## Features

- JSON config file at `~/.gatorconfig.json` that tracks the database URL and the current user
- User registration backed by Postgres
- Login that switches the active user (and rejects unknown ones)
- Database reset command to quickly wipe records during development
- `users` command to list all registered users, marking the currently logged-in one
- RSS feed fetching and parsing into structured Go types, with HTML entities decoded
- `addfeed` command to add a new RSS feed to the database, which automatically follows it for the active user
- `feeds` command to list all feeds in the database, including the name of the user who added each one
- **Follow and unfollow system** to track or remove subscriptions to existing feeds
- **Background fetching loop (`agg`)** that continuously sweeps and stores new posts from the oldest updated feeds
- **Configurable browsing (`browse`)** to catch up on the most recent posts from your followed feeds right inside your terminal

## Prerequisites

Before running `gator`, ensure you have the following installed on your machine:

- **Go** (version 1.22 or higher recommended)
- **PostgreSQL** (running locally or remotely)

## Installation

You can install the `gator` CLI globally to your machine's `$GOPATH/bin` directory by running the following command inside the project root:

```bash
go install .
```

Ensure your terminal's shell path includes your Go bin directory (typically `~/go/bin`) so you can run the `gator` binary from anywhere.

## Setup & Getting Started

### 1. Initialize the Configuration

Create a configuration file in your home directory named `.gatorconfig.json`. Populate it with your database connection string and a placeholder for your current user:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

### 2. Run Database Migrations

Make sure your target PostgreSQL database exists, then use Goose to run the migration schemas up to date:

```bash
goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```

### 3. Try Out the Commands

Start using the tool by running the compiled binary directly. Here are a few core commands to get you started:

- **Register a new account:**
  ```bash
  gator register <username>
  ```
- **Log in as an existing user:**
  ```bash
  gator login <username>
  ```
- **Add a new feed source:**
  ```bash
  gator addfeed "Hacker News RSS" "https://hnrss.org"
  ```
- **Follow an existing feed:**
  ```bash
  gator follow "https://hnrss.org"
  ```
- **Run the background crawler worker (set duration pace):**
  ```bash
  gator agg 1m
  ```
- **Browse your personalized feed items (defaults to a limit of 2):**
  ```bash
  gator browse 5
  ```
