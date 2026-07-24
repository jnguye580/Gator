# Gator

Gator is a local multi-user CLI blog aggregator built in Go that stores configuration in JSON and uses PostgreSQL for shared data.

## Requirements

Before running Gator, make sure you have the following installed:

- **Go** (1.26 or later) — [install instructions](https://go.dev/doc/install)
- **PostgreSQL** — [install instructions](https://www.postgresql.org/download/)

Gator connects to a Postgres database to store users, feeds, and posts, so you'll need a running Postgres server and a database created for Gator to use.

## Installing the CLI

With Go installed, you can install the `gator` CLI directly from the repo:

```bash
go install github.com/jnguye580/GATOR-PROJECT@latest
```

This builds the binary and places it in your `$GOPATH/bin` (or `$HOME/go/bin`), so make sure that directory is on your `PATH`.

## Setting up the config file

Gator reads its configuration from a `.gatorconfig.json` file in your home directory. Create it with your Postgres connection string:

```bash
echo '{"db_url":"postgres://username:password@localhost:5432/gator?sslmode=disable","current_user_name":""}' > ~/.gatorconfig.json
```

Update the `db_url` to match your local Postgres setup (username, password, host, port, and database name). The `current_user_name` field is managed automatically by the `login`/`register` commands.

Once the config file is in place, run the database migrations in `sql/schema` against your database (e.g. using [goose](https://github.com/pressly/goose)) to create the required tables.

## Running the program

Once installed and configured, run commands like:

```bash
gator register alice
gator login alice
```

Some commands you can run:

- `gator register <name>` — create a new user and log in as them
- `gator login <name>` — switch the current user
- `gator addfeed <name> <url>` — add a new RSS feed and follow it
- `gator follow <url>` — follow an existing feed
- `gator following` — list the feeds you're following
- `gator agg <duration>` — continuously fetch new posts from feeds on an interval (e.g. `gator agg 1m`)
- `gator browse [limit]` — print recent posts from feeds you follow (defaults to 2 posts)
