FROM golang:latest as builder

WORKDIR /app

# Copy the Go files to the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bot ./cmd/klubbot/main.go

# Path: Dockerfile
FROM alpine:latest as production

# Install ca-certificates
RUN apk update && apk add chromium ca-certificates && rm -rf /var/cache/apk/*

# Set the working directory inside the container
WORKDIR /app

# Copy the build artifact into the container
COPY --from=builder /app/bot .

# Copy the .env file into the container
# COPY .env .
COPY go.mod .

EXPOSE 8080

# Set the command to run the binary
CMD ["./bot"]