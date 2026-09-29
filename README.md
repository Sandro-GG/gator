# gator

A CLI blog aggregator written in Go. Work in progress.

## What it is

`gator` is a command-line tool for following RSS feeds. It stores users, feeds,
and posts in a PostgreSQL database so you can browse what you've subscribed to
without leaving the terminal.

## Stack

- **Go** for the CLI
- **PostgreSQL** for storage
- **[Goose](https://github.com/pressly/goose)** for schema migrations
- **[SQLC](https://sqlc.dev/)** to generate type-safe Go from raw SQL

## Current features

- JSON config file at `~/.gatorconfig.json` that tracks the database URL and the current user
- User registration backed by Postgres
- Login that switches the active user (and rejects unknown ones)
- Database reset command to quickly wipe records during development
- `users` command to list all registered users, marking the currently logged-in one
- RSS feed fetching and parsing into structured Go types, with HTML entities decoded
- `agg` command that fetches a feed and prints its channel and items
- `addfeed` command to add a new RSS feed to the database, linked to the current user
- `feeds` command to list all feeds in the database, including the name of the user who added each one

## Roadmap

- Follow and unfollow feeds
- Background fetching and storing of posts
- Browse recent posts for the logged-in user
