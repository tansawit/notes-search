FROM golang:latest
RUN mkdir $GOPATH/src/notesearch/
WORKDIR $GOPATH/src/notesearch/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY app.go connection.go load.go search.go go.mod go.sum ./
RUN go run app.go connection.go load.go search.go &

FROM node:carbon
RUN mkdir /usr/src/app
WORKDIR /usr/src/app
ENV PATH /app/node_modules/.bin:$PATH
COPY public/ .
RUN ls
RUN npm install --silent
RUN npm install react-scripts -g --silent
CMD ["npm", "start"]