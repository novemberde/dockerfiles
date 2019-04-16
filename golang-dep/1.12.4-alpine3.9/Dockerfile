FROM golang:1.12.4-alpine3.9

RUN apk update && \
    apk upgrade && \
    apk add git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/kardianos/govendor

CMD ["go"]