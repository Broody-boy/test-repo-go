FROM golang:1.25.3-alpine

WORKDIR /app

# Copy go.mod and go.sum first to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the binary from cmd/main.go and output as "app"
RUN go build -o app ./cmd/main.go

# Command to run the binary
CMD ["/app/app"]
