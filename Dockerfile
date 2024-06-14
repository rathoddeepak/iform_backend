# Start with the official Golang image
FROM golang:1.21-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o bin/main .

# Expose port 8888 to the outside world
EXPOSE 8888

# Command to migrate and start server
CMD ./bin/main  -iform migrate ; ./bin/main -iform start
