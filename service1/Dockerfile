# Use the official Go image for building
FROM golang:1.20-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o service

# Expose the required port
EXPOSE 8080

# Command to run the executable
CMD ["./service"]
