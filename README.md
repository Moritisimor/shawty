# shawty
URL shortener written in Go

## What is this project about?
This is a very simple and straightforward URL shortener written in Go.
 
It uses `net/http` for its HTTP API and a SQLite3 database for storing key-value pairs.

## Cloning and Compilation
For this you will require a go compiler, preferably the latest, as well as git.

A small script for cloning and compiling the application:
```bash
git clone https://github.com/Moritisimor/shawty
cd shawty
go build -o shawty -ldflags="-s -w" cmd/shawty/Main.go
```

Now to run it:
```bash
./shawty
```

Try visiting `http://localhost:8080` in your browser!

## Environment Variables
You can set a couple of environment variables before running the application for a bit of configuration.

### DB_PATH
This is the path for the SQLite3 database that is to be used.

Usually, this is set to `./shawty.db`.

### PORT
The port on which the server listens.

Usually, this is set to `8080`.

### ADDRESS
This is the bind-address for the server.

Usually, this is set to `0.0.0.0`.

### DELETE_AFTER_HOURS
This is how long a URL-Alias may exist after its creation.

Usually, this is set to `24`.

### SLEEP_TIME_MINS
This is how often the cleaner goroutine runs, deleting expired Aliases.

Usually, this is set to `20`.

## REST Endpoints
These are the REST endpoints of the application.

### GET /api/status
Returns the status of the web backend.

If everything is alright, status code `200 OK` will be returned.

### GET /link/{link}
Redirects the client to the given URL that is behind `link`. 

If it is not found, `404 Not Found` is returned.

If the alias does exist, `302 Found` is returned, and the client is redirected to the actual URL.

### POST /api/alias
Posts a new alias for a URL.

If the alias already exists, `409 Conflict` is returned.

If it doesn't already exist, the alias is stored in the database in association with the URL. `200 Ok` is returned.

If, for any reason, saving the alias to the database fails, `500 Internal Server Error` is returned. In this case, check the logs which are written to stdout by the application.

### GET /
Gets the frontend, which is embedded in the binary.
