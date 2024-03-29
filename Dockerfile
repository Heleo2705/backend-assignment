FROM golang:1.21.3-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/api/main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY .env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./db/migrations


EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]