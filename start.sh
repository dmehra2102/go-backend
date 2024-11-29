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

# Install dos2unix to convert line endings
RUN apk add --no-cache dos2unix && dos2unix /app/start.sh

# Ensure start.sh is executable
RUN chmod +x /app/start.sh

EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["/app/start.sh"]