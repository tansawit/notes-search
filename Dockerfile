FROM golang:latest
RUN mkdir $GOPATH/src/notesearch/
WORKDIR $GOPATH/src/notesearch/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN ls
CMD ["go", "run", "app.go"]
