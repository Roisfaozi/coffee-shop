# Product API Spec

## Create Product API

Endpoint : POST /products/


Request Body :

```json
{
  "name": "Nasi Goreng v#",
  "price": 30000,
  "currency": "IDR",
  "description": "Deskripsi produk kopi yang luar biasa",
  "category": "Makanan",
  "image_url": "https://example.com/tshirt.jpg"
}
```

Response Body Success :

```json
{
  "id": "2b6f9510-fa3e-4780-b2c1-44ef8ea7ade6"
}
```


## Update Product API

Endpoint : PUT /products/:productsId

Request Body :

```json
{
  "name": "Nasi Goreng v3",
  "price": 30000,
  "currency": "IDR",
  "description": "Deskripsi produk kopi yang luar biasa",
  "category": "Makanan",
  "image_url": "https://example.com/tshirt.jpg"
}
```

Response Body Success :

```json
{
  "message": "Product updated successfully"
}
```

## Get List Product API

Endpoint : GET /products/


Response Body Success :

```json
[
  {
    "id": "1c442516-a2ad-4415-8382-55a3ea54f438",
    "name": "Nasi Goreng",
    "price": 30000,
    "currency": "IDR",
    "description": "Deskripsi produk kopi yang luar biasa",
    "image_url": "https://example.com/tshirt.jpg",
    "category": "Makanan",
    "created_at": "2024-04-24T23:52:11.442218Z",
    "updated_at": "2024-04-24T23:52:11.442218Z",
    "sizes": [
      {
        "id": "fe140f05-ddbe-4529-b65b-c07111e0d45c",
        "size_name": "R"
      },
      {
        "id": "bcfef6fe-d2a1-423d-8623-3cd4157a7e21",
        "size_name": "L"
      },
      {
        "id": "5c8a1a69-346c-4cb6-8921-b93f61a92d88",
        "size_name": "XL"
      }
    ]
  }
]
```


## Get Product API

Endpoint : GET /products/:productsId

Response Body Success :

```json 
{
  "id": "2b6f9510-fa3e-4780-b2c1-44ef8ea7ade6",
  "name": "Nasi Goreng v3",
  "price": 30000,
  "currency": "IDR",
  "description": "Deskripsi produk kopi yang luar biasa",
  "image_url": "https://example.com/tshirt.jpg",
  "category": "Makanan",
  "created_at": "2024-04-25T00:40:48.245353Z",
  "updated_at": "2024-04-25T00:41:14.858454Z",
  "sizes": [
    {
      "id": "fe140f05-ddbe-4529-b65b-c07111e0d45c",
      "size_name": "R"
    },
    {
      "id": "bcfef6fe-d2a1-423d-8623-3cd4157a7e21",
      "size_name": "L"
    },
    {
      "id": "5c8a1a69-346c-4cb6-8921-b93f61a92d88",
      "size_name": "XL"
    }
  ]
}
```


## Remove Product API

Endpoint : DELETE /products/:productsId

Response Body Success :

```json
{
  "message": "Product deleted successfully"
}
```
