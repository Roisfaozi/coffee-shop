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
  "image_url": "https://example.com/tshirt.jpg" //"path/to/file
}
```

Response Body Success :

```json
{
  "status": "Created",
  "data": {
    "id": "80d4defc-b576-4e12-a6f3-7cb2cbf607ff",
    "username": "Testing-1714658925",
    "email": "testing@example.com"
  },
  "description": "1 data user created"
}
```

Response Body Failed token permission :

```json
{
  "status": "Unauthorized",
  "description": "You not have permission"
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
  "image_url": "https://example.com/tshirt.jpg" //"path/to/file
}
```

Response Body Success :

```json
{
  "status": "OK",
  "description": " Product updated successfully"
}
```

Response Body Failed token permission :

```json
{
  "status": "Unauthorized",
  "description": "You not have permission"
}
```

## Get List Product API

Endpoint : GET /products/


Response Body Success :

```json
{
  "status": "OK",
  "data": [
    {
      "id": "474782d6-a163-4aa8-8f8c-5508c6cf92ff",
      "name": "Nasi Goreng v#",
      "price": 30000,
      "currency": "IDR",
      "description": "Deskripsi produk kopi yang luar biasa",
      "image_url": "https://example.com/tshirt.jpg",
      "category": "Makanan",
      "created_at": "2024-04-28T22:26:25.406898Z",
      "updated_at": "2024-04-28T22:26:25.406898Z",
      "sizes": [
        {
          "id": "ac53f05a-2d88-4562-bef1-301d32dff13e",
          "size_name": "R"
        },
        {
          "id": "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
          "size_name": "L"
        },
        {
          "id": "6ab889b0-a10d-4f92-ab26-d70710205a9e",
          "size_name": "XL"
        }
      ]
    },
    {
      "id": "4f620515-2ac7-4434-bb9e-a1d511d50708",
      "name": "Nasi Goreng v#",
      "price": 30000,
      "currency": "IDR",
      "description": "Deskripsi produk kopi yang luar biasa",
      "image_url": "https://example.com/tshirt.jpg",
      "category": "Makanan",
      "created_at": "2024-04-28T22:26:25.92881Z",
      "updated_at": "2024-04-28T22:26:25.92881Z",
      "sizes": [
        {
          "id": "ac53f05a-2d88-4562-bef1-301d32dff13e",
          "size_name": "R"
        },
        {
          "id": "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
          "size_name": "L"
        },
        {
          "id": "6ab889b0-a10d-4f92-ab26-d70710205a9e",
          "size_name": "XL"
        }
      ]
    }
  ]
}
```

Response Body Failed token permission :

```json
{
  "status": "Unauthorized",
  "description": "You not have permission"
}
```


## Get Product API

Endpoint : GET /products/:productsId

Response Body Success :

```json 
{
  "status": "OK",
  "data": {
    "id": "8e61447b-8d41-4f6c-acce-a8772c17b52e",
    "name": "Nasi Goreng Pedas",
    "price": 50000,
    "currency": "IDR",
    "description": "Deskripsi produk kopi yang luar biasa",
    "image_url": "http://res.cloudinary.com/dfs7nermk/image/upload/v1714658962/Coffee_shop/mxfeymtf2lc73jjm1lxz.jpg",
    "category": "Makanan",
    "created_at": "2024-05-02T20:39:33.115893Z",
    "updated_at": "2024-05-02T21:09:21.390339Z",
    "sizes": [
      {
        "id": "ac53f05a-2d88-4562-bef1-301d32dff13e",
        "size_name": "R"
      },
      {
        "id": "be25d8ae-ccc0-4b52-9e21-12633bc77bd2",
        "size_name": "L"
      },
      {
        "id": "6ab889b0-a10d-4f92-ab26-d70710205a9e",
        "size_name": "XL"
      }
    ]
  }
}
```

Response Body Failed token permission :

```json
{
  "status": "Unauthorized",
  "description": "You not have permission"
}
```


## Remove Product API

Endpoint : DELETE /products/:productsId

Response Body Success :

```json
{
  "status": "OK",
  "description": "Product deleted successfully"
}
```
