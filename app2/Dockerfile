# Start from golang base image
FROM golang:1.19.3-alpine3.16 as builder

WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# Start a new stage from scratch
FROM alpine:3.15

# Copy the Pre-built binary files
COPY --from=builder /app/main .

# Expose port
EXPOSE 8082

#Command to run the executable
CMD ["./main"]
