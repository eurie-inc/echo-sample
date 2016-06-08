# echo-sample

Sample Application Structure for micro Web Application Framework [Echo](https://github.com/labstack/echo).

## HTTP Request Sample

### POST

```bash
curl -H "Content-Type: application/json" -X POST -d '{"name":"Luis", "number": 9, "position": "FW"}' http://localhost:8888/api/v1/members
curl -H "Content-Type: application/json" -X POST -d '{"name":"Lionel", "number": 10, "position": "FW"}' http://localhost:8888/api/v1/members

# => {"number":10, "name":"Lionel", "position":"FW", "createdAt":1465415304}

```

### GET a record

```bash
curl -H "Content-Type: application/json" -X GET http://localhost:8888/api/v1/members/10

# => {"number":10, "name":"Lionel", "position":"FW", "createdAt":1465415304}
```

### GET records

```bash

curl -H "Content-Type: application/json" -X GET http://localhost:8888/api/v1/members

# => [
       {"number":9, "name":"Luis", "position":"FW", "createdAt":1465415304},
       {"number":10,"name":"Lionel", "position":"FW", "createdAt":1465415304}
     ]
```
