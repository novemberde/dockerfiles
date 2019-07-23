FROM golang:1.12.4-alpine3.9

LABEL IMAGE_NAME="load-testing-server"

ENV PORT 8080
EXPOSE 8080

COPY main.go main.go

CMD ["go", "run", "main.go"]