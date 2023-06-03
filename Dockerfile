# Start from a base Go image
FROM golang:1.17-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the go.mod file to the working directory
COPY go.mod .

# Download Go dependencies
RUN go mod download

# Copy the rest of the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o app

# Start a new stage
FROM alpine:latest

# Set the working directory in the new stage
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/app .

# Expose the port that the server listens on
EXPOSE 8080

# Run the Go application
CMD ["./app"]
