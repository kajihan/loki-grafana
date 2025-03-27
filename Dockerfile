# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
COPY pkg/ ./pkg/
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o webapp main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/webapp .
COPY public/ ./public/
COPY .env .env  # Optional: Include if you have a .env file
EXPOSE 8080
CMD ["./webapp"]
