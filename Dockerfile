# Use an official Go runtime as a parent image
FROM golang:1.20-alpine AS builder


# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o khetiapp ./cmd/kheti

# Use a minimal Alpine image for the final runtime
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the executable from the builder stage into the final image
COPY --from=builder /app/khetiapp .

# Create a directory for the config file
RUN mkdir -p /app/configs

# Copy the config file
COPY ./cmd/kheti/configs/config.json /app/configs/

# Expose the port on which the Go application will run
EXPOSE 8080

# Command to run the Go application
CMD ["./khetiapp"]