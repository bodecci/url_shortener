# Stage 1: Build the Go application
FROM golang:1.21-alpine as builder

# Install git (required by `go mod download` if using remote dependencies)
RUN apk add --no-cache git

# Set the build directory
WORKDIR /build

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies early, which can be cached
RUN go mod download

# Copy the entire source code after dependencies are downloaded
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run the built application
FROM alpine:latest

# Add a non-root user
RUN adduser -S -D -H -h /app appuser

# Set the working directory
WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /build/main /app/

# Expose port for the application
EXPOSE 3000

# Set the user to the non-root user
USER appuser

# Run the app
CMD ["./main"]
