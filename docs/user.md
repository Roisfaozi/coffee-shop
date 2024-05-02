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
  "status": "Created",
  "data": {
    "id": "6b45af65-e7f9-42eb-b47b-ed952a894c4b",
    "username": "Testing-1714658608",
    "email": "testing@example.com"
  },
  "description": "1 data user created"
}
```
## Auth login user

Endpoint : POST /auth/login


Request Body :

```json
{
  "username": "Testing-1713977486",
  "password": "rahasiainimah"
}
```

Response Body Success :

```json
{
  "status": "OK",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijg2MjIyODkwLWY1ZWEtNGZkMy04MWM5LTIzMjMzNzQ4NjBlMCIsInJvbGUiOiJhZG1pbiIsImlzcyI6ImJhY2tHb2xhbmciLCJleHAiOjE3MTQ2NTg2NTZ9.G0dQcVvb4YRkiaAv2prQ5aEwb3wty1-6QLFji62mSVc"
}
```

## Get List User API

Endpoint : GET /user/


Response Body Success :

```json
[
  {
    "status": "OK",
    "data": [
      {
        "id": "f446ce7c-57f6-4edf-8a6b-a67dc7162c61",
        "username": "Testing-1714311147",
        "email": "testing@example.com",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
      },
      {
        "id": "28e54fb7-fee5-4fda-8a46-e1bc86d7cc95",
        "username": "Testing-1714312883",
        "email": "testing@example.com",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
      }
    ]
  }
]
```


## Get User API

Endpoint : GET /user/:userId

Response Body Success :

```json 
{
  "status": "OK",
  "data": {
    "id": "6b45af65-e7f9-42eb-b47b-ed952a894c4b",
    "username": "Testing-1714658608",
    "email": "testing@example.com",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  "description": "Success get user Data"
}
```