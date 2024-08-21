# Use the official Go image as the base
FROM golang:1.23-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the Go mod file
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the application
RUN go build -o main server.go

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]
