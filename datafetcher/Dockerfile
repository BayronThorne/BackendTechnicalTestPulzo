# Build stage: compile the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy entire project (go.mod is at root)
COPY . .

# Build the datafetcher binary from the datafetcher subfolder
RUN go build -o datafetcher ./datafetcher

# Final stage: lightweight alpine image with necessary certificates
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/datafetcher .

CMD ["./datafetcher"]
