FROM golang:1.24 AS builder
LABEL authors="neenayal"

ENTRYPOINT ["top", "-b"]

# Set the working directory
WORKDIR /penny-pot

# Copy go.mod and go.sum first (to cache deps)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o penny-pot .

# --- Stage 2: Deploy ---
FROM alpine:latest

# Set working directory
WORKDIR /penny-pot

# Copy the binary from the builder
COPY --from=builder /penny-pot/penny-pot .

# Expose port (optional, just for clarity)
EXPOSE 8080

# Command to run the executable
CMD ["./penny-pot"]

