FROM golang:alpine

# Set working directory and copy code into working directory
WORKDIR /go/src/github.com/edwintcloud/goForum
COPY . .

# install git so we can use go get to get the database driver for postgres
RUN apk update && apk add git

# build app
RUN go get -u "github.com/lib/pq"
RUN go build

# Run the app
CMD ["./goForum"]