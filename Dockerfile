# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz
RUN tar -xvzf migrate.tar.gz

# Run stage
FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh /app/
COPY db/migration ./db/migration

# Install dos2unix and convert the line endings of start.sh
RUN apk add --no-cache dos2unix
RUN dos2unix /app/start.sh

EXPOSE 8080

# Ensure the start.sh script has the correct permissions
RUN chmod +x /app/start.sh

# Debug step to list the contents of the /app directory and check permissions
RUN ls -l /app

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]