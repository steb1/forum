# # Use an official Golang runtime as the base image
# FROM golang:1.16 AS builder

# # Set the working directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download and install go dependencies
# RUN go mod download

# # Copy the entire project to the container
# COPY . .

# RUN go mod download golang.org/x/net

# # Use a minimal base image for the final container
# FROM alpine:latest

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the binary from the builder stage
# COPY --from=builder /app/webforum .

# # Expose the port that your web forum listens on
# EXPOSE 8080

# # Run the web forum binary
# CMD ["./webforum"]

# syntax=docker/dockerfile:1
FROM golang:1.19
## To make things easier when running the rest of our commands, let’s create a directory inside the image that we are building

WORKDIR /app
## Copying go.mod
COPY go.mod ./  go.sum ./

# Download and install go dependencies
RUN go mod download

## we need now to add our source code to our image
COPY . .
## We would like to compile our application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Expose the port that your web forum listens on
EXPOSE 8080

## Now, all that is left to do is to tell Docker what command to execute when our image is used to start a container.
CMD ["/docker-gs-ping"]