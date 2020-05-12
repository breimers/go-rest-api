# Sample Go REST API

Supports Create, Read, Update, Delete

## Build

```bash
  go build
```

## Run

```bash
  ./go-rest-api
```

## Test

```bash
  $ curl -d '{...}' localhost:8000/api/books
    {...}
  $ curl localhost:8000/api/books
    {...}
  $ curl -X PUT -d '{...}' localhost:8000/api/books/{id}
    {...}
  $ curl -X DELETE localhost:8000/api/books/{i}
```
