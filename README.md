```
  ___    _         
 | __|__| |_  ___  
 | _|/ _| ' \/ _ \ 
 |___\__|_||_\___/                 
```

---

![Build Status](https://github.com/devem-tech/echo/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/devem-tech/echo)](https://goreportcard.com/report/github.com/devem-tech/echo)

**Echo** -- easy-to-use mock server that accelerates api emulation with json.

### Features

- simple json format
- response latency
- reverse proxy mode
- request and response logs

## How to use

### Create a mock file

The json format is as follows:

- **key** -- route
- **value** -- response

The **key** is a colon-separated string which contains 3 parts:

| # | Part        | Required | Default |
|---|-------------|:--------:|:-------:|
| 1 | Method      |    *     |         |
| 2 | Endpoint    |    *     |         |
| 3 | Status code |          |   200   |

#### Mapping example

| Key               | Method | Endpoint  | Status code |
|-------------------|:------:|-----------|:-----------:|
| GET:/orders       |  GET   | /orders   |     200     |
| POST:/orders:201  |  POST  | /orders   |     201     |
| GET:/orders/1:404 |  GET   | /orders/1 |     404     |
| GET:/orders/2:500 |  GET   | /orders/2 |     500     |

### Example

```json
{
  "GET:/v1/orders/1": {
    "id": 1,
    "status": "confirmed"
  },
  "POST:/v1/orders:201": {
    "success": true
  }
}
```

### Go

```bash
echo routes.json
```

```bash
Options:
  --port PORT, -p PORT   port [default: 8080]
  --color, -c            color output [default: true]
  --latency LATENCY, -l LATENCY
                         response latency [default: 0]
  --print PRINT          string specifying what the output should contain:
    'H' request headers
    'B' request body
    'h' response headers
    'b' response body
  --verbose, -v          verbose output
  --help, -h             display this help and exit
  --version              display version and exit
```

### Testing

```bash
curl 127.0.0.1:8080/v1/orders/1
```

```bash
curl -X POST 127.0.0.1:8080/v1/orders
```

## Reverse proxy mode

Redirect all non-defined routes to the specified host:

```json
{
  "*": "https://jsonplaceholder.typicode.com"
}
```

Override a specific route:

```json
{
  "*": "https://jsonplaceholder.typicode.com",
  "GET:/todos/1": {
    "id": 1,
    "userId": 1,
    "title": "default",
    "completed": false
  }
}
```