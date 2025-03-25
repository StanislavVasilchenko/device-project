FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd/server

RUN go build -o main .

COPY ./init.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 8080

CMD ["./main"]