FROM golang:1.12.4-alpine3.9

LABEL Maintainer="https://github.com/novemberde"

RUN apk update && \
    apk upgrade && \
    apk add git

RUN apk add --update nodejs nodejs-npm
RUN apk add --no-cache git bash build-base libffi-dev openssl-dev bzip2-dev zlib-dev readline-dev sqlite-dev
RUN apk add linux-headers

RUN apk add --no-cache python3 && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --upgrade pip setuptools

RUN pip3 install awscli --upgrade
RUN apk add --no-cache openssh

RUN apk -v --purge del py-pip && \
    rm /var/cache/apk/*

CMD ["go"]