# Use golang 1.17 as the builder stage
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /build

# Copy go.mod and go.sum to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies using go mod
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application with CGO disabled for a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app main.go

# Use a minimal Alpine image for the final stage
FROM alpine:3.14

# Install CA certificates necessary for HTTPS communication
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage to the final image
COPY --from=builder /build/app /bin/

# Set the command to run the binary executable
CMD ["/bin/app"]
