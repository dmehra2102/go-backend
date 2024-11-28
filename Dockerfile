# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

# Run stage
FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "/app/main" ]