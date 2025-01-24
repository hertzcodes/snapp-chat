    FROM golang:1.23-alpine3.19 AS builder
    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server    ./cmd/main.go
    
    FROM alpine:3.19
    WORKDIR /app
    
    COPY --from=builder /app/server    .
    
    CMD ["./server"]
        