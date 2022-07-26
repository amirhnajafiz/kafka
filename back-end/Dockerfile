# Start from the latest golang base image
FROM golang:1.18:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Installing Git for alpine
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
WORKDIR /app/cmd/runner
RUN go build -o /runner

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /runner .

ENTRYPOINT ["./runner"]