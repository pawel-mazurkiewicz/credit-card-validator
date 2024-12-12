# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o credit-card-validator

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/credit-card-validator .
EXPOSE 8080
CMD ["./credit-card-validator"]