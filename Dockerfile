FROM golang:alpine

# Set working directory and copy code into working directory
WORKDIR /go/src/github.com/edwintcloud/goForum
COPY . .

# build app
RUN go build

# Run the app
CMD ["./goForum"]