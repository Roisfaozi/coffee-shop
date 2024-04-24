<h1 align="center">
  Golang Coffe shop Restfull Api With Gin-gonic
</h1>

<p align="center"><img src="https://yt3.ggpht.com/ytc/AKedOLT7YD9x6PiR-CfbBbFC3wz2WatiIZFrI_I0v-6k=s900-c-k-c0x00ffffff-no-rj" width="300px" alt="fazztrack.svg" /></p>

<p align="center">
    <a href="https://www.roisfaozi.com/" target="blank">Our Website</a>
    ·
    <a href="https://www.fazztrack.com/class/fullstack-website-dan-golang">Join With Us</a>
    ·
</p>

## 🛠️ Project Structure

```bash
.
├── README.md
├── cmd
│   └── main.go
├── database
│   └── database.sql
├── exception
│   ├── error_handler.go
│   └── not_found_error.go
├── helper
│   └── error.go
├── internals
│   ├── handlers
│   ├── models
│   ├── repository
│   ├── routers
├── pkg
│   ├── postgres.go
    └── db.go
```

# API Specification

## User

Endpoint to create new user read <a href="https://github.com/Roisfaozi/coffee-shop/blob/main/docs/user.md" target="blank">this documentation</a>

## Products

Endpoint to create new user read <a href="https://github.com/Roisfaozi/coffee-shop/blob/main/docs/product.md" target="blank">this documentation</a>

## Products

Endpoint to create new user read <a href="https://github.com/Roisfaozi/coffee-shop/blob/main/docs/favorite.md" target="blank">this documentation</a>

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/Roisfaozi/coffee-shop.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```

3. Run the app

```bash
go run ./cmd/main.go
```

🌟 You are all set!

## 💻 Built with

- [Golang](https://go.dev/): programming language
- [Gin-gonic](https://gin-gonic.com/): for handle http request
- [Sqlx](http://jmoiron.github.io/sqlx/): for query database
- [Postgres](https://www.postgresql.org/): for DBMS

<hr>
<p align="center">
Developed with ❤️ in Indonesia 	🇮🇩
</p>
