
# crud-golang

This pet project is an implementation of CRUD using golang, and has some UI for usability

## Run Locally

Clone the project

```bash
  git clone https://github.com/tgkzz/crud-golang.git
```

Go to the project directory

```bash
  cd crud-golang
```

Install dependencies

```bash
  go mod download
```

Start the server

```bash
  go run .
```

server will be started on localhost:4000


## Used Technology

- [HTML](https://www.w3schools.com/html/) (Front-end)
- [CSS](https://www.w3schools.com/css/) (Stylize)
- [Go](https://go.dev/) (Server part)
- [Postgres](https://www.postgresql.org/) (Database)
- SOON: [Redis](https://redis.io/) (Database seems to work faster with it)

## TODO

- Implement handlers for read (done) and update (done)
- ErrorHandling
- Add [redis](https://redis.io/)
- Transfer project to clean architecture

## Important menthion

as soon as i am not good programmer, in file [dbconn.go](https://github.com/tgkzz/crud-golang/blob/main/internal/database/dbconn.go), you need to change the path to your database, so it will work for your postgres database. Remember it, while executing this project

## Feedback

If you have any feedback, please reach out to us at [telegram](https://t.me/tgkmdk). I will answer as fast as I can)