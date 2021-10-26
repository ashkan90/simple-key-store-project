<h1 align="center">Yemeksepeti Case</h1>

![example workflow](https://github.com/ashkan90/ys-project/actions/workflows/github-actions.yml/badge.svg)


> This project aims to implement very simple in-memory. Written for Delivery Hero case with Hexagonal Architecture.

### ✨ [Project Runs On](https://ys-project.herokuapp.com/)

# Hexagonal Architecture
The idea of Hexagonal Architecture is to put inputs and outputs at the edges of our design. Business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from — a database, a microservice API exposed via gRPC or REST, or just a simple CSV file. <br>

The pattern allows us to isolate the core logic of our application from outside concerns. Having our core logic isolated means we can easily change data source details without a significant impact or major code rewrites to the codebase.

## Run on Local Machine

```shell
go get .
PORT=8080 go run ./cmd/
```

## Endpoints

```console
GET /key?key=myKey
POST /key ## needs json payload as body
```

## Example Usages

```shell
curl --location --request POST 'http://localhost:8080/key' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "k",
    "value": "v"
}'

curl --location --request GET 'http://localhost:8080/key?key=myKey'
```

## Run Tests

```console
go test ./...
go test -race ./domain/*
```

## Deployment

> To deploy the case, I used `Heroku`. The deployment is automated with `github actions` and it's `containerized`

## Author

👤 **Emirhan Ataman**


## 📝 License

Copyright © 2021 [Emirhan Ataman](https://github.com/ashkan90). <br />