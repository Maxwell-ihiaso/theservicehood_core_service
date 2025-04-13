FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go binary
RUN go build -o go-auth-app /app/cmd/app/main.go
# Expose the port on which the app will run
EXPOSE 8081

# Start the Go binary
CMD ["./go-auth-app"]
