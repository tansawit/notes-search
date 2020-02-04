FROM golang:alpine
COPY . /go/src/app
WORKDIR /go/src/app
CMD ls
CMD ["go", "run", "server/app.go"]
