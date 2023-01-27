# BUILD
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o main cmd/app/main.go

# DEPLOY
FROM alpine:3.16.2

COPY --from=build /app/main main

ENTRYPOINT ["./main"]