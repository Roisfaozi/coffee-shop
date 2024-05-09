APP=server
BUILD="./build/$(APP)"
DB_DRIVER=postgres
DB_SOURCE="postgres://rois:rois@localhost:5436/go-coffee-shop?sslmode=disable"
MIGRATIONS_DIR=./database/migrations
# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


install:
	go get -u ./... && go mod tidy

run:
	CGO_ENABLED=0 GOOS=linux go build -o ${BUILD} ./cmd/main.go

test:
	go test -cover -v ./...

migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql $(name)

migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0

compose-up:
	docker compose up -d --force-recreate

compose-down:
	docker compose stop && docker compose down && docker rmi go-server

docker-build:
	docker rmi rois/cafe-server && docker build -t rois/cafe-server .