# Use the official Golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -o myapp .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./myapp"]
