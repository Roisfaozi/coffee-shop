FROM golang:1.22.2-alpine AS build
LABEL authors="Rois Faozi"website:"roisfaozi.com"

WORKDIR coffeback

COPY . .

RUN go mod download

RUN apk --no-cache add bash
COPY wait_for_postgres.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/wait_for_postgres.sh

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /coffeback/server ./cmd/main.go

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["wait-for-postgres.sh", "cafe_db", "migrate", "-path", "./database/migrations", "-database", "postgresql://rois:rois@cafe_db:5432/go-coffee-shop?sslmode=disable", "-verbose", "up"]

FROM alpine:3.14

WORKDIR /app

COPY --from=build /coffeback /app/
ENV PATH="/app:${PATH}"

EXPOSE 8081

ENTRYPOINT ["server"]