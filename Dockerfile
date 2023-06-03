FROM golang:1.16

WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Install the required dependencies
RUN go get github.com/jaytaylor/html2text@v0.0.0-20200129193226-29cc3fb31f6e
RUN go get github.com/james-bowman/nlp@v0.4.0

# Build the Go application
RUN go build -o main .

# Expose the port on which the server will listen
EXPOSE 8080

# Run the Go application
CMD ["./main"]
