FROM golang:latest
RUN mkdir $GOPATH/src/notesearch/
WORKDIR $GOPATH/src/notesearch/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD ["go", "run", "app.go", "connection.go","load.go","search.go"]
WORKDIR $GOPATH/src/notesearch/public
RUN npm install
RUN npm start