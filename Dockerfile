# Start from the golang base image with Alpine
FROM golang:alpine as builder

# Install git, required for fetching the dependencies.
# gcc and musl-dev are required to compile any dependencies that require cgo.
RUN apk add --no-cache git gcc musl-dev

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go mod.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
# Disabling cgo with CGO_ENABLED=0 to make the binary statically linked.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app .

# Start a new stage from scratch using Alpine for a smaller, secure final image
FROM alpine:latest

# Install ca-certificates, tzdata for SSL and time zone configurations,
# and dumb-init for proper signal forwarding and zombie process reaping.
RUN apk add --no-cache ca-certificates tzdata dumb-init

# Set the Time Zone
ENV TZ=Asia/Makassar

# Create a directory to hold uploads if your app needs to save files
RUN mkdir -p /uploads

# Copy the built binary from the builder stage to the production stage
COPY --from=builder /go/bin/app /app

# Optionally, if your application requires environment variables defined in a .env file:
# Copy .env.example and rename it to .env
# COPY .env.example /env.example
# RUN mv /env.example /.env
COPY .env /.env

# Define the entrypoint to use dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# Run the web service on container startup.
CMD ["/app"]
