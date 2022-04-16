FROM golang:alpine as builder

LABEL maintainer="Devon Tingley <dtingley@twilit.io>"

WORKDIR /usr/src/app
RUN apk add build-base git

COPY . .
RUN rm vorona.db

RUN go mod download
RUN go build -o vorona

EXPOSE 8080


FROM alpine:latest
RUN apk --update-cache upgrade
COPY --from=builder /usr/src/app/vorona /usr/loca/bin/vorona
CMD ["/usr/local/bin/vorona"]
