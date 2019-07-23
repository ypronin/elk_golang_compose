# Start from golang:1.12 base image
FROM golang:1.12

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/elk_golang_compose

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...


RUN make build

# This container exposes port 8080 to the outside world
EXPOSE 8000

# Run the executable
CMD ["./bin/app"]