# Start from the official Go image
FROM golang:1.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application and name the output binary as "short-url"
RUN CGO_ENABLED=0 GOOS=linux go build -o short-url ./app/short-url

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary built in the previous stage
COPY --from=builder /app/short-url .

# Expose port 8080
EXPOSE 8080

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./short-url"]
