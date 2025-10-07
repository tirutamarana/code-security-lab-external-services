# Use an official lightweight Go image
FROM golang:1.18 AS builder

# Set up working directory
WORKDIR /app

# Copy source code
COPY . .

# Build the Go binary
RUN go mod tidy && go build -o image-processing-api

# Use a minimal runtime image
FROM alpine:latest

# Set up working directory
WORKDIR /app

# Copy the built Go binary from builder stage
COPY --from=builder /app/image-processing-api /app/image-processing-api

CMD ["/app/image-processing-api"]
