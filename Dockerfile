FROM golang:1.22-alpine AS backend_builder
WORKDIR /app
# Install required build tools
RUN apk add --no-cache gcc musl-dev

# Copy and download Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the backend code
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -o api ./cmd/api/main.go

# Build stage for Frontend
FROM node:20-alpine AS frontend_builder
WORKDIR /app

# Copy frontend files
COPY frontend/package*.json ./
# Install pnpm
RUN npm install -g pnpm
# Install dependencies
RUN pnpm install

COPY frontend/ .
# Build frontend
RUN pnpm run build

# Final stage for backend
FROM alpine:3.19 AS backend
WORKDIR /app
# Copy the binary from backend builder
COPY --from=backend_builder /app/api .
# Copy any other necessary files (like .env if needed)
COPY .env .

# Create directory for SQLite database
RUN mkdir -p /app/db && \
    chmod 777 /app/db

EXPOSE 8080
CMD ["./api"]

# Final stage for frontend
FROM nginx:alpine AS frontend
# Copy built frontend files
COPY --from=frontend_builder /app/dist /usr/share/nginx/html
# Copy custom nginx configuration if needed
# COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
