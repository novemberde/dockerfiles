FROM golang:1.16rc1-buster as builder

WORKDIR /app

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main main.go

FROM ubuntu

ENV DELIVERY_STREAM_NAME=test
ENV PORT=8080
WORKDIR /app

COPY --from=builder /app/main ./main

CMD ["./main"]