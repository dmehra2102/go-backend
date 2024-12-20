# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Run stage
FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/main .

COPY app.env .

COPY start.sh .

COPY db/migration ./db/migration

EXPOSE 8080

RUN chmod +x /app/start.sh

CMD [ "/app/main" ]

ENTRYPOINT [ "/app/start.sh" ]