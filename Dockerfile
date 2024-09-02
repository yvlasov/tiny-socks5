# Start with a base image that includes Go
FROM golang:1.20-alpine


RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod ./
COPY go.sum ./

# Download all dependencies (using go mod download to reduce build times during subsequent runs)
# Download all dependencies (using go get to fetch the latest version)
#RUN go get -d -v ./...

#RUN go mod tidy
#RUN go mod download
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o socks5-proxy main.go

# Expose the port the app runs on
EXPOSE 1080

# Command to run the executable
CMD ["./socks5-proxy"]


