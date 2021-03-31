FROM golang:1.16.2-alpine3.13

WORKDIR /usr/src/

COPY ./ /usr/src/

RUN echo "building"

CMD go run main.go
