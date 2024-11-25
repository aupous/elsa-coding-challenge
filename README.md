# Elsa Coding Challenge

# Prerequisites

- Having `redis` installed with version >= 7.2.0, or having `docker` installed
- If you have `docker` installed, there's a `docker-compose.yml` file already, run this command at repository level to start redis container:

```
docker-compose up -d
```

# Run the project

- Run the `main.go` file in `cmd/script` to prepare and add data to redis
- Run the `main.go` file in `cmd/http` to start http server
- There're two API that you can test with the http server

### Get the leaderboard

```
curl --location 'http://localhost:9000/v1/leaderboard/test-key'
```

### Get rank and score of an user

```
curl --location 'http://localhost:9000/v1/leaderboard/test-key/user-rank?userId=user-100'
```
