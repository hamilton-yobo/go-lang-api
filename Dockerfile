FROM golang:1.22.4

# set working directory
WORKDIR /go/src/app

# copy the source code
copy . .

# expose the port
EXPOSE 8000

# Build the Go app
RUN go buil -o main cmd/main.go

# Run the executable
CMD ["./main"]