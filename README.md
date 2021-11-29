# Classes Scheduler

A simple class scheduler written in Go.

## Requirements

- [Golang 1.17](https://go.dev/)
- [Docker](https://www.docker.com/)

## Running

Firstly, create a .env file with all the variables needed as follows:

```yaml
DATABASE_HOST=classes-scheduler-db
DATABASE_PORT=5432
DATABASE_USER=[CHOSE_A_USER]
DATABASE_PASS=[CHOSE_A_PASSWORD]
DATABASE_NAME=[CHOSE_A_DATABASE_NAME]
```

Run docker compose to start the environment:

```bash
$ docker compose up -d
```

The environment will startup with three containers:

- classes-scheduler: the application
- classes-scheduler-db: a PostgreSQL database
- adminer: a tiny database admin interface

## Tests

With all the containers running, run:

```bash
go test ./e2e
```

You can manually test the application using cURL too:

```bash
$ curl -X POST http://localhost:8080/api/classes \
    -H 'Content-Type: application/json' \
    -d '{"name":"Class Test","start_date":"2021-11-10", "end_date":"2021-11-20", "capacity": 50}'
```

```bash
curl -X POST http://localhost:8080/api/bookings \
    -H 'Content-Type: application/json' \
    -d '{"name":"Mateus","date":"2021-11-10"}'
```
