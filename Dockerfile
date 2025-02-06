# Use an official Golang image as a base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache go modules
RUN go mod download

# Copy the rest of the application code
COPY .. .

# Build the application
RUN go build -o main ./cmd/app

# Expose the Prometheus metrics port
EXPOSE 2112

# Start the application
CMD ["./main"]
