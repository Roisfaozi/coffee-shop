# Favorite API Spec

## Create Favorite API

Endpoint : POST /favorite/:userId


Request Body :

```json
{
  "product_id": "2c8ccbe3-d652-4b76-8288-65743ca4b6ac"
}
```

Response Body Success :

```json
{
  "status": "Created",
  "data": {
    "id": "1633b9e2-eb07-4e92-8d92-894b40cb7ac1",
    "username": "Testing-1714658742",
    "email": "testing@example.com"
  },
  "description": "1 data user created"
}
```

Response Body Failed token :

```json
{
  "status": "Unauthorized",
  "description": "token has invalid claims: token is expired"
}
```

## Get List Favorite API

Endpoint : GET /favorite/:userId


Response Body Success :

```json
{
  "status": "OK",
  "data": [
    {
      "id": "900602f0-a1eb-48c6-b4e0-6769edf89d9f",
      "product_id": "2c8ccbe3-d652-4b76-8288-65743ca4b6ac",
      "user_id": "1633b9e2-eb07-4e92-8d92-894b40cb7ac1",
      "created_at": "2024-05-02T21:06:36.575991Z"
    },
    {
      "id": "6c5adad4-ea2f-4e6e-852f-25ca2f817f8c",
      "product_id": "2c8ccbe3-d652-4b76-8288-65743ca4b6ac",
      "user_id": "1633b9e2-eb07-4e92-8d92-894b40cb7ac1",
      "created_at": "2024-05-02T21:06:54.118864Z"
    }
  ],
  "description": "Favorites retrieved successfully"
}
```

Response Body Failed token :

```json
{
  "status": "Unauthorized",
  "description": "token has invalid claims: token is expired"
}
```

## Remove Favorite API

Endpoint : DELETE /Favorite/:favoriteId

Response Body Success :

```json
{
  "status": "OK",
  "description": "Product updated successfully"
}
```

Response Body Failed token :

```json
{
  "status": "Unauthorized",
  "description": "token has invalid claims: token is expired"
}
```