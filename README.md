```
  ___    _         
 | __|__| |_  ___  
 | _|/ _| ' \/ _ \ 
 |___\__|_||_\___/                 
```

**Echo** -- effortless mock server for rapid API emulation with JSON

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

| # | Part   | Required | Default |
|---|--------|:--------:|:-------:|
| 1 | Method |          |   GET   |
| 2 | Code   |          |   200   |
| 3 | Path   |    *     |         |

#### Mapping example

| Key              | Method | Path      | Code |
|------------------|:------:|-----------|:----:|
| /orders          |  GET   | /orders   | 200  |
| 404:/orders/1    |  GET   | /orders/1 | 404  |
| GET:/orders/2    |  GET   | /orders/2 | 200  |
| POST:201:/orders |  POST  | /orders   | 201  |

### Example

```json
{
  "/v1/orders/1": {
    "id": 1,
    "status": "confirmed"
  },
  "POST:201:/v1/orders": {
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
curl -i localhost:8080/v1/orders/1
```

```bash
curl -i -X POST localhost:8080/v1/orders
```

## Reverse proxy mode

Redirect all non-defined routes to the specified host:

```json
{
  "*": "http://localhost:8080"
}
```

Override a specific route:

```json
{
  "*": "http://localhost:8080",
  "/v1/orders/1": {
    "id": 1,
    "status": "confirmed"
  }
}
```