# ProgWebTPE — Quick Start Guide
# Prerequisites
    - Linux, Bash, and Make
    - Docker
    - Go 1.22+ (or compatible)
    - sqlc (only if you intend to modify the queries)
    - templ (only if you intend to modify DOM elements)

# Variables (Default values in the Makefile)
- Container: `some-postgres`
- DB User: `postgres`
- Password: `XYZ`
- DB: `tpespecial`
- Image: `postgres:16`
- HTTP Port: `8080`

These can be overridden when invoking make, for example:
`make DB_PASSWORD=another_pass DB_NAME=another_db setup`

# STEPS:

# 1- Prepare the database (2 options) via Linux console:
Create or start the container, wait for availability, create the DB if missing, and apply the schema.
```bash
make setup
```

Alternatively, more simply (this also starts the server and executes requests.bash):
```bash
    make test
```

## 2- Start the server:
In a terminal:
```bash
go run ./cmd/server
```
The API will be available in `http://localhost:8080`.

## 3- Test the API (2 options):
    - Run the request script:
```bash
        chmod +x ./requests.bash
        ./requests.bash
```
    - Or use the task that starts the server in the background, runs the script, and shuts it down:
```bash
    make test
```

## 4 - View resolution:
    Open in your browser:

    http://localhost:8080/

## 5-Port and container cleanup
    Stop and remove the container:
```bash
    make clean
```
