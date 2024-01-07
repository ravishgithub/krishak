# Use an official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy necessary files to the working directory
COPY ./cmd/kheti /app
COPY ./authentication /app/authentication

# Build the Go application
RUN go build -o main .

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
