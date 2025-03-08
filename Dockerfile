# Use the official Golang image as a base
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Install dependencies required for Air
RUN apk add --no-cache git curl && \
    go install github.com/cosmtrek/air@latest && \
    cp /go/bin/air /usr/local/bin/air  # Ensure Air is in PATH

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose the application port
EXPOSE 8080

# Run Air to enable live reload
CMD ["air", "-c", ".air.toml"]
