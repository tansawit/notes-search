FROM golang:latest
RUN mkdir $GOPATH/src/notesearch/
WORKDIR $GOPATH/src/notesearch/
COPY . .
RUN go mod download
RUN chmod +x wait-for-elasticsearch.sh