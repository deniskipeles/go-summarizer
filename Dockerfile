# Use an appropriate base image for your application
FROM golang:1.16

# Set the working directory
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which the server will listen
EXPOSE 8080

# Run the Go application
CMD ["./main"]
