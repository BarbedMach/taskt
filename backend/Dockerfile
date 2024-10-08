FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY ./db/tables ./db/tables

RUN go build -o main .

FROM ubuntu:22.04
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/db/tables ./db/tables

CMD ["./main"]


