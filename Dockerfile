FROM golang:1.19.2

WORKDIR /yana_project

COPY main.go .
COPY go.sum .
COPY go.mod .

RUN go get -d -v ./...
RUN go install -v ./...
