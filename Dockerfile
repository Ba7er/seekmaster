# Use a specific Go version as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/search/main.go

# Use a specific Alpine version for the final stage
FROM alpine:3.19

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=0 /app/main .

# Expose the port the app runs on
EXPOSE 9100

# Run the binary
CMD ["./main"]