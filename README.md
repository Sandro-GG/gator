# gator

A CLI blog aggregator written in Go.

## Prerequisites

- [Go](https://go.dev/doc/install) 1.22 or later

## Installation

```sh
go install github.com/Sandro-GG/gator@latest
```

Or clone and build from source:

```sh
git clone https://github.com/Sandro-GG/gator.git
cd gator
go build
```

## Configuration

gator reads its configuration from `~/.gatorconfig.json`. Create the file
before running any commands:

```json
{
  "db_url": "postgres://example"
}
```

| Field               | Description                                  |
| ------------------- | -------------------------------------------- |
| `db_url`            | Connection string for the Postgres database. |
| `current_user_name` | Set automatically by the `login` command.    |

## Usage

```sh
gator <command> [args...]
```

During development you can run commands without installing:

```sh
go run . <command> [args...]
```

### Commands

#### `login <username>`

Sets the current user in the config file.

```sh
gator login alice
```

```
The username alice has been set
```

## Exit codes

| Code | Meaning                                              |
| ---- | ---------------------------------------------------- |
| `0`  | Command ran successfully.                            |
| `1`  | No command given, missing arguments, or handler error. |

## Project structure

```
.
├── main.go              # entrypoint: reads config, registers commands, dispatches
├── types.go             # state, command, and commands registry
├── handlers.go          # command handler functions
└── internal/
    └── config/
        └── config.go    # reading and writing ~/.gatorconfig.json
```

## Adding a command

1. Write a handler matching the signature `func(*state, command) error` in `handlers.go`.
2. Register it in `main.go`:

```go
commands.register("yourcommand", handlerYourCommand)
```