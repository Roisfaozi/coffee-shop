# User API Spec

## Create User API

Endpoint : POST /user/


Request Body :

```json
{
  "username": "Testing-1713977486",
  "password": "rahasiainimah",
  "email": "testing@example.com",
  "role": "user"
}
```

Response Body Success :

```json
{
  "id": "65e5361b-1d04-46ee-ba6f-9488b5693654",
  "username": "Testing-1713977486",
  "email": "testing@example.com",
  "role": "user"
}
```


## Update User API

Endpoint : PUT /user/:userId

Request Body :

```json
{
  "username": "Testing-1713977486",
  "password": "rahasiainimah",
  "email": "testing@example.com",
  "role": "user"
}
```

Response Body Success :

```json
{
  "id": "65e5361b-1d04-46ee-ba6f-9488b5693654",
  "username": "Testing-1713977486",
  "email": "testing@example.com",
  "role": "user"
}
```

## Get List User API

Endpoint : GET /user/


Response Body Success :

```json
[
    {
      "id": "65e5361b-1d04-46ee-ba6f-9488b5693654",
      "username": "Testing-1713977486",
      "email": "testing@example.com",
      "role": "user"
    },
    {
        "id": "c91c37c2-4749-45d0-b53e-d495faec9001",
        "username": "Testing-1713977487",
        "email": "testing@example.com",
        "role": "user"
    }
]
```


## Get User API

Endpoint : GET /user/:userId

Response Body Success :

```json 
{
  "id": "a3c1d8e0-5ce2-4e5e-bda6-2aaf5b067b34",
  "username": "Testing-1713980003",
  "email": "testing@example.com",
  "role": "user"
}
```


## Remove Address API

Endpoint : DELETE /user/:userId

Response Body Success :

```json
{
  "message": "User deleted successfully"
}
```
