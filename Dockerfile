# Use an official Golang runtime as the base image
FROM golang:1.16 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download and install go dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webforum .

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/webforum .

# Expose the port that your web forum listens on
EXPOSE 8080

# Run the web forum binary
CMD ["./webforum"]
