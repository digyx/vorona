FROM golang:1.17-alpine

LABEL maintainer="Devon Tingley <dtingley@twilit.io>"

WORKDIR /usr/src/app
RUN apk add build-base

COPY . .
RUN go build -o app

EXPOSE 8080
CMD ["./app"]
