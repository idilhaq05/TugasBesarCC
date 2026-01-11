    FROM golang:latest AS builder

    WORKDIR /app/src/backend
    COPY go.mod go.sum ./
    RUN go mod tidy

    COPY . .
    
    RUN CGO_ENABLED=0 GOOS=linux go build -o backend

    EXPOSE 8081

    CMD ["./backend"]

