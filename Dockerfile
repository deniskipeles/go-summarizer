# Start with the official Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which the server will listen
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./main"]
