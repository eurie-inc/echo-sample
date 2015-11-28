# echo-sample

Sample Application Structure for micro Web Framework [Echo](https://github.com/labstack/echo).

## HTTP Request Sample

### POST new record

```bash
curl -H "Content-Type: application/json" -X POST -d '{"name":"Luis", "number": 9}' http://localhost:8080/api/v1/members
curl -H "Content-Type: application/json" -X POST -d '{"name":"Lionel", "number": 10}' http://localhost:8080/api/v1/members

# => {"number":10,"name":"Lionel","createdAt":1448723606599520746}

```

### GET a record

```bash
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/v1/members/11

# => {"number":10,"name":"Lionel","createdAt":1448723606599520746}
```

### GET records

```bash

curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/v1/members

# => [{"number":9,"name":"Luis","createdAt":1448723839698017936},{"number":10,"name":"Lionel","createdAt":1448723236912795004}]
```
