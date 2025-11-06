# Stage 1: Build the Go binary
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/server ./cmd/server

# Stage 2: Create a minimal final image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Copy the data files
COPY data ./data

EXPOSE 8080

CMD ["/app/server"]