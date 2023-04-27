FROM golang:latest as build

# Add Maintainer Info
LABEL maintainer="ali saboori"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY ["*.go" ,"./"]

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get . 

# Build the Go app
RUN go build -o shortlink -ldflags=-s .
# Command to run the executable
from centos:centos8
RUN mkdir /app
COPY --from=build /app/shortlink /app
WORKDIR /app
CMD ["./shortlink"]