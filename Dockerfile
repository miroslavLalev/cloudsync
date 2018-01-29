FROM golang:alpine

ENV SRC_DIR /go/src/github.com/miroslavLalev/cloudsync/src

ADD ./src ${SRC_DIR}

WORKDIR ${SRC_DIR}

RUN apk update && apk add git
RUN go get -d ./...

CMD [ "go", "run", "main.go" ]