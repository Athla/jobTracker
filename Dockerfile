# Build stage
FROM golang:1.23.4-alpine AS builder

RUN apk add --no-cache \
    sqlite \
    sqlite-dev \
    gcc \
    musl-dev \
    make

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main cmd/api/main.go

# Create directory for SQLite database
RUN mkdir -p /app/db && chmod 777 /app/db

# Run migrations and setup before starting the app
CMD ["sh", "-c", "go run cmd/migrations/main.go && go run cmd/setup/main.go && ./main"]
