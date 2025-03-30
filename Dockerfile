FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY ./configs/.env ./configs/.env

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/configs/.env ./configs/.env

RUN chmod +x ./main

EXPOSE 8888

CMD ["./main"]
