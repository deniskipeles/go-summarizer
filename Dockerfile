FROM golang:1.16

WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Download Go dependencies
# RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port on which the server will listen
EXPOSE 8080

# Run the Go application
CMD ["./main"]
