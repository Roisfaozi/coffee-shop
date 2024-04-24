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
  "id": "a7c025c3-f1b6-4f5a-a9e0-d3370d8ee4bc",
  "product_id": "2c8ccbe3-d652-4b76-8288-65743ca4b6ac",
  "user_id": "a3c1d8e0-5ce2-4e5e-bda6-2aaf5b067b34",
  "created_at": "2024-04-25T00:44:48.886238Z"
}
```

## Get List Favorite API

Endpoint : GET /favorite/:userId


Response Body Success :

```json
[
  {
    "id": "a7c025c3-f1b6-4f5a-a9e0-d3370d8ee4bc",
    "product_id": "2c8ccbe3-d652-4b76-8288-65743ca4b6ac",
    "user_id": "a3c1d8e0-5ce2-4e5e-bda6-2aaf5b067b34",
    "created_at": "2024-04-25T00:44:48.886238Z"
  }
]
```

## Remove Favorite API

Endpoint : DELETE /Favorite/:favoriteId

Response Body Success :

```json
{
  "message": "Favorite deleted successfully"
}
```
