# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Add Maintainer Info
LABEL maintainer="xiaolong.gou@googlemail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

WORKDIR /dist
RUN cp /app/main .
FROM scratch

# Copy binary from build to main folder
COPY --from=builder /dist/main /

# Command to run the executable
ENTRYPOINT ["/main"]

# Expose port 8080 to the outside world
EXPOSE 8080