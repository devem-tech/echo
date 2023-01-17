```
  __  __            _      _       
 |  \/  | ___   ___| | __ (_) ___  
 | |\/| |/ _ \ / __| |/ / | |/ _ \ 
 | |  | | (_) | (__|   < _| | (_) |
 |_|  |_|\___/ \___|_|\_(_)_|\___/ 
                                   
```

A very simple mock server for quick API emulation.

### Features

- simple json format
- simulated latency

## How to use

### Create a mock file

```json
{
  "/v1/orders": {
    "success": true
  },
  "/v1/orders/1": {
    "id": 1,
    "status": "confirmed"
  }
}
```

### Go

```bash
mockio -i data.json [-p 8080] [-l 1000]
```

### Testing

```bash
curl -i localhost:8080/v1/orders/1
```

```bash
curl -i -X POST localhost:8080/v1/orders
```