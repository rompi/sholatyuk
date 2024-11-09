# Stage 1: Build the Go application
FROM golang:1.21.6 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o myapp command/app/*.go

# Stage 2: Create a smaller image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/myapp .

# Set the entry point to run the application
ENTRYPOINT ["./myapp"]