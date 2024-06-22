# Start from a minimal Alpine Linux base image with the latest version of Go installed
FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go binary into the container
COPY exporter .

# Command to run the executable
CMD ["./exporter"]